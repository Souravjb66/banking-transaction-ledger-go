package main

import(
	// "github.com/gofiber/fiber/v2"
	"tr-ledger/server"
	"tr-ledger/connect"
	"tr-ledger/routes"
	"tr-ledger/service"
	"sync"
)

func main(){
	var wg sync.WaitGroup
	wg.Add(1)
	connect.ConnectMq()   //rmq set object in total
	routes.Routes()
	go service.Recv(&wg)
	server.ServerSetup()
	wg.Wait()





}