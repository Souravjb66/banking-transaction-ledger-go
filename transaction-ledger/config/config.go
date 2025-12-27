package config

import (
	// "log"
     "github.com/gofiber/fiber/v2"
	"github.com/wagslane/go-rabbitmq"
	// "encoding/json"
	// "tr-ledger/"
	"database/sql"
	// "github.com/google/uuid"  
	// sqlcdb "tr-ledger/db"
	
)

type AllConection struct{
	RbMq *rabbitmq.Conn
	App *fiber.App
    PostGresDB *sql.DB

}
var connections *AllConection

func GetConnection()*AllConection{
	return connections

}
func SetConnection(con *AllConection){
	connections=con

}

