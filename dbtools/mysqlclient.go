package dbtools

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

/*
NOTICE:
	1, mysql manipulating sql without commit will not update the db
	2, how to batch the inserting or updating data
*/

type MysqlClient struct {
	client *sql.DB
}

func NewMysqlClient(host string, user string, password string, database string, charset string) *MysqlClient {
	info := fmt.Sprintf("%s:%s@%s/%s?charset=%s", user, password, host, database, charset)
	// e.x. : root:@/NameGender?charset=utf8
	fmt.Println(info)
	db, err := sql.Open("mysql", info)
	if err != nil {
		panic(err)
	}
	mysqlClient := &(MysqlClient{})
	mysqlClient.client = db
	return mysqlClient
}

// execute the query with returned rows
func (self *MysqlClient) Query(query string) *sql.Rows {
	rows, err := self.client.Query(query)
	if err != nil {
		panic(err)
	}
	return rows
}

// execute the query with at most one row
func (self *MysqlClient) QueryRow(query string) *sql.Row {
	row := self.client.QueryRow(query)
	return row
}

/*the usage for fetching the data from rows:
for rows.Next(){
	var id int
	var name string
	err := rows.Scan(&id, &name)
}
*/

// begin one transaction
func (self *MysqlClient) Begin() *sql.Tx {
	tx, err := self.client.Begin()
	if err != nil {
		panic(err)
	}
	return tx
}

// execute the query without returned rows, mainly for udpating sqls
func (self *MysqlClient) Exec(query string) {
	tx := self.Begin()
	_, err := tx.Exec(query)
	if err != nil {
		panic(err)
	}
	tx.Commit()
}

// close the client
func (self *MysqlClient) Close() {
	if self.client != nil {
		self.client.Close()
	}
}

// some frequently used query and the result
func (self *MysqlClient) Count(tablename string) int64 {
	q := fmt.Sprintf("select count(1) from %s", tablename)
	row := self.QueryRow(q)
	var count int64
	err := row.Scan(&count)
	if err != nil {
		panic(err)
	}
	return count
}

func (self *MysqlClient) DropTable(tablename string) {
	q := fmt.Sprintf("drop table %s", tablename)
	self.Exec(q)
}

func (self *MysqlClient) TruncateTable(tablename string) {
	q := fmt.Sprintf("truncate %s", tablename)
	self.Exec(q)
}

func (self *MysqlClient) IsTableExist(tablename string) bool {
	_, err := self.client.Query(fmt.Sprintf("select count(1) from %s", tablename))
	if err != nil {
		return false
	}
	return true
}

func (self *MysqlClient) Ping() {

}
