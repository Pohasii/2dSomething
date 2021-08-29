package main

import (
	flatbuffers "github.com/google/flatbuffers/go"
	mesSchem "github.com/pohasii/2dsomething"
	// messages "mod"
)

type Router struct {
	users   *users
	players *players
	fb      *flatbuffers.Builder
	from    *chan messageFromTCPUser
}

func initRouter(users *users, players *players, fb *flatbuffers.Builder, from *chan messageFromTCPUser) Router {
	return Router{
		users:   users,
		players: players,
		fb:      fb,
		from:    from,
	}
}

func (r *Router) start() {

	for m := range *r.from {
		t, _ := readRootMessage(m.data)
		if t == 1 {
			if us := r.users.getUserById(m.id); us != nil {
				us.write(makeRootMessage(r.fb, 1, []byte("1")))
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
