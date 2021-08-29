package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"sync"
	"time"
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
	defer u.Unlock()
	u.i++

	newUser := user{
		IP:     ip,
		id:     u.i,
		status: true,
		Conn:   tcp,
	}

	u.u = append(u.u, newUser)

	go u.getUserById(newUser.id).reader()

}

func (u *users) getUserById(id int) *user {
	u.Lock()
	defer u.Unlock()

	for _, us := range u.u {
		if us.id != id {
			continue
		}
		return &us
	}
	return nil
}

// ============
// user
// ============

type user struct {
	IP     string
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
			log.Printf("error the the gameServer try to read from user: %v - %v", u, err)
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
				date: time.Now(),
			}
		}
	}
}

func (u *user) write(data []byte) {
	_, err := u.Conn.Write(data)
	if err != nil {
		log.Println(err)
	}
}
