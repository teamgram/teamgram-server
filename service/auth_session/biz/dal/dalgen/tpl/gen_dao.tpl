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

package mysql_dao

import(
	"github.com/nebula-chat/chatengine/mtproto"
	"github.com/nebula-chat/chatengine/service/auth_session/biz/dal/dataobject"
	"github.com/jmoiron/sqlx"
	"github.com/golang/glog"
	"fmt"
)

type {{.Name}}DAO struct {
	db *sqlx.DB
}

func New{{.Name}}DAO(db* sqlx.DB) *{{.Name}}DAO {
	return &{{.Name}}DAO{db}
}

{{range $i, $v := .Funcs }}
{{if eq .QueryType "INSERT"}}
{{template "INSERT" $v}}
{{else if eq .QueryType "SELECT_STRUCT_SINGLE"}}
{{template "SELECT_STRUCT_SINGLE" $v}}
{{else if eq .QueryType "SELECT_STRUCT_LIST"}}
{{template "SELECT_STRUCT_LIST" $v}}
{{else if eq .QueryType "SELECT_STRUCT_MAP"}}
{{template "SELECT_STRUCT_MAP" $v}}
{{else if eq .QueryType "UPDATE"}}
{{template "UPDATE" $v}}
{{else if eq .QueryType "DELETE"}}
{{template "DELETE" $v}}
{{end}}
{{end}}

