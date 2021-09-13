package main

import (
	flatbuffers "github.com/google/flatbuffers/go"
	mesSchem "github.com/pohasii/2dsomething"
	"sync"
	"time"
)

type gameServer struct {
	sync.Mutex
	size    worldMapSize // struct
	players              // []players - struct
	//forRead chan int
	tasks
}

func initServer(x, y float32) gameServer {
	return gameServer{
		size: worldMapSize{
			width:  x,
			height: y,
		},
		players: make(players, 0, 6),
		//forRead: make(chan int, 5),
		tasks: make([]task, 0, 20),
	}
}

func (s *gameServer) start(settings *Settings, toGameServer <-chan messageForGameMachine, fromGMs chan toUsersChan) {

	//wg := &sync.WaitGroup{}

	ticker := time.Tick(settings.getWorldCycleTimeTypeDuration() * time.Second)

	b:= flatbuffers.NewBuilder(0) // remove

	go taskMaker(s, toGameServer)

	for range ticker {
		for _, t := range s.tasks {
			//wg.Add(1)

			s.Lock()
			t.run(s)
			s.Unlock()
			//go func() {
			//	s.Lock()
			//	t.run(s)
			//	s.Unlock()
			//	//wg.Done()
			//}()
		}
		//wg.Wait()
		s.Lock() // remove
		for _, p := range s.players {
			id := p.getId()
			x,y:= p.getPos()
			PlayerState := makeMoveMessage(b, struct {
				id ID
				x  float32
				y  float32
			}{id: id, x: x, y:y}) // remove
			messFromGMs := makeGMMessage(b, 0, PlayerState)
			fromGMs <- toUsersChan{
				userID: id,
				data:   makeRootMessage(b,3,messFromGMs),
			}
		}

		s.Unlock() // remove
	}
}

type worldMapSize struct {
	width  float32
	height float32
}

type tasks []task

type task interface {
	run(gs *gameServer) error //waiter *sync.WaitGroup
}

func taskMaker(gs *gameServer, messageFromRouter <-chan messageForGameMachine) {
	for mes := range messageFromRouter {
		switch mes.mType {
		case 1:
			go func() {
				defer gs.Unlock()
				gs.Lock()
				id, x,y := readMoveMessage(mes.data)
				gs.tasks = append(gs.tasks, move{
					id,
					x,
					y,
				})
			}()
		}
	}
}


//
//type taskMaker struct {
//	taskType
//	taskStruct map[int]struct{}
//}
//type taskType map[int]getData
//
//type getData interface {
//	get() interface{}
//}
//
//type test struct{
//	id ID
//	x,y float32
//}
//
//func (tm *taskMaker) run(gs *gameServer, messageFromRouter <-chan messageForGameMachine) {
//	for mes := range messageFromRouter {
//		data := tm.taskType[mes.mType].get().(test)
//
//		go func() {
//			defer gs.Unlock()
//			gs.Lock()
//			id, x,y := readMoveMessage(mes.data)
//			gs.tasks = append(gs.tasks, move{
//				id,
//				x,
//				y,
//			})
//		}()
//	}
//}


// messages from Router
//func makeGMMessage(b *flatbuffers.Builder, mType int, data []byte) []byte {
//
//	b.Reset()
//
//	dataMessage := b.CreateByteVector(data)
//
//	mesSchem.GameMachineMessageStart(b)
//	mesSchem.GameMachineMessageAddMType(b, int32(mType))
//	mesSchem.GameMachineMessageAddContents(b, dataMessage)
//
//	MessagePosition := mesSchem.GameMachineMessageEnd(b)
//	b.Finish(MessagePosition)
//
//	return append(b.Bytes[b.Head():], []byte("\n")...)
//}
//
//func readGMMessage(buf []byte) messageForGameMachine {
//	mess := mesSchem.GetRootAsGameMachineMessage(buf, 0)
//
//	mType := int(mess.MType())
//	data := mess.ContentsBytes()
//
//	return messageForGameMachine{
//		mType: mType,
//		data:  data,
//	}
//}


func makeMoveMessage(b *flatbuffers.Builder, data struct{id ID; x float32; y float32}) []byte {

	b.Reset()
	mesSchem.MoveMessageStart(b)
	mesSchem.MoveMessageAddId(b, uint32(data.id))
	mesSchem.MoveMessageAddX(b, data.x)
	mesSchem.MoveMessageAddY(b, data.y)

	MessagePosition := mesSchem.MoveMessageEnd(b)
	b.Finish(MessagePosition)

	return append(b.Bytes[b.Head():], []byte("\n")...)
}

func readMoveMessage(buf []byte) (id ID, x,y float32) {
	mess := mesSchem.GetRootAsMoveMessage(buf, 0)
	id = ID(mess.Id())
	x = mess.X()
	y = mess.Y()

	return
}

type messageForGameMachine struct {
	mType int
	data []byte
}





/// tasks

type move struct {
	id   ID
	dirX float32
	dirY float32
}

func (m move) run(gs *gameServer) error {
	p := gs.players.findPlayerById(m.id)
	if p != nil {
		p.updatePos(m.dirX,m.dirY)
	}
	//errors.New()
	return nil
}

/// actions game server