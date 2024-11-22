package main

import (
	"fmt"

	"github.com/cengsin/oracle"
	"gorm.io/gorm"
)

func main() {
	urlOptions := map[string]string{
		"SID": "xe",
	}
	connStr := go_ora.BuildUrl("", 1521, "", "SYSTEM", "root", urlOptions)
	// conn, err := sql.Open("oracle", connStr)
	oracle.Open("system/oracle@127.0.0.1:1521/XE")

	db, err := gorm.Open(oracle.Open("system/oracle@localhost:1521/XE"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if err != nil {
		fmt.Println("Can't open the driver: ", err)
		return
	}

	defer func() {
		err = conn.Close()
		if err != nil {
			fmt.Println("Can't close connection: ", err)
		}
	}()

	err = conn.Ping()
	if err != nil {
		fmt.Println("Can't ping connection: ", err)
		return
	}

	rows, err := conn.Query("SELECT 'Hello World! 123' FROM dual")
	if err != nil {
		fmt.Println("Erro ao fazer select: ", err)
		return
	}

	defer func() {
		err = rows.Close()
		if err != nil {
			fmt.Println("Can't close dataset: ", err)
		}
	}()

	var (
		name string
	)
	for rows.Next() {
		err = rows.Scan(&name)
		fmt.Println("\tName: ", name)
	}
}
