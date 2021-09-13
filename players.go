package main

type players []*player

type ID uint

func (p players) findPlayerById(id ID) *player {
	for _, Player := range p {
		if Player.getId() != id {
			continue
		}
		return Player
	}

	return nil
}

func (p *players) addPlayer(id ID, x float32, y float32, speed float32) {
	Player := player{
		id:       id,
		position: pos{x: x, y: y},
		speed:    speed,
	}
	*p = append(*p, &Player)
}

func (p *players) CheckThePlayerById(id ID) bool {
	for _, player := range *p{
		if player.id == id {
			return true
		}
	}

	return false
}

//remove
func (p *players) removePlayer(id ID) {
	Player := p.findPlayerById(id)
	if Player != nil {
		// p = remove(p, Player)
	}
}

// PLAYER PART

type pos struct {
	x float32
	y float32
}

type player struct {
	id       ID
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

func (p player) getId() ID {
	return p.id
}
