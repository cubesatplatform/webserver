package main

//go get github.com/lib/pq

import (
	"webserver/mydb"
	"webserver/myweb"
)

func main() {
	defer mydb.DB.Close()

	//myjson.TestJSON()
	//myjson.TestunJSON()
	myweb.ListenWeb()
}
