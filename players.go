package main

type players []player 

func (p players) findPlayerById(id uint) Player player {
	for Player := range players {
		if Player.getId() != id { continue }
		return Player
	}
}

func (p *players) addPlayer(id uint, x float32, y float32) {
	Player = player{id, pos{x,y}}
	p = append(p, Player)
}

//remove 
func (p *players) removePlayer(id uint) {
	Player = findPlayerById(id)
	p = remove(p, Player)
}

// PLAYER PART

type pos struct {
	x: float32
	y: float32
}

type player struct {
	id: uint
	position: pos
}

func (p player) getPos() (x, y float32){
	return pos.x, pos.y
}

// will it work?)
func (p *player) updatePos(x, y float32) {
	p.pos.x = x
	p.pos.y = y
}

func (p player) getId() id uint {
	return p.id
}
