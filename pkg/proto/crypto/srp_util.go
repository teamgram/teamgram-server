// Copyright 2024 Teamgram Authors
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
//
// Author: teamgramio (teamgram.io@gmail.com)
//

/***

# 安全远程口令(SRP)
> 原始出处: https://blog.csdn.net/haoranstone/article/details/46407959
>
> SRP是一个基于口令的身份认证和密钥交换协议。SRP的优点在于，
> 认证过程中不会有密钥明文传递的现象，用户只需要持有口令即可，
> 此外，服务端存储的非用户的口令，而是与口令相关的信息，
> 即便服务器被敌手俘获，敌手也无法伪造一个合法的客户端(无法拿到口令) 从而保证了双方的安全。
> 下面介绍SRP协议的内容

## 使用参数说明
- NN 一个非常大的素数(N=2q+1N=2q+1,q是一个素数，所有的运算都是在mod N 上完成
- gg mod N上的生成元
- kk 乘数因子(SRP-6中，k=3)
- ss 用户的盐值(saltsalt)
- II 用户名
- pp 明文口令
- H()H() 单向hash
- ^^ 模幂运算
- uu 随机计算值
- a,ba,b 临时秘密值(Secret ephemeral values)
- A,BA,B 临时公开值(Public ephemeral values)
- xx 私钥(由pp和ss生成)
- vv 口令校验值

## 服务端存储信息
- 服务端(host）用以下的方程式计算用户的信息并保存， 用于验证客户端登陆。
- x=H(s,p) (s,随机选取)
- v=gx (v就是口令验证值)
- 服务端将I,s,v作为用户的认证信息保存起来

## 认证与协商过程
- User —> Host: I,A=gaI,A=ga (a 是一个随机选取的值， A是公开值)
- Host —> User: s,B=kv+gbs,B=kv+gb(发送盐值，和公开的B，其中bb是一个随机选取的值
- 双方计算：u=H(A,B)u=H(A,B)
- User: x=H(s,p)x=H(s,p)(注意，p是用户自己输入的口令)
- User: S=(B−kgx)(a+ux)S=(B−kgx)(a+ux) (S就是用户计算的会话密钥生成值)
- User: K=H(S)K=H(S) (用Hash S得出会话密钥)
- Host: S=(Avu)bS=(Avu)b(如果用户输入的口令是对的，那这里算出的S值和客户端算出的S值是一致的)
- Host:K=H(S)K=H(S) (同样Hash得出会话密钥）

## 验证双方生成密钥是否一致
- User—-> Host : M=H(H(N)xorH(g),H(I),s,A,B,K)M=H(H(N)xorH(g),H(I),s,A,B,K)
- Host—–>User: H(A,M,K)H(A,M,K)

## 分析
在这个过程中，双方只交换和公开了s, A, B, I 这四个信息，任何一个因子都不携带与口令相关的信息，
用户端在获取了s,Bs,B之后，利用ss和口令pp生成中间值xx, 即该值与服务端的计算v时的x值是一致的，
继而算出会话密钥，实际上S=(B−kgx)(a+ux)S=(B−kgx)(a+ux) 和S=(Avu)bS=(Avu)b是完全等价的，
从而完成了密钥的协商。这个过程中，敌手即便窃听了s,A,B,Is,A,B,I若是没有口令，也无法算出会话密钥。
同样，敌手即便攻破了服务器，他也只能获取到s,v,Is,v,I同样无法获取到用户的密钥，
当然，敌手利用s,v,Is,v,I 可以伪造服务器与客户进行通信。不过这不影响口令的安全。

## 参考资料
https://github.com/symeapp/srp-client

https://github.com/bitwiseshiftleft/sjcl

http://srp.stanford.edu/ndss.html

http://srp.stanford.edu/demo/demo.html

http://srp.stanford.edu/links.html

https://code.google.com/p/srpforjava

http://milk-36.iteye.com/blog/1687443

*/

package crypto

import (
	"bytes"
	"math/big"
	"math/rand"
	"time"
)

