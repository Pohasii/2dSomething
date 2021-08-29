package main

import (
	"fmt"
)

var messageFrom chan messageFromTCPUser = make(chan messageFromTCPUser, 100)

func main() {

	Settings := initSettings("test", "localhost", "8081", 0, 3)
	Server := initServer(50, 50)
	Users := initUsers()
	NWs := initNetworkSocket(Settings.ip, Settings.port)

	fmt.Printf("Hello time machine gameServer %v :)\n", Settings.name)

	go Server.start(&Settings)

	NWs.handleConnection(&Users)
}

//1) игровой цикл
//2) игрок
//3) создать сокет
//4) роутер команд
//5) запусть это все)
