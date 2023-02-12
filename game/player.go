package game

type player struct {
	hand          []card
	score         int
	wannaTakeMore bool
	isComputer    bool
}

func NewPlayer(isComputer bool) *player {
	return &player{
		hand:          make([]card, 0, 21),
		score:         0,
		wannaTakeMore: true,
		isComputer:    isComputer,
	}
}

func (p *player) cleanHand() {
	p.hand = make([]card, 0, 21)
}

func (p *player) getHandScore() int {
	handScore := 0
	for _, handCard := range p.hand {
		if handCard.value == 11 && handScore+handCard.value > 21 {
			handScore++
		} else {
			handScore += handCard.value
		}
	}
	return handScore
}