type PasswordKdfAlgoModPow struct {
	Salt1 []byte
	Salt2 []byte
	G     int32
	P     []byte
}

type SRPUtil struct {
	*PasswordKdfAlgoModPow
	g      *big.Int
	gBytes []byte
	p      *big.Int
	k      *big.Int
}

func MakeSRPUtil(algo *PasswordKdfAlgoModPow) *SRPUtil {
	rand.Seed(time.Now().UnixNano())

	// TODO(@benqi): check algo
	srp := &SRPUtil{
		PasswordKdfAlgoModPow: algo,
		g:                     big.NewInt(int64(algo.G)),
		p:                     new(big.Int).SetBytes(algo.P),
	}

	srp.gBytes = getBigIntegerBytes(srp.g)
	kBytes := calcSHA256(algo.P, srp.gBytes)
	srp.k = new(big.Int).SetBytes(kBytes)

	return srp
}

func (m *SRPUtil) CheckNewSalt1(newSalt1 []byte) bool {
	if len(newSalt1) < 8 {
		return false
	}

	return bytes.Equal(m.Salt1, newSalt1[:8])
}

func (m *SRPUtil) CalcSRPB(vBytes []byte) ([]byte, []byte) {
	v := new(big.Int).SetBytes(vBytes)

	bNonce := RandomBytes(256)
	//fmt.Println(hex.EncodeToString(bNonce))
	b := new(big.Int).SetBytes(bNonce)
	//bBytes := getBigIntegerBytes(b)
	//fmt.Println(hex.EncodeToString(bBytes))
	// Host —> User: s,B=k*v+g**b(发送盐值，和公开的B，其中b是一个随机选取的值
	B := new(big.Int).Mod(new(big.Int).Add(new(big.Int).Mul(m.k, v), new(big.Int).Exp(m.g, b, m.p)), m.p)
	//B := new(big.Int).Add(new(big.Int).Mul(m.k, v), new(big.Int).Exp(m.g, b, m.p))
	BBytes := getBigIntegerBytes(B)

	return bNonce, BBytes
}

func (m *SRPUtil) CalcSRPB2(bNonce, vBytes []byte) []byte {
	v := new(big.Int).SetBytes(vBytes)

	// bNonce := RandomBytes(256)
	//fmt.Println(hex.EncodeToString(bNonce))
	b := new(big.Int).SetBytes(bNonce)
	//bBytes := getBigIntegerBytes(b)
	//fmt.Println(hex.EncodeToString(bBytes))
	// Host —> User: s,B=k*v+g**b(发送盐值，和公开的B，其中b是一个随机选取的值
	B := new(big.Int).Mod(new(big.Int).Add(new(big.Int).Mul(m.k, v), new(big.Int).Exp(m.g, b, m.p)), m.p)
	//B := new(big.Int).Add(new(big.Int).Mul(m.k, v), new(big.Int).Exp(m.g, b, m.p))

	BBytes := getBigIntegerBytes(B)

	return BBytes
}

