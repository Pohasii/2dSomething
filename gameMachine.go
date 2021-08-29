package main

import (
	"time"
)

type gameServer struct {
	size    worldMapSize // struct
	players              // []players - struct
}

func initServer(x, y float32) gameServer {
	return gameServer{
		worldMapSize{
			width:  x,
			height: y,
		},
		make(players, 0, 6),
	}
}

func (s *gameServer) start(settings *Settings) {

	ticker := time.Tick(settings.getWorldCycleTimeTypeDuration() * time.Second)

	for range ticker {
		// fmt.Println(settings.name)
	}
}

type worldMapSize struct {
	width  float32
	height float32
}
