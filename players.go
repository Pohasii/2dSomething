package main

type players []player

func (p players) findPlayerById(id int) player {
	for _, Player := range p {
		if Player.getId() != id {
			continue
		}
		return Player
	}

	return player{
		id:       -1,
		position: pos{},
	}
}

func (p *players) addPlayer(id int, x float32, y float32, speed float32) {
	Player := player{
		id:       id,
		position: pos{x: x, y: y},
		speed:    speed,
	}
	*p = append(*p, Player)
}

//remove
func (p *players) removePlayer(id int) {
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
	id       int
	position pos
	speed    float32
}

func (p player) getPos() (x, y float32) {
	return p.position.x, p.position.y
}

// will it work?)
func (p *player) updatePos(x, y float32) {
	p.position.x = x
	p.position.y = y
}

func (p player) getId() int {
	return p.id
}
