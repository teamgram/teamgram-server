package mysql_dao

import (
	"fmt"
	"github.com/golang/glog"
	"github.com/jmoiron/sqlx"
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/dal/dataobject"
	"github.com/nebula-chat/chatengine/mtproto"
)

type MessageReactDAO struct {
	db *sqlx.DB
}

func NewMessageReactDAO(db *sqlx.DB) *MessageReactDAO{
	return &MessageReactDAO{db}
}

func (dao *MessageReactDAO) Insert(do *dataobject.MessageReactDO) int64 {
	var query = "insert ignore into message_datas(text, file_id, file_hash, file_size, width, height) values (:text, :file_id, :file_hash, :file_size, :date3, :deleted, :created_at, :updated_at)"
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

func (dao *MessageReactDataDAO) SelectByReactId(react_id int64) *dataobject.MessageReactDO{
	var query =  "select reaction_id, text, file_id, file_hash, file_size, width, height where react_id = ? and deleted = 0 limit 1"
	rows, err := dao.db.Queryx(query, react_id)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectByReactId(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	do := &dataobject.MessageReactDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectByReactId(_), error: %v", err)
			glog.Error(errDesc)
			//panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
	} else {
		return nil
	}

	err = rows.Err()
	if err != nil {
		errDesc := fmt.Sprintf("rows in SelectByReactId(_), error: %v", err)
		glog.Info(errDesc)
		glog.Error(errDesc)
		//panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return do
}

