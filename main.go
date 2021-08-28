package main

import (
	"fmt"
	"time"
)

func main() {

	Settings := initSettings("test", "0.0.0.0", 1, 3)
	Server := initServer(1,1)

	fmt.Printf("Hello time machine server %v :)\n", Settings.name)

	go Server.start(&Settings)

	for {

	}
}

type server struct {
	size int // struct
	players int // []players - struct
}

func initServer (size int, players int) server {
	return struct {
		size    int
		players int
	}{size, players}
}

func (s *server) start (settings *Settings) {

	ticker := time.Tick(settings.getWorldCycleTimeTypeDuration() * time.Second)

	for range ticker {
		fmt.Println(settings.name)
	}
}

type Settings struct {
	name string
	id int64
	ip string
	worldCycleTime int64
	users []users
}

func initSettings (name, ip string, id, worldCycleTime int64) Settings {
	return Settings{
		name,
		id,
		ip,
		worldCycleTime,
		make([]users, 0, 5),
	}
}

func (s Settings) getWorldCycleTimeTypeDuration() time.Duration {
	return time.Duration(s.worldCycleTime)
}

type users struct {
	token string
	IPs []string
	id int
}

func (u users) init (token string, ip string, id int) users{
	return users{
		token,
		[]string{ip},
		id,
	}
}
//1) игровой цикл
//2) игрок
//3) создать сокет
//4) роутер команд
//5) запусть это все)