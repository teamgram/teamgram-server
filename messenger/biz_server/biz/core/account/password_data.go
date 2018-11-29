// Copyright (c) 2018-present,  NebulaChat Studio (https://nebula.chat).
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Author: Benqi (wubenqi@gmail.com)

package account

import (
	"bytes"
	"encoding/hex"
	"github.com/golang/glog"
	"github.com/nebula-chat/chatengine/mtproto"
)

/*
	account.noPassword#96dabc18 new_salt:bytes email_unconfirmed_pattern:string = account.Password;
	account.password#7c18141c current_salt:bytes new_salt:bytes hint:string has_recovery:Bool email_unconfirmed_pattern:string = account.Password;
	account.passwordSettings#b7b72ab3 email:string = account.PasswordSettings;
	account.passwordInputSettings#86916deb flags:# new_salt:flags.0?bytes new_password_hash:flags.0?bytes hint:flags.0?string email:flags.1?string = account.PasswordInputSettings;
	auth.passwordRecovery#137948a5 email_pattern:string = auth.PasswordRecovery;

	byte[] new_salt = currentPassword.new_salt;
	byte[] hash = new byte[new_salt.length * 2 + newPasswordBytes.length];
	System.arraycopy(new_salt, 0, hash, 0, new_salt.length);
	System.arraycopy(newPasswordBytes, 0, hash, new_salt.length, newPasswordBytes.length);
	System.arraycopy(new_salt, 0, hash, hash.length - new_salt.length, new_salt.length);
	req.new_settings.flags |= 1;
	req.new_settings.hint = hint;
	req.new_settings.new_password_hash = Utilities.computeSHA256(hash, 0, hash.length);
	req.new_settings.new_salt = new_salt;
*/

// TODO(@benqi): add error code
// PASSWORD_HASH_INVALID
// NEW_PASSWORD_BAD
// NEW_SALT_INVALID
// EMAIL_INVALID
// EMAIL_UNCONFIRMED

// EC F8 73 76 65 BC 77 5A

// case 1: 未设置密码
// case 2: 设置密码了但未设置email
// case 3: 设置密码了，已设置email但email未验证
// case 4: email已经验证

const (
	kStatePasswordNone             = 0
	kStateNoRecoveryPassword       = 1
	kStateEmailUnconfirmedPassword = 2
	kStateConfirmedPassword        = 3
)

const (
	kServerSaltLen = 8
	kSaltLen       = 16
	kHashLen       = 32
)

type passwordData struct {
	userId     int32
	serverSalt []byte
	salt       []byte
	hash       []byte
	// TODO(@benqi): process hint
	hint string
	// hasRecovery bool
	email string
	state int
	dao   *accountsDAO
}

func makeEMailPattern(email string) string {
	// TODO(@benqi): make pattern --> axxxa@domain.com
	return email
}

// hasRecovery

func (m *AccountModel) MakePasswordData(userId int32) (*passwordData, error) {
	var (
		err                    error
		serverSalt, salt, hash []byte
	)

	do := m.dao.UserPasswordsDAO.SelectByUserId(userId)
	if do == nil {
		err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_INTERNAL_SERVER_ERROR)
		glog.Error(err, ": not found user_password row, user_id: ", userId)
		return nil, err
	}

	serverSalt, err = hex.DecodeString(do.ServerSalt)
	if err != nil || len(serverSalt) != kServerSaltLen {
		err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_INTERNAL_SERVER_ERROR)
		glog.Error(err, ": db's server_salt error")
		return nil, err
	}

	if len(do.Salt) > 0 {
		salt, err = hex.DecodeString(do.Salt)
		if err != nil || len(salt) != kSaltLen {
			err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_INTERNAL_SERVER_ERROR)
			glog.Error(err, ": db's salt error")
			return nil, err
		}
	}

	if len(do.Hash) > 0 {
		hash, err = hex.DecodeString(do.Hash)
		if err != nil || len(hash) != kHashLen {
			err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_INTERNAL_SERVER_ERROR)
			glog.Error(err, ": db's hash error")
			return nil, err
		}
	}

	// TODO(@benqi): check data.
	data := &passwordData{
		userId:     userId,
		serverSalt: serverSalt,
		salt:       salt,
		hash:       hash,
		hint:       do.Hint,
		email:      do.Email,
		state:      int(do.State),
		dao:        m.dao,
	}
	return data, nil
}

