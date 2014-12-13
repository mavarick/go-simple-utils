package main

import (
    "lib/dbtools"
    "fmt"
    "log"
    "errors"
)

func main(){
    //host := "localhost"
    host := ""
    user := "root"
    password := ""
    charset := "utf8"
    db := "NameGender"

    mysqlClient := dbtools.NewMysqlClient(host, user, password, db, charset)
    count := mysqlClient.Count("url")
	if count != 827 {
		info := fmt.Sprintf("count is not right: %d", count)
		log.Fatal(errors.New(info))
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
			log.Fatal(err)
		}
		fmt.Println(fmt.Sprintf("%d:%s:%s:%d:%d", id, url, add_time, is_fetched, is_finished))
	}
	// create table
	q := "create table test_go (id int not null, name varchar(40))"
	mysqlClient.Exec(q)
	tablename := "test_go"
	isTableExist := mysqlClient.IsTableExist(tablename)
	if isTableExist != true {
		log.Fatal(errors.New("Error table test_go not exist"))
	}
	// insert into the table
	q = "insert into test_go values(1, 'lxf')"
	count = mysqlClient.Count(tablename)
	if count != 1 {
		log.Fatal(errors.New("Insert Error"))
	}
    //fmt.Println(count)
}