func (m *SRPUtil) CalcM(newSalt1, vBytes, srpA, srpb, srpB []byte) []byte {
	v := new(big.Int).SetBytes(vBytes)

	A := new(big.Int).SetBytes(srpA)
	if A.Cmp(bigIntZero) <= 0 || A.Cmp(m.p) >= 0 {
		return nil
	}
	ABytes := getBigIntegerBytes(A)

	//// Host —> User: s,B=kv+gb(发送盐值，和公开的B，其中b是一个随机选取的值
	b := new(big.Int).SetBytes(srpb)

	// B := new(big.Int).SetBytes(srpB)
	// BBytes := getBigIntegerBytes(B)
	// BBytes := srpB // getBigIntegerBytes(B)

	//_ = BBytes

	uBytes := calcSHA256(ABytes, srpB)
	u := new(big.Int).SetBytes(uBytes)
	if u.Cmp(bigIntZero) == 0 {
		return nil
	}
	//fmt.Println("uBytes: " + hex.EncodeToString(uBytes))
	//fmt.Println("vBytes: " + hex.EncodeToString(vBytes))

	// Host: S=(A*v**u)**b(如果用户输入的口令是对的，那这里算出的S值和客户端算出的S值是一致的)
	S := new(big.Int).Exp(new(big.Int).Mod(new(big.Int).Mul(A, new(big.Int).Exp(v, u, m.p)), m.p), b, m.p)
	SBytes := getBigIntegerBytes(S)

	KBytes := calcSHA256(SBytes)

	//fmt.Println("pHash: " + hex.EncodeToString(m.P))
	//fmt.Println("gHash: " + hex.EncodeToString(m.gBytes))
	//fmt.Println("newSalt1: " + hex.EncodeToString(newSalt1))
	//fmt.Println("m.Salt2: " + hex.EncodeToString(m.Salt2))
	//fmt.Println("ABytes: " + hex.EncodeToString(ABytes))
	//fmt.Println("BBytes: " + hex.EncodeToString(srpB))
	//fmt.Println("KBytes: " + hex.EncodeToString(KBytes))

	pHash := calcSHA256(m.P)
	gHash := calcSHA256(m.gBytes)
	for i := 0; i < len(pHash); i++ {
		pHash[i] = gHash[i] ^ pHash[i]
	}

	return calcSHA256(pHash, calcSHA256(newSalt1), calcSHA256(m.Salt2), ABytes, srpB, KBytes)
}

func (m *SRPUtil) GetX(newSalt1, passwordBytes []byte) []byte {
	var xBytes []byte

	xBytes = calcSHA256(newSalt1, passwordBytes, newSalt1)
	xBytes = calcSHA256(m.Salt2, xBytes, m.Salt2)
	xBytes = calcPBKDF2(xBytes, newSalt1)

	return calcSHA256(m.Salt2, xBytes, m.Salt2)
}

func (m *SRPUtil) GetV(newSalt1, passwordBytes []byte) *big.Int {
	xBytes := m.GetX(newSalt1, passwordBytes)
	x := new(big.Int).SetBytes(xBytes)

	return new(big.Int).Exp(m.g, x, m.p)
}

func (m *SRPUtil) GetVBytes(newSalt1, passwordBytes []byte) []byte {
	return getBigIntegerBytes(m.GetV(newSalt1, passwordBytes))
}

func (m *SRPUtil) CalcClientM(newSalt1, xBytes, srpB []byte) ([]byte, []byte) {
	if len(xBytes) == 0 || len(srpB) == 0 || !isGoodPrime(m.P, int(m.G)) {
		// fmt.Println("check error")
		return nil, nil
	}

	x := new(big.Int).SetBytes(xBytes)

	aBytes := RandomBytes(256)
	//fmt.Println(hex.EncodeToString(aBytes))
	a := new(big.Int).SetBytes(aBytes)

	A := new(big.Int).Exp(m.g, a, m.p)
	ABytes := getBigIntegerBytes(A)

	B := new(big.Int).SetBytes(srpB)
	if B.Cmp(bigIntZero) <= 0 || B.Cmp(m.p) >= 0 {
		// fmt.Println("b error")
		return nil, nil
	}
	BBytes := getBigIntegerBytes(B)

	uBytes := calcSHA256(ABytes, BBytes)
	u := new(big.Int).SetBytes(uBytes)
	if u.Cmp(bigIntZero) == 0 {
		// fmt.Println("u error")
		return nil, nil
	}

	//// Host: S=(A*v**u)**b(如果用户输入的口令是对的，那这里算出的S值和客户端算出的S值是一致的)
	//S := new(big.Int).Exp(new(big.Int).Mod(new(big.Int).Mul(A, new(big.Int).Exp(v, u, m.p)), m.p), b, m.p)
	//SBytes := getBigIntegerBytes(S)

	// new(big.Int).Exp(v, u, m.p)
	// new(big.Int).Mul(A, new(big.Int).Exp(v, u, m.p))
	// new(big.Int).Exp(new(big.Int).Mul(A, new(big.Int).Exp(v, u, m.p)), b, m.p)
	// User: S=(B−k*g**x)**(a+u*x) (S就是用户计算的会话密钥生成值)
	BKgx := new(big.Int).Sub(B, new(big.Int).Mod(new(big.Int).Mul(m.k, new(big.Int).Exp(m.g, x, m.p)), m.p))
	if BKgx.Cmp(bigIntZero) < 0 {
		// fmt.Println("<0")
		BKgx = new(big.Int).Add(BKgx, m.p)
	}

	if !isGoodGaAndGb(BKgx, m.p) {
		// fmt.Println("isGoodGaAndGb error")
		return nil, nil
	}

	S := new(big.Int).Exp(BKgx, new(big.Int).Add(a, new(big.Int).Mul(u, x)), m.p)
	SBytes := getBigIntegerBytes(S)

	KBytes := calcSHA256(SBytes)

	pHash := calcSHA256(m.P)
	gHash := calcSHA256(m.gBytes)
	for i := 0; i < len(pHash); i++ {
		pHash[i] = gHash[i] ^ pHash[i]
	}

	return ABytes, calcSHA256(pHash, calcSHA256(newSalt1), calcSHA256(m.Salt2), ABytes, BBytes, KBytes)

	// result.M1 = Utilities.computeSHA256(p_hash, Utilities.computeSHA256(algo.salt1), Utilities.computeSHA256(algo.salt2), A_bytes, B_bytes, K_bytes);

}

