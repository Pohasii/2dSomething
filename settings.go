package main

import "time"

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
