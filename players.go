package main

type players []player 

func (p players) findPlayerById(id uint) player {
	for _, Player := range p {
		if Player.getId() != id { continue }
		return Player
	}

	return player{
		id:       -1,
		position: pos{},
	}
}

func (p *players) addPlayer(id uint, x float32, y float32) {
	Player := player{id, pos{x,y}}
	*p = append(*p, Player)
}

//remove 
func (p *players) removePlayer(id uint) {
	Player := p.findPlayerById(id)
	if Player.getId() != -1 {
		// p = remove(p, Player)
	}
}

// PLAYER PART

type pos struct {
	x float32
	y float32
}

type player struct {
	id uint
	position pos
}

func (p player) getPos() (x, y float32){
	return p.position.x, p.position.y
}

// will it work?)
func (p *player) updatePos(x, y float32) {
	p.position.x = x
	p.position.y = y
}

func (p player) getId() uint {
	return p.id
}
