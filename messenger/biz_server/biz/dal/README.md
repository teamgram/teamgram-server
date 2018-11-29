# DAL -- Data Access Layer

> 术语
> * DAL: Data Access Layer
> * DO:  Data Object
> * DAO: Data Access Object

```
// DO  --> 对应于数据库表
// DAO --> 对表的操作

/**
 <?xml version="1.0" encoding="UTF-8"?>
 <table sqlname="users">
	<operation name="insert">
 <sql>
 INSERT INTO
 users(app_id,user_id,avatar,nick,status,created_at,updated_at)
 VALUES (?,?,?,?,?,?,?)
 </sql>
	</operation>
	<operation name="selectByID">
 <sql>
 SELECT app_id,user_id,avatar,nick,status,created_at,updated_at FROM users WHERE id=?
 </sql>
	</operation>
 </table>
 */
// 如上, 可以通过配置自动生成DO,DAO,DAOImpl对象
// users表对应UserDO
// DAO: insert, selectByID

```
