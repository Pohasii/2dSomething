package main

import (
	"fmt"
	flatbuffers "github.com/google/flatbuffers/go"
	mesSchem "github.com/pohasii/2dsomething"
	// messages "mod"
)

type Router struct {
	users   *users
	players *players
	fb      *flatbuffers.Builder
	from    chan messageFromTCPUser
}

func initRouter(users *users, players *players, fb *flatbuffers.Builder, from chan messageFromTCPUser) Router {
	return Router{
		users:   users,
		players: players,
		fb:      fb,
		from:    from,
	}
}

func (r *Router) start() {

	for m := range r.from {
		t, _ := readRootMessage(m.data)
		if t == 1 {
			if us := r.users.getUserById(m.id); us != nil {
				us.write(makeRootMessage(r.fb, 1, []byte("1")))
				fmt.Println("ping from: ", m.ip)
			}
		}
		if t == 2 {
			if !r.players.CheckThePlayerById(m.id) {
				r.players.addPlayer(m.id, 0,0,5)
			}
		}


		if t == 3 { //up
			fmt.Println("up")
			if r.players.CheckThePlayerById(m.id) {
				p := r.players.findPlayerById(m.id)
				x,y := p.getPos()
				fmt.Println("x, y : ", x, y)
				y++
				fmt.Println("x, y : ", x, y)
				fmt.Println("x, y : ", x, y+1)
				(*p).updatePos(x,y)
			}
		}
		if t == 4 { // down
			fmt.Println("d")
			if r.players.CheckThePlayerById(m.id) {
				p := r.players.findPlayerById(m.id)
				x,y := p.getPos()
				(*p).updatePos(x,y-1)
			}
		}
		if t == 5 { // left
			fmt.Println("l")
			if r.players.CheckThePlayerById(m.id) {
				p := r.players.findPlayerById(m.id)
				x,y := p.getPos()
				p.updatePos(x+1,y)
			}
		}
		if t == 6 { // right
			fmt.Println("r")
			if r.players.CheckThePlayerById(m.id) {
				p := r.players.findPlayerById(m.id)
				x,y := p.getPos()
				p.updatePos(x-1,y)
			}
		}
	}
}


// ================
// message type

func makeRootMessage(b *flatbuffers.Builder, mType int, data []byte) []byte {

	b.Reset()

	dataMessage := b.CreateByteVector(data)

	mesSchem.RootMessageStart(b)
	mesSchem.RootMessageAddMType(b, int32(mType))
	mesSchem.RootMessageAddContents(b, dataMessage)

	rootMessagePosition := mesSchem.RootMessageEnd(b)
	b.Finish(rootMessagePosition)

	return append(b.Bytes[b.Head():], []byte("\n")...)
}

func readRootMessage(buf []byte) (mType int, data []byte) {
	mess := mesSchem.GetRootAsRootMessage(buf, 0)

	mType = int(mess.MType())
	data = mess.ContentsBytes()

	return
}
