package main

import (
	"fmt"
	flatbuffers "github.com/google/flatbuffers/go"
)

//var messageFrom chan messageFromTCPUser = getData(chan messageFromTCPUser, 100)

func main() {

	// start settings
	Settings := initSettings("test", "localhost", "8081", 0, 3)

	// game machine server
	Server := initServer(50, 50)
	messageToGMs := make(chan messageForGameMachine, 100)

	// network
	messageFrom := make(chan messageFromTCPUser, 100)
	Users := initUsers(messageFrom)
	messageForUsers := make(chan toUsersChan)
	NWs := initNetworkSocket(Settings.ip, Settings.port)
	networkMessageRouter := initRouter(&Users,
		&Server.players,
		flatbuffers.NewBuilder(0),
		messageFrom,
		messageToGMs,
		messageForUsers)

	// server hello message
	fmt.Printf("Hello time machine gameServer %v :)\n", Settings.name)

	// game machine cycle start
	go Server.start(&Settings, messageToGMs, messageForUsers)

	// routing messages
	go networkMessageRouter.start()

	// connections accept
	NWs.handleConnection(&Users)
}

//1) игровой цикл
//2) игрок
//3) создать сокет
//4) роутер команд
//5) запусть это все)