func (m *AccountModel) CheckRecoverCode(userId int32, code string) error {
	do := m.dao.UserPasswordsDAO.SelectCode(userId)
	if do == nil {
		err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_INTERNAL_SERVER_ERROR)
		glog.Error(err, ": not found user_password row, user_id - ", userId)
		return err
	}

	// TODO(@benqi): FLOOD_WAIT, check attempts??

	if do.Code != code {
		err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_CODE_INVALID)
		glog.Errorf("%s: userId - %d, code - %s", err, userId, code)
		return err
	}
	return nil
}

// SESSION_PASSWORD_NEEDED
func (m *AccountModel) CheckSessionPasswordNeeded(userId int32) bool {
	// TODO(@benqi):  仅仅从数据库里取state字段
	do := m.dao.UserPasswordsDAO.SelectByUserId(userId)
	if do == nil {
		return false
	}
	return do.State == kStateNoRecoveryPassword || do.State == kStateConfirmedPassword
}

func (p *passwordData) saveToDB() {
	var (
		salt, hash string
	)
	if len(p.salt) != 0 {
		salt = hex.EncodeToString(p.salt)
	}
	if len(p.hash) != 0 {
		hash = hex.EncodeToString(p.hash)
	}

	p.dao.UserPasswordsDAO.Update(salt, hash, p.hint, p.email, int8(p.state), p.userId)
}

func (p *passwordData) getHasRecovery() *mtproto.Bool {
	return mtproto.ToBool(p.state == kStateConfirmedPassword)
}

func (p *passwordData) GetPassword() *mtproto.Account_Password {
	switch p.state {
	//case kStatePasswordNone:
	//	noPassword := &mtproto.TLAccountNoPassword{Data2: &mtproto.Account_Password_Data{
	//		NewSalt:                 p.serverSalt,
	//		EmailUnconfirmedPattern: "",
	//	}}
	//	return noPassword.To_Account_Password()
	//case kStateNoRecoveryPassword:
	//	password := &mtproto.TLAccountPassword{Data2: &mtproto.Account_Password_Data{
	//		NewSalt:     p.serverSalt,
	//		CurrentSalt: p.salt,
	//		Hint:        p.hint,
	//		HasRecovery: mtproto.ToBool(false),
	//		// TODO(@benqi): make pattern
	//		EmailUnconfirmedPattern: "",
	//	}}
	//	return password.To_Account_Password()
	//case kStateEmailUnconfirmedPassword:
	//	noPassword := &mtproto.TLAccountNoPassword{Data2: &mtproto.Account_Password_Data{
	//		NewSalt:                 p.serverSalt,
	//		EmailUnconfirmedPattern: makeEMailPattern(p.email),
	//	}}
	//	return noPassword.To_Account_Password()
	//case kStateConfirmedPassword:
	//	password := &mtproto.TLAccountPassword{Data2: &mtproto.Account_Password_Data{
	//		NewSalt:     p.serverSalt,
	//		CurrentSalt: p.salt,
	//		Hint:        p.hint,
	//		HasRecovery: mtproto.ToBool(true),
	//		// TODO(@benqi): make pattern
	//		EmailUnconfirmedPattern: "",
	//	}}
	//	return password.To_Account_Password()
	default:
		// bug.
		err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_INTERNAL_SERVER_ERROR)
		glog.Error(err)
		panic(err)
		return nil
	}
}