func (m *SRPUtil) CalcClientM2(newSalt1, aBytes, ABytes, xBytes, srpB []byte) []byte {
	if len(xBytes) == 0 || len(srpB) == 0 || !isGoodPrime(m.P, int(m.G)) {
		// fmt.Println("check error")
		return nil
	}

	x := new(big.Int).SetBytes(xBytes)

	// aBytes := RandomBytes(256)
	// fmt.Println(hex.EncodeToString(aBytes))
	a := new(big.Int).SetBytes(aBytes)
	//
	//A := new(big.Int).Exp(m.g, a, m.p)
	//ABytes := A //getBigIntegerBytes(A)

	B := new(big.Int).SetBytes(srpB)
	if B.Cmp(bigIntZero) <= 0 || B.Cmp(m.p) >= 0 {
		// fmt.Println("b error")
		return nil
	}
	BBytes := getBigIntegerBytes(B)

	uBytes := calcSHA256(ABytes, BBytes)
	u := new(big.Int).SetBytes(uBytes)
	if u.Cmp(bigIntZero) == 0 {
		// fmt.Println("u error")
		return nil
	}

	// User: S=(B−k*g**x)**(a+u*x) (S就是用户计算的会话密钥生成值)
	BKgx := new(big.Int).Sub(B, new(big.Int).Mod(new(big.Int).Mul(m.k, new(big.Int).Exp(m.g, x, m.p)), m.p))
	if BKgx.Cmp(bigIntZero) < 0 {
		// fmt.Println("<0")
		BKgx = new(big.Int).Add(BKgx, m.p)
	}

	if !isGoodGaAndGb(BKgx, m.p) {
		// fmt.Println("isGoodGaAndGb error")
		return nil
	}

	//fmt.Println("uBytes: " + hex.EncodeToString(uBytes))
	S := new(big.Int).Exp(BKgx, new(big.Int).Add(a, new(big.Int).Mul(u, x)), m.p)
	SBytes := getBigIntegerBytes(S)
	KBytes := calcSHA256(SBytes)
	pHash := calcSHA256(m.P)
	gHash := calcSHA256(m.gBytes)

	for i := 0; i < len(pHash); i++ {
		pHash[i] = gHash[i] ^ pHash[i]
	}

	return calcSHA256(pHash, calcSHA256(newSalt1), calcSHA256(m.Salt2), ABytes, BBytes, KBytes)
}
