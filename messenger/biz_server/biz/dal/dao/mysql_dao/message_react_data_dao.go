package mysql_dao

import (
	"fmt"
	"github.com/golang/glog"
	"github.com/jmoiron/sqlx"
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/dal/dataobject"
	"github.com/nebula-chat/chatengine/mtproto"
)

type MessageReactDataDAO struct {
	db *sqlx.DB
}

func NewMessageReactDataDAO(db *sqlx.DB) *MessageReactDataDAO{
	return &MessageReactDataDAO{db}
}

func (dao *MessageReactDataDAO) Insert(do *dataobject.MessageReactDataDO) int64 {
	glog.Info(dao.db)
	glog.Info("Insert pushreact")
	var query = "insert ignore into message_react_data(react_data_id, react_id, message_data_id, sender_user_id,  `date3`, edit_date, deleted, created_at, updated_at) values (:react_data_id, :react_id, :message_data_id, :sender_user_id, :date3, :edit_date, :deleted, :created_at, :updated_at)"	
	r, err := dao.db.NamedExec(query, do)
	if err != nil {
		errDesc := fmt.Sprintf("NamedExec in Insert(%v), error: %v", do, err)
		glog.Info(errDesc)
		glog.Error(errDesc)
	}

	id, err := r.LastInsertId()
	if err != nil {
		errDesc := fmt.Sprintf("LastInsertId in Insert(%v)_error: %v", do, err)
		glog.Info(errDesc)
		glog.Error(errDesc)
	}
	return id
}

func (dao *MessageReactDataDAO) SelectByMessageId(sender_user_id, message_data_id int64) *dataobject.MessageReactDataDO{
	var query =  "select reaction_id, message_data_id, sender_user_id, date3, edit_date, deleted, created_at, updated_at where sender_user_id = ? and message_data_id = ? and deleted = 0 limit 1"
	rows, err := dao.db.Queryx(query, message_data_id)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectByMessageId(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	do := &dataobject.MessageReactDataDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectByMessageId(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
	} else {
		return nil
	}

	err = rows.Err()
	if err != nil {
		errDesc := fmt.Sprintf("rows in SelectByMessageId(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return do
}

