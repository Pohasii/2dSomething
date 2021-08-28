package main

import (
	"fmt"
	"time"
)

type server struct {
	size int // struct
	players int // []players - struct
}

type playersMap struct {
	width float32
	height float32
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

