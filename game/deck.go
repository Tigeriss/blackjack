package game

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
	"math/rand"
	"strconv"
	"time"
)

type card struct {
	img    *ebiten.Image
	imgSrc string
	value  int
}

type deck struct {
	cardBackImg string
	cards       []card
}

func NewDeck() deck {
	d := make([]card, 0, 52)
	for i := 0; i < 52; i++ {
		val := getValueForCard(i)
		img, _, err := ebitenutil.NewImageFromFile("./img/deck_green.png")
		if err != nil {
			log.Fatal(err)
		}
		c := card{
			img:    img,
			imgSrc: "./img/" + strconv.Itoa(i) + ".png",
			value:  val,
		}
		d = append(d, c)
	}
	return deck{cardBackImg: "./img/deck_green.png", cards: d}
}

func (d *deck) shuffle() deck {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d.cards), func(i, j int) {
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	})
	return deck{}
}

func (d *deck) dealCard() card {
	// get top card
	c := d.cards[0]
	// remove dealt card from deck
	d.cards[0] = card{}
	d.cards = d.cards[1:]

	return c
}

func getValueForCard(index int) int {
	switch index % 13 {
	case 0:
		return 11
	case 1:
		return 2
	case 2:
		return 3
	case 3:
		return 4
	case 4:
		return 5
	case 5:
		return 6
	case 6:
		return 7
	case 7:
		return 8
	case 8:
		return 9
	default:
		return 10
	}
}

func (d deck) TestShuffle() {
	d.shuffle()
	for _, c := range d.cards {
		fmt.Println(c.imgSrc)
	}
}
