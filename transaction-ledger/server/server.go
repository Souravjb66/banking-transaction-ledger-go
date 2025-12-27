package server

import (
	"log"
	// "tr-ledger/connect"
	// sqlcdb "tr-ledger/db"
	"tr-ledger/config"
	
	
)



func ServerSetup(){
    // connect.ConnectPostGres()
	if err:=config.GetConnection().App.Listen(":8082");err!=nil{
		log.Fatal(err)
	}

}


// func OpenMysqlConnectionQuery()*sqlcdb.Queries{
// 	db:=sqlcdb.New(connect.GetConnection().PostGresDB.PostGres)
// 	return db

// }