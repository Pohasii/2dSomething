package main

import (
	"fmt"
	"time"
)

func main() {

	Settings := initSettings("test", "localhost", "8081", 0, 3)
	Server := initServer(1, 1)
	Users := initUsers()
	NWs := initNetworkSocket(Settings.ip, Settings.port)

	fmt.Printf("Hello time machine server %v :)\n", Settings.name)

	go Server.start(&Settings)

	NWs.handleConnection(&Users)
}

type Settings struct {
	name           string
	id             int64
	ip             string
	port           string
	worldCycleTime int64
}

func initSettings(name, ip, port string, id, worldCycleTime int64) Settings {
	return Settings{
		name,
		id,
		ip,
		port,
		worldCycleTime,
	}
}

func (s Settings) getWorldCycleTimeTypeDuration() time.Duration {
	return time.Duration(s.worldCycleTime)
}

//1) игровой цикл
//2) игрок
//3) создать сокет
//4) роутер команд
//5) запусть это все)