package main

import (
	"bufio"
	"fmt"
	flatbuffers "github.com/google/flatbuffers/go"
	"io"
	"log"
	"net"
	"os"
	"time"
)

var messageFromServer = make(chan []byte, 100)

func main() {

	fmt.Println("Hi I am a Client for tcp server")

	b := flatbuffers.NewBuilder(0)

	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8081")
	if err != nil {
		log.Println("setAddr: ", err)
	}

	conn, err := net.Dial(addr.Network(), addr.String())
	if err != nil {
		log.Println("check: ", err)
	}

	go readMess(&conn, messageFromServer)

	go func() {
		for {
			var messToServer string
			_, err := fmt.Fscan(os.Stdin, &messToServer)
			if err != nil {
				log.Println(err)
			}
			fmt.Println("message to Server: ", messToServer)

			mes := makeRootMessage(b, 2, []byte(messToServer))
			count, err := conn.Write(mes)
			if err != nil {
				log.Println("send to server fail: ", err)
			}
			fmt.Println("message to Server: ", mes, "count :", count)
			// messageForServer <- append([]byte(messToServer), newline[0])
		}
	}()

	var sentTime time.Time
	var Ping time.Duration
	ticker := time.NewTicker(2 * time.Second)

	for {
		select {
		case val, ok := <-messageFromServer:

			if ok {
				mType, _ := readRootMessage(val)

				if mType == 1 {
					Ping = time.Now().Sub(sentTime)
					fmt.Println("Ping: ", Ping)
				}

				//if mType == 1 {
				//	_, err := conn.Write(makeRootMessage(b, 1, []byte("")))
				//	if err != nil {
				//		return
				//	}
				//}
				fmt.Println("message form server: ", string(val))
			}
		case <-ticker.C:
			sentTime = time.Now()
			_, err := conn.Write(makeRootMessage(b, 1, []byte("")))
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("send ping test")
		}
	}

}

func readMess(Conn *net.Conn, c chan []byte) {

	// mess := make([]byte, 0, 10240)
	for {

		bufferBytes, err := bufio.NewReader(*Conn).ReadBytes('\n')
		if err != nil {
			log.Printf("error the the gameServer try to read from user: %v", err)
			if err == io.EOF {
				log.Printf("loss of connection of: %v", err)
				break
			}
			break
		} else {
			c <- bufferBytes
		}

		//count, err := (*Conn).Read(mess)
		//if err != nil {
		//	log.Println("Read: ", err)
		//	if err == io.EOF {
		//		err := (*Conn).Close()
		//		if err != nil {
		//			log.Fatal(err)
		//			return
		//		}
		//		break
		//	}
		//}
		//
		//fmt.Println(mess)
		//mess = mess[:count]
		//c <- mess
	}
}


func makeRootMessage(b *flatbuffers.Builder, mType int, data []byte) []byte {

	b.Reset()

	dataMessage := b.CreateByteVector(data)

	RootMessageStart(b)
	RootMessageAddMType(b, int32(mType))
	RootMessageAddContents(b, dataMessage)

	rootMessagePosition := RootMessageEnd(b)
	b.Finish(rootMessagePosition)

	return append(b.Bytes[b.Head():], []byte("\n")...)
}

func readRootMessage(buf []byte) (mType int32, data []byte) {
	mess := GetRootAsRootMessage(buf, 0)

	mType = mess.MType()
	data = mess.ContentsBytes()

	return
}