{{define "INSERT"}}
// {{.Sql}}
// TODO(@benqi): sqlmap
func (dao *{{.TableName}}DAO) {{.FuncName}}(do *dataobject.{{.TableName}}DO) int64 {
	var query = "{{.Sql}}"
	r, err := dao.db.NamedExec(query, do)
	if err != nil {
		errDesc := fmt.Sprintf("NamedExec in {{.FuncName}}(%v), error: %v", do, err)
		glog.Error(errDesc)
	    panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	id, err := r.LastInsertId()
	if err != nil {
		errDesc := fmt.Sprintf("LastInsertId in {{.FuncName}}(%v)_error: %v", do, err)
		glog.Error(errDesc)
	    panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}
	return id
}
{{end}}

{{define "SELECT_STRUCT_SINGLE"}}
// {{.Sql}}
// TODO(@benqi): sqlmap
func (dao *{{.TableName}}DAO) {{.FuncName}}({{ range $i, $v := .Params }} {{if ne $i 0 }} , {{end}} {{.FieldName}} {{.Type}} {{end}}) (*dataobject.{{.TableName}}DO) {
{{if eq .ParamHasList "true"}}  var q = "{{.CompiledByNamedSql}}"
    query, a, err := sqlx.In(q, {{range $i, $v := .QueryParams }} {{if ne $i 0 }} , {{end}} {{.FieldName}} {{end}})
    rows, err := dao.db.Queryx(query, a...)
{{else}} var query = "{{.CompiledByNamedSql}}"
    rows, err := dao.db.Queryx(query, {{range $i, $v := .QueryParams }} {{if ne $i 0 }} , {{end}} {{.FieldName}} {{end}})
{{end}}
	if err != nil {
		errDesc := fmt.Sprintf("Queryx in {{.FuncName}}(_), error: %v", err)
		glog.Error(errDesc)
	    panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	do := &dataobject.{{.TableName}}DO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
            errDesc := fmt.Sprintf("StructScan in {{.FuncName}}(_), error: %v", err)
            glog.Error(errDesc)
            panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
	} else {
		return nil
	}

    err = rows.Err()
    if err != nil {
        errDesc := fmt.Sprintf("rows in {{.FuncName}}(_), error: %v", err)
        glog.Error(errDesc)
        panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
    }

	return do
}
{{end}}

{{define "SELECT_STRUCT_LIST"}}
// {{.Sql}}
// TODO(@benqi): sqlmap
func (dao *{{.TableName}}DAO) {{.FuncName}}({{ range $i, $v := .Params }} {{if ne $i 0 }} , {{end}} {{.FieldName}} {{.Type}} {{end}}) ([]dataobject.{{.TableName}}DO) {
{{if eq .ParamHasList "true"}}  var q = "{{.CompiledByNamedSql}}"
    query, a, err := sqlx.In(q, {{range $i, $v := .QueryParams }} {{if ne $i 0 }} , {{end}} {{.FieldName}} {{end}})
    rows, err := dao.db.Queryx(query, a...)
{{else}} var query = "{{.CompiledByNamedSql}}"
    rows, err := dao.db.Queryx(query, {{range $i, $v := .QueryParams }} {{if ne $i 0 }} , {{end}} {{.FieldName}} {{end}})
{{end}}
	if err != nil {
        errDesc := fmt.Sprintf("Queryx in {{.FuncName}}(_), error: %v", err)
        glog.Error(errDesc)
        panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	var values []dataobject.{{.TableName}}DO
	for rows.Next() {
        v := dataobject.{{.TableName}}DO{}

        // TODO(@benqi): 不使用反射
        err := rows.StructScan(&v)
        if err != nil {
            errDesc := fmt.Sprintf("StructScan in {{.FuncName}}(_), error: %v", err)
            glog.Error(errDesc)
            panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
        }
		values = append(values, v)
    }

    err = rows.Err()
    if err != nil {
        errDesc := fmt.Sprintf("rows in {{.FuncName}}(_), error: %v", err)
        glog.Error(errDesc)
        panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
    }

    return values
}
{{end}}

{{define "SELECT_STRUCT_MAP"}}
// {{.Sql}}
// TODO(@benqi): sqlmap
func (dao *{{.TableName}}DAO) {{.FuncName}}({{ range $i, $v := .Params }} {{if ne $i 0 }} , {{end}} {{.FieldName}} {{.Type}} {{end}}) ([]map[string]interface{}) {
{{if eq .ParamHasList "true"}}  var q = "{{.CompiledByNamedSql}}"
    query, a, err := sqlx.In(q, {{range $i, $v := .QueryParams }} {{if ne $i 0 }} , {{end}} {{.FieldName}} {{end}})
    rows, err := dao.db.Queryx(query, a...)
{{else}} var query = "{{.CompiledByNamedSql}}"
    rows, err := dao.db.Queryx(query, {{range $i, $v := .QueryParams }} {{if ne $i 0 }} , {{end}} {{.FieldName}} {{end}})
{{end}}
	if err != nil {
        errDesc := fmt.Sprintf("Queryx in {{.FuncName}}(_), error: %v", err)
        glog.Error(errDesc)
        panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	values := make([]map[string]interface{}, 0)

	for rows.Next() {
    	v := make(map[string]interface{})

        // TODO(@benqi): 不使用反射
        err := rows.MapScan(v)
        if err != nil {
            errDesc := fmt.Sprintf("MaptScan in {{.FuncName}}(_), error: %v", err)
            glog.Error(errDesc)
            panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
        }
		values = append(values, v)
    }

    err = rows.Err()
    if err != nil {
        errDesc := fmt.Sprintf("rows in {{.FuncName}}(_), error: %v", err)
        glog.Error(errDesc)
        panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
    }

    return values
}
{{end}}


{{define "UPDATE"}}
// {{.Sql}}
// TODO(@benqi): sqlmap
func (dao *{{.TableName}}DAO) {{.FuncName}}({{ range $i, $v := .Params }} {{if ne $i 0 }} , {{end}} {{.FieldName}} {{.Type}} {{end}}) int64 {
{{if eq .ParamHasList "true"}}  var q = "{{.CompiledByNamedSql}}"
    query, a, err := sqlx.In(q, {{range $i, $v := .QueryParams }} {{if ne $i 0 }} , {{end}} {{.FieldName}} {{end}})
    r, err := dao.db.Exec(query, a...)
{{else}} var query = "{{.CompiledByNamedSql}}"
    r, err := dao.db.Exec(query, {{range $i, $v := .QueryParams }} {{if ne $i 0 }} , {{end}} {{.FieldName}} {{end}})
{{end}}
	if err != nil {
        errDesc := fmt.Sprintf("Exec in {{.FuncName}}(_), error: %v", err)
        glog.Error(errDesc)
        panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	rows, err := r.RowsAffected()
	if err != nil {
        errDesc := fmt.Sprintf("RowsAffected in {{.FuncName}}(_), error: %v", err)
        glog.Error(errDesc)
        panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return rows
}
{{end}}

{{define "DELETE"}}
// {{.Sql}}
// TODO(@benqi): sqlmap
func (dao *{{.TableName}}DAO) {{.FuncName}}({{ range $i, $v := .Params }} {{if ne $i 0 }} , {{end}} {{.FieldName}} {{.Type}} {{end}}) int64 {
{{if eq .ParamHasList "true"}}  var q = "{{.CompiledByNamedSql}}"
    query, a, err := sqlx.In(q, {{range $i, $v := .QueryParams }} {{if ne $i 0 }} , {{end}} {{.FieldName}} {{end}})
    r, err := dao.db.Exec(query, a...)
{{else}} var query = "{{.CompiledByNamedSql}}"
    r, err := dao.db.Exec(query, {{range $i, $v := .QueryParams }} {{if ne $i 0 }} , {{end}} {{.FieldName}} {{end}})
{{end}}
	if err != nil {
        errDesc := fmt.Sprintf("Exec in {{.FuncName}}(_), error: %v", err)
        glog.Error(errDesc)
        panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	rows, err := r.RowsAffected()
	if err != nil {
        errDesc := fmt.Sprintf("RowsAffected in {{.FuncName}}(_), error: %v", err)
        glog.Error(errDesc)
        panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}
	return rows
}
{{end}}
