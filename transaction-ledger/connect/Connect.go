package connect

import (
	"log"
    //  "github.com/gofiber/fiber/v2"
	"github.com/wagslane/go-rabbitmq"
	// "encoding/json"
	// "tr-ledger/"
	// "database/sql"
	"github.com/google/uuid"  
	"tr-ledger/config"
)

// type AllConection struct{
// 	RbMq *rabbitmq.Conn
// 	App *fiber.App
//     PostGresDB *sql.DB

// }
// var connections *AllConection
func ConnectMq(){
	conn, err := rabbitmq.NewConn(
		"amqp://guest:guest@localhost",
		rabbitmq.WithConnectionOptionsLogging,
	)
	if err != nil {
		log.Fatalln("error in connecting rabbitmq: ", err)
	}
	log.Println("con-----",conn)
	// config.connections=&config.AllConection{
	// 	RbMq: conn,
	// }
	config.SetConnection(&config.AllConection{
		RbMq: conn,
		PostGresDB: ConnectPostGres(),
	})
	log.Println("set connection")
	// connections.RbMq=conn
	

}
// func GetConnection()*AllConection{
// 	return connections

// }
type Transaction struct{
	Id             uuid.UUID `json:"id"`
    FromAccountId  uuid.UUID `json:"from_account_id"`
    // TransactionId  string    `json:"transaction_id"`  // <-- string, not UUID
    Amount         float64   `json:"amount"`
    ToAccountId    uuid.UUID `json:"to_account_id"`
    Status         string    `json:"status"`
    CreatedAt      int64     `json:"created_at"`
    UpdatedAt      int64     `json:"updated_at"`
}

// func Recv(){
// 	log.Println("in recv")
// 	consumer, err := rabbitmq.NewConsumer(
// 		connections.RbMq,
// 		"my_queue",
// 		rabbitmq.WithConsumerOptionsRoutingKey("transaction.created"),
// 		rabbitmq.WithConsumerOptionsExchangeName("transaction-exchange"),
// 		rabbitmq.WithConsumerOptionsExchangeDeclare,
// 		rabbitmq.WithConsumerOptionsExchangeKind("direct"), // same as Spring's DirectExchange
// 		rabbitmq.WithConsumerOptionsExchangeDurable,        // ✅ make it durable
// 		rabbitmq.WithConsumerOptionsQueueDurable,           // optional but recommended
		 
// 	)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	err = consumer.Run(func(d rabbitmq.Delivery) rabbitmq.Action {
// 		var t Transaction
// 		if err := json.Unmarshal(d.Body, &t); err != nil {
//             log.Println("Error decoding:", err)
        
//         }
// 		log.Println("consumed1: ",t)
// 		log.Printf("consumed: %v", string(d.Body))
// 		// rabbitmq.Ack, rabbitmq.NackDiscard, rabbitmq.NackRequeue
// 		// dB:=server.OpenMysqlConnectionQuery()
// 		//i have to credit to to account and debit from from account

// 		return rabbitmq.Ack
// 	})
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// }
