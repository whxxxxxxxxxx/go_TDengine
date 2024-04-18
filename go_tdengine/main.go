package main

import (
	"database/sql"
	"fmt"
	_ "github.com/taosdata/driver-go/v3/taosSql"
)

func main() {
	var taosUri = "root:taosdata@tcp(localhost:6030)/"
	taos, err := sql.Open("taosSql", taosUri)
	if err != nil {
		fmt.Println("failed to connect TDengine, err:", err)
		return
	}
	println(taos)
	defer taos.Close()

}
