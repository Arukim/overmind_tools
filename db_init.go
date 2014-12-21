package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
)

func getValue(key string) string {
	var buffer bytes.Buffer
	for j := 0; j < 10000; j++ {
		buffer.WriteString(key)
	}
	return buffer.String()
}

func main() {

	connectString = flag.String("cs", "", "connection string")

	flag.Parse()

	db, err := sql.Open("mysql", connectString)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if _, err = db.Exec("drop table if exists `cache`"); err != nil {
		panic(err)
	}

	create := "create table if not exists `cache`(`id` integer auto_increment,`key` varchar(128) unique,`value` text,primary key(id))"

	if _, err = db.Exec(create); err != nil {
		panic(err)
	}

	for i := 0; i < 10000; i++ {

		key := strconv.Itoa(i)
		var valueInterface interface{}
		valueInterface = getValue(key)

		buffer := new(bytes.Buffer)
		enc := json.NewEncoder(buffer)
		enc.Encode(valueInterface)

		insert := fmt.Sprintf("insert into `cache` (`key`, `value`) values ('%s','%s')", key, buffer.String())
		fmt.Println("inserting " + key)
		if _, err = db.Exec(insert); err != nil {
			panic(err)
		}
	}
}
