package main

import (
	"net"
	"sync"
)

type users struct {
	sync.Mutex
	i int
	u []user
}

func initUsers() users {
	return users{
		i: 0,
		u: make([]user, 0, 5),
	}
}

func (u *users) addUser(ip string, tcp *net.TCPConn) {
	u.Lock()
	newUser := user{
		[]string{ip},
		u.i,
		true,
		tcp,
	}

	u.i++

	u.u = append(u.u, newUser)
	u.Unlock()
}

type user struct {
	IPs    []string
	id     int
	status bool
	Conn   *net.TCPConn
}