// TODO(@benqi): check salt, currentPasswordHash, newPasswordHash
func (p *passwordData) UpdatePasswordSetting(currentPasswordHash, newSalt, newPasswordHash []byte, hint, email string) error {
	var err error

	switch p.state {
	case kStatePasswordNone:
		// set password.
		//
		if len(currentPasswordHash) != 0 {
			err = mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_PASSWORD_HASH_INVALID)
			glog.Error(err, ": current_password_hash need empty.")
			return err
		}

		// check salt
		if !bytes.Equal(p.serverSalt, newSalt[:8]) {
			err = mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_NEW_SALT_INVALID)
			glog.Error(err, ": current_password_hash need empty.")
			return err
		}

		p.salt = newSalt
		// check new_password_hash
		p.hash = newPasswordHash

		// TODO(@benqi): hint
		p.hint = hint
		p.email = email

		if p.email != "" {
			// TODO(@benqi): Check email invalid
			p.state = kStateEmailUnconfirmedPassword
		} else {
			p.state = kStateNoRecoveryPassword
			err = mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_EMAIL_UNCONFIRMED)
		}

	case kStateEmailUnconfirmedPassword:
		// check currentPasswordHash, currentPasswordHash: req.current_password_hash = new byte[0]
		//
		// email未验证状态时，客户端可以取消但不能更改密码
		// email验证后由服务端修改状态 --> kStateConfirmedPassword
		if len(currentPasswordHash) != 0 {
			err = mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_PASSWORD_HASH_INVALID)
			glog.Error(err, ": current_password_hash need empty.")
			return err
		}

		// check currentPasswordHash, currentPasswordHash: req.current_password_hash = new byte[0]
		if email != "" {
			err = mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_EMAIL_INVALID)
			glog.Error(err, ": email need empty.")
			return err
		}

		if len(newPasswordHash) != 0 {
			err = mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_NEW_PASSWORD_BAD)
			glog.Error(err, ": new_password_hash need empty.")
			return err
		}

		if len(newSalt) != 0 {
			err = mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_NEW_SALT_INVALID)
			glog.Error(err, ": new_salt need empty.")
			return err
		}

		p.salt = []byte{}
		p.hash = []byte{}
		p.hint = ""
		p.email = ""
		p.state = kStatePasswordNone

	case kStateConfirmedPassword, kStateNoRecoveryPassword:
		// email已经有了
		if email != "" {
			err = mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_EMAIL_INVALID)
			glog.Error(err, ": email need empty.")
			return err
		}

		// check hash
		if !bytes.Equal(p.hash, currentPasswordHash) {
			err = mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_PASSWORD_HASH_INVALID)
			glog.Error(err, ": current_password_hash need empty.")
			return err
		}

		if len(newPasswordHash) != 0 {
			// check salt
			if !bytes.Equal(p.serverSalt, newSalt[:8]) {
				err = mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_NEW_SALT_INVALID)
				glog.Error(err, ": current_password_hash need empty.")
				return err
			}

			p.salt = newSalt
			p.hash = newPasswordHash
			// TODO(@benqi): hint
			p.hint = hint
		} else {
			p.salt = []byte{}
			p.hash = []byte{}
			p.hint = ""
			p.email = ""
			p.state = kStatePasswordNone
		}
	default:
		// bug.
		err = mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_INTERNAL_SERVER_ERROR)
		glog.Error(err)
		panic(err)
		return err
	}

	p.saveToDB()
	return err
}

func (p *passwordData) CheckPassword(passwordHash []byte) bool {
	return bytes.Equal(p.hash, passwordHash)
}

func (p *passwordData) GetPasswordSetting(currentPasswordHash []byte) (*mtproto.Account_PasswordSettings, error) {
	checked := p.CheckPassword(currentPasswordHash)
	if !checked {
		err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_PASSWORD_HASH_INVALID)
		glog.Error(err, ": current_password_hash - ", currentPasswordHash)
		return nil, err
	}

	setting := &mtproto.TLAccountPasswordSettings{Data2: &mtproto.Account_PasswordSettings_Data{
		Email: p.email,
	}}
	return setting.To_Account_PasswordSettings(), nil
}

func (p *passwordData) RequestPasswordRecovery() (*mtproto.Auth_PasswordRecovery, error) {
	// TODO(@benqi): FLOOD_WAIT
	passwordRecovery := &mtproto.TLAuthPasswordRecovery{Data2: &mtproto.Auth_PasswordRecovery_Data{
		EmailPattern: p.email,
	}}
	return passwordRecovery.To_Auth_PasswordRecovery(), nil
}
