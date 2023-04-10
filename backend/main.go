package main

import (
	"database/sql"
	"log"
	"github.com/vutranhoang1411/tmdt/api"
	"github.com/vutranhoang1411/tmdt/util"
	_ "github.com/lib/pq"
)

func main(){
	config,err:=util.LoadConfig(".")
	if err!=nil{
		log.Fatal(err)
	}
	conn,err:=sql.Open(config.DBDriver,config.DBSource)
	if err!=nil{
		log.Fatal(err)
	}

	server:=api.NewServer(conn)
	server.Start(config.HttpServerAddress)
}