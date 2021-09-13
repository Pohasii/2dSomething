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
	toGMs   chan messageForGameMachine
	toUsers chan toUsersChan
}

func initRouter(users *users,
	players *players,
	fb *flatbuffers.Builder,
	from chan messageFromTCPUser,
	toGMs chan messageForGameMachine,
	toUsers chan toUsersChan,
	) Router {
	return Router{
		users:   users,
		players: players,
		fb:      fb,
		from:    from,
		toGMs:   toGMs,
		toUsers: toUsers,
	}
}

func (r *Router) start() {

	go func() {
		for m := range r.toUsers {
			u := r.users.getUserById(int(m.userID))
			if u != nil {
				u.write(m.data)
			}
		}
	}()
	for m := range r.from {
		mType, data := readRootMessage(m.data)
		if mType == 1 {
			if us := r.users.getUserById(m.id); us != nil {
				us.write(makeRootMessage(r.fb, 1, []byte("1")))
				fmt.Println("ping from: ", m.ip)
			}
		}
		if mType == 2 {
			if !r.players.CheckThePlayerById(ID(m.id)) {
				r.players.addPlayer(ID(m.id), 0, 0, 5)
			}
		}

		if mType == 3 { //up
			r.toGMs <- readGMMessage(data)
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

func makeGMMessage(b *flatbuffers.Builder, mType int, data []byte) []byte {

	b.Reset()

	dataMessage := b.CreateByteVector(data)

	mesSchem.GameMachineMessageStart(b)
	mesSchem.GameMachineMessageAddMType(b, int32(mType))
	mesSchem.GameMachineMessageAddContents(b, dataMessage)

	MessagePosition := mesSchem.GameMachineMessageEnd(b)
	b.Finish(MessagePosition)

	return append(b.Bytes[b.Head():], []byte("\n")...)
}

func readGMMessage(buf []byte) messageForGameMachine {
	mess := mesSchem.GetRootAsGameMachineMessage(buf, 0)

	mType := int(mess.MType())
	data := mess.ContentsBytes()

	return messageForGameMachine{
		mType: mType,
		data:  data,
	}
}

type toUsersChan struct {
	userID ID
	data []byte
}

//
//type messageForGameMachine struct {
//	mType int
//	data []byte
//}
//
//func makeMoveMessage(b *flatbuffers.Builder, data struct{ID int; x float32; y float32}) []byte {
//
//	b.Reset()
//
//	mesSchem.MoveMessageStart(b)
//	mesSchem.MoveMessageAddId(b, uint32(data.ID))
//	mesSchem.MoveMessageAddX(b, data.x)
//	mesSchem.MoveMessageAddY(b, data.y)
//
//	MessagePosition := mesSchem.MoveMessageEnd(b)
//	b.Finish(MessagePosition)
//
//	return append(b.Bytes[b.Head():], []byte("\n")...)
//}
//
//func readMoveMessage(buf []byte) (ID int, x,y float32) {
//	mess := mesSchem.GetRootAsMoveMessage(buf, 0)
//
//	ID = int(mess.Id())
//	x = mess.X()
//	y = mess.Y()
//
//	return
//}
