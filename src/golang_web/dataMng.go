package main

import (
	"github.com/ziutek/mymysql/mysql"
	_ "github.com/ziutek/mymysql/thrsafe"
	"log"
)

type DataMng struct {
	mysql.Conn

}

func NewDataMng() *DataMng {
	conn:=mysql.New("tcp","","127.0.0.1:3306","root","root","golang_web")
	err:=conn.Connect()
	if err!=nil {
		log.Println(err)
	}
	return &DataMng{conn}
}
