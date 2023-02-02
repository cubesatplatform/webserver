package main

//go get github.com/lib/pq

import (
	"github.com/cubesatplatform/webserver/webserver/mydb"
	"github.com/cubesatplatform/webserver/webserver/myweb"
)

func main() {
	defer mydb.DB.Close()

	//myjson.TestJSON()
	//myjson.TestunJSON()
	myweb.ListenWeb()
}
