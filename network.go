package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

type networkSocket struct {
	port string
	ip   string
	l    *net.TCPListener
}

func initNetworkSocket(ip, port string) networkSocket {
	n := networkSocket{
		port: port,
		ip:   ip,
		l:    nil,
	}
	n.makeListener()

	return n
}

func (n *networkSocket) handleConnection(Users *users) {
	defer func(l *net.TCPListener) {
		err := l.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(n.l)

	for {
		c, err := n.l.AcceptTCP()
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("new connections with addr: ", c.RemoteAddr().String())
		Users.addUser(c.RemoteAddr().String(), c)
	}
}

func (n *networkSocket) makeListener() {
	l, err := net.ListenTCP("tcp", n.getTCPAddr())
	if err != nil {
		log.Fatal(err)
	}

	n.l = l
}

func (n *networkSocket) getAddr() string {
	return n.ip + ":" + n.port
}

func (n *networkSocket) getTCPAddr() *net.TCPAddr {
	tcpAddr, err := net.ResolveTCPAddr("tcp", n.getAddr())
	if err != nil {
		log.Fatal(err)
	}

	return tcpAddr
}

// =======================
// message
// =======================

type messageFromTCPUser struct {
	ip   string
	data []byte
	id   int
	date time.Time
}
