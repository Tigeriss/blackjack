package game

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"time"
)

type Game struct {
	players []*player
	deck    deck
	round   int
}

func NewGame() *Game {
	pl := make([]*player, 2, 2)
	pl[0] = NewPlayer(false)
	pl[1] = NewPlayer(true)
	d := NewDeck()
	d.shuffle()

	return &Game{
		players: pl,
		deck:    d,
		round:   1,
	}
}

func (g *Game) StartNewRound() {
	for _, pl := range g.players {
		pl.cleanHand()
		pl.wannaTakeMore = true
	}
	fmt.Printf("===========================\n")
	fmt.Printf("Round %d starting.\n", g.round)
	fmt.Printf("===========================\n\n")
}

func (g *Game) FinishRound() {
	fmt.Printf("===========================\n")
	fmt.Printf("Round %d finished.\n", g.round)
	fmt.Printf("===========================\n\n")
	fmt.Printf("Current score(user/computer):\n")
	fmt.Printf("%d : %d\n\n", g.players[0].score, g.players[1].score)
	fmt.Printf("===========================\n")
	g.round++
}

func (g *Game) GameOver() {
	fmt.Printf("===========================\n")
	fmt.Printf("Game is finished.\n")
	fmt.Printf("===========================\n\n")
	fmt.Printf("Current score(user/computer):\n")
	fmt.Printf("%d : %d\n\n", g.players[0].score, g.players[1].score)
	pScore := g.players[0].score
	cScore := g.players[1].score
	if pScore > cScore {
		fmt.Printf("Player wins!\n")
	} else if pScore == cScore {
		fmt.Printf("Player and Computer ties. No one wins.\n")
	} else {
		fmt.Printf("Computer wins!\n")
	}
	fmt.Printf("===========================\n")
}

func (g *Game) Round() {
	for g.players[0].wannaTakeMore {
		c := g.deck.dealCard()
		fmt.Printf("Your got yourself %s.\n", c.imgSrc)
		g.players[0].hand = append(g.players[0].hand, c)
		handScore := g.players[0].getHandScore()
		if handScore > 21 {
			fmt.Printf("Your score is: %d.It's more than 21, you loose this round\n", handScore)
			g.players[0].wannaTakeMore = false
			g.players[1].score++
			return
		} else {
			fmt.Printf("Your score is: %d.Do you need one more card?\n", handScore)
			fmt.Scanf("%t\n", &g.players[0].wannaTakeMore)
		}

	}
	for g.players[1].wannaTakeMore {
		c := g.deck.dealCard()
		fmt.Printf("Computer got himself %s.\n", c.imgSrc)
		g.players[1].hand = append(g.players[1].hand, c)
		fmt.Printf("Computer's score is: %d.\n", g.players[1].getHandScore())
		if g.players[1].getHandScore() > 21 {
			fmt.Printf("It's more, than 21. Computer looses this round\n")
			g.players[1].wannaTakeMore = false
			g.players[0].score++
			return
		} else if g.players[1].getHandScore() > 18 {
			fmt.Printf("Computer thinks it's enough.\n")
			g.players[1].wannaTakeMore = false
		} else {
			fmt.Printf("Computer would take one more card.\n")
		}
		time.Sleep(1 * time.Second)
	}
	pScore := g.players[0].getHandScore()
	cScore := g.players[1].getHandScore()
	fmt.Printf("Player's score is %d and Computer's score is %d\n", pScore, cScore)
	if pScore > cScore {
		fmt.Printf("Player wins this round\n")
		g.players[0].score++
	} else if pScore == cScore {
		fmt.Printf("Player and Computer ties. No one wins this round. Score stays unchanged.\n")
	} else {
		fmt.Printf("Computer wins this round\n")
		g.players[1].score++
	}
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{
		R: 0xa8,
		G: 0xc2,
		B: 0xf5,
		A: 0xff,
	})
	for x, c := range g.players[0].hand {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(x)*5, 75)
		screen.DrawImage(c.img, op)
	}
}
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}
