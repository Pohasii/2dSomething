package main

import (
	"bufio"
	"io"
	"log"
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
		IP:ip,
		id: u.i,
		status: true,
		Conn: tcp,
	}

	u.i++

	u.u = append(u.u, newUser)
	u.Unlock()
}

type user struct {
	IP    string
	id     int
	status bool
	Conn   *net.TCPConn
}

func (u *user) reader() {
	defer func(Conn *net.TCPConn) {
		err := Conn.Close()
		if err != nil {
			log.Println("Close the Connection for user: ", u.id, err)
		}
	}(u.Conn)

	// _ = u.Conn.SetReadDeadline(time.Now().Add(time.Minute))

	for {

		bufferBytes, err := bufio.NewReader(u.Conn).ReadBytes('\n')
		if err != nil {
			log.Printf("error the the gameServer try to read from user: %v - %v", u ,err)
			one := make([]byte, 1)
			if _, err := u.Conn.Read(one); err == io.EOF {
				log.Printf("loss of connection of %v: %v", u, err)
				break
			}
			break
		} else {
			messageFrom <- messageFromTCPUser{
				ip:   u.Conn.RemoteAddr().String(),
				data: bufferBytes,
				id:   u.id,
			}
		}
	}
}

func (u *user) write (data []byte) {
	_, err := u.Conn.Write(data)
	if err != nil {
		log.Println(err)
	}
}
