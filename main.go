package main

import (
	"migrationScript/common"
	"migrationScript/migrations"
)

func main() {
	common.ConnectToDatabase()
	err := migrations.Run()
	if err != nil {
		println("Error: " + err.Error())
	} else {
		println("Ok")
	}
}
