package service

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	// "github.com/wagslane/go-rabbitmq"
	"encoding/json"
	"tr-ledger/connect"
	"tr-ledger/db"

	// "github.com/google/uuid"
	"github.com/wagslane/go-rabbitmq"
	"strconv"
	"sync"
	"tr-ledger/config"
	
)

func Transaction(ctx *fiber.Ctx) error {
	// go Recv()
	return nil

}
func Recv(wg *sync.WaitGroup) {
	log.Println("in recv")
	defer wg.Done()
	consumer, err := rabbitmq.NewConsumer(
		config.GetConnection().RbMq,
		"my_queue",
		rabbitmq.WithConsumerOptionsRoutingKey("transaction.created"),
		rabbitmq.WithConsumerOptionsExchangeName("transaction-exchange"),
		rabbitmq.WithConsumerOptionsExchangeDeclare,
		rabbitmq.WithConsumerOptionsExchangeKind("direct"), // same as Spring's DirectExchange
		rabbitmq.WithConsumerOptionsExchangeDurable,        //  make it durable
		rabbitmq.WithConsumerOptionsQueueDurable,           // optional but recommended

	)
	if err != nil {
		log.Fatal(err)
	}

	maxRetries:=3	
	err = consumer.Run(func(d rabbitmq.Delivery) rabbitmq.Action {
		var t connect.Transaction
		if err := json.Unmarshal(d.Body, &t); err != nil {
	        log.Println("Error decoding:", err)

	    }
		retryCount := 0
		if v, ok := d.Headers["x-retry-count"]; ok {
            retryCount = v.(int)
        }
		log.Println("consumed1: ",t)
		log.Printf("consumed: %v", string(d.Body))
		// rabbitmq.Ack, rabbitmq.NackDiscard, rabbitmq.NackRequeue
		dB:=connect.OpenMysqlConnectionQuery()
        // log.Println("after db openmy----------------")
		trParams:=db.UpdateTransactionStatusParams{
			ID: t.Id,
			Status: "FAILED",
		}
		if retryCount >= maxRetries {
			if err:=dB.UpdateTransactionStatus(context.Background(),trParams);err!=nil{
				log.Println(err)
				return rabbitmq.Ack
			}
            return rabbitmq.Ack // goes to DLQ or dropped
        }
		tx,err:=config.GetConnection().PostGresDB.BeginTx(context.Background(),nil)
		if err!=nil{
			log.Println("in create tx: ",err)
			log.Fatal(err)
			
		}
		// log.Println("after tx create----------------")
		qtx:=dB.WithTx(tx)
		sender,err:=qtx.GetAccountById(context.Background(),t.FromAccountId)
		// sender,err:=dB.GetAccountById(context.Background(),t.From_account_id)
		if err!=nil{
			log.Println("fetching sender: ",err)
			// log.Fatal(err)
			d.Headers["x-retry-count"] = retryCount + 1
			tx.Rollback()
			return rabbitmq.NackRequeue   // retry
		}
		// log.Println("after query sender----------------")
		param:=db.UpdateBalanceParams{
			ID: sender.ID,
			Balance: func()string{
	            // sB,err:=strconv.Atoi(sender.Balance)
				f, err := strconv.ParseFloat(sender.Balance, 64)
				if err!=nil{
					tx.Rollback()
				}
				left:=f-t.Amount
				floatStr2 := strconv.FormatFloat(left, 'f', -1, 64)
				return floatStr2


			}(),

		}
		err=qtx.UpdateBalance(context.Background(),param)
		if err!=nil{
			log.Println("updating sender balance: ",err)
			// log.Println(err)
			d.Headers["x-retry-count"] = retryCount + 1
			tx.Rollback()
			return rabbitmq.NackRequeue   // retry

		}
		// dB.UpdateBalance(context.Background(),)
		receiver,err:=qtx.GetAccountById(context.Background(),t.ToAccountId)
		if err!=nil{
			log.Println("fetching receiver: ",err)
			// log.Println(err)
			d.Headers["x-retry-count"] = retryCount + 1
			tx.Rollback()
			return rabbitmq.NackRequeue   // retry

		}
		rcParam:=db.UpdateBalanceParams{
			ID: receiver.ID,
			Balance: func()string{
	            
				f, err := strconv.ParseFloat(receiver.Balance, 64)
				if err!=nil{
					tx.Rollback()
				}
				total:=f+t.Amount
				floatStr3 := strconv.FormatFloat(total, 'f', -1, 64)
				return floatStr3

			}(),

		}
		err=qtx.UpdateBalance(context.Background(),rcParam)
		if err!=nil{
			log.Println("updating receiver balance: ",err)
			d.Headers["x-retry-count"] = retryCount + 1
			tx.Rollback()
			return rabbitmq.NackRequeue   // retry
		}
		//i have to credit to to account and debit from from account
		par:=db.UpdateTransactionStatusParams{
			Status :"SUCCESS",
            ID : t.Id,
		}
        err=qtx.UpdateTransactionStatus(context.Background(),par)
		if err!=nil{
			log.Println("updating transaction status: ",err)
			d.Headers["x-retry-count"] = retryCount + 1
			tx.Rollback()
			return rabbitmq.NackRequeue   // retry

		}
		if err := tx.Commit(); err != nil {
            log.Println("commit failed:", err)
			d.Headers["x-retry-count"] = retryCount + 1
			tx.Rollback()
            return rabbitmq.NackRequeue   // retry
        }
		return rabbitmq.Ack
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("end--------------------------------------------rab")
	

}
