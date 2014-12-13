package dbtools

import (
	"errors"
	"fmt"
	"testing"
)

func TestMysqlClient(t *testing.T) {
	host := ""
	user := "root"
	password := ""
	charset := "utf8"
	db := "NameGender"

	mysqlClient := NewMysqlClient(host, user, password, db, charset)
	count := mysqlClient.Count("url")
	if count != 827 {
		info := fmt.Sprintf("count is not right: %d", count)
		t.Fatal(errors.New(info))
	}
	// get the data
	rows := mysqlClient.Query("select id, url, add_time, is_fetched, is_finished from url")
	for rows.Next() {
		var id int
		var url string
		var add_time string
		var is_fetched int
		var is_finished int
		err := rows.Scan(&id, &url, &add_time, &is_fetched, &is_finished)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(fmt.Sprintf("%d:%s:%s:%d:%d", id, url, add_time, is_fetched, is_finished))
	}
	// create table
	q := "create table test_go (id int not null, name varchar(40))"
	mysqlClient.Exec(q)
	tablename := "test_go"
	isTableExist := mysqlClient.IsTableExist(tablename)
	if isTableExist != true {
		t.Fatal(errors.New("Error table test_go not exist"))
	}
	// insert into the table

	q = "insert into test_go values(1, 'lxf')"
	mysqlClient.Exec(q)
	//count = mysqlClient.Count(tablename)
	fmt.Println("count : ", count)
	// drop the table
	mysqlClient.DropTable(tablename)
	// check it
	if mysqlClient.IsTableExist(tablename) != false {
		t.Fatal(errors.New("error when drop the table"))
	}
	mysqlClient.Close()
}
