package connect

import (
	"database/sql"
	"fmt"
	"log"
	_"github.com/lib/pq"
	sqlcdb "tr-ledger/db"
	"tr-ledger/config"
	
)


// type MySqlDbInstance struct {
// 	PostGres *sql.DB
// }
type MySqlRequired struct {
	Url      string
	UserName string
	Password string
	Port     string
	DbName   string
}

// var MySqlInstance MySqlDbInstance

func ConnectPostGres() *sql.DB{
	// "root:@tcp(127.0.0.1:3306)/mydb"
	// dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true", "postgres", "root", "localhost", 5432, "bank_test")
	dsn := "postgres://postgres:root@localhost:5432/bank_test?sslmode=disable"
	
    
	
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	// Test the connection
	err = db.Ping()
	if err != nil {
		log.Fatal("Error pinging the database:", err)
	}

	fmt.Println("Successfully connected to Postgres!")
	// config.GetConnection().PostGresDB=db
	return db
	// return &MySqlDbInstance{PostGres: db}

}

func OpenMysqlConnectionQuery()*sqlcdb.Queries{
	db:=sqlcdb.New(config.GetConnection().PostGresDB)
	return db

}

