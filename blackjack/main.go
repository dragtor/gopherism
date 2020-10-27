package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/dragtor/gopherism/deck"
	"os"
	"strconv"
	"strings"
)

var (
	numberOfPlayers *int
)

func validateFlags() {
	if *numberOfPlayers < 1 {
		panic("Invalid number of players")
	}
}

func init() {
	numberOfPlayers = flag.Int("n", 1, "Number of players")
	flag.Parse()
	validateFlags()
}

const (
	PLAYER_TYPE_HUMAN  = "HUMAN"
	PLAYER_TYPE_DEALER = "DEALER"
)

type Game struct {
	Deck        *deck.Cards
	Players     []*Player
	PlayerCount int
	Dealer      *Player
}

func New(deck *deck.Cards) Game {
	return Game{Deck: deck}
}

func gameInit(deck *deck.Cards, options ...func(*Game) *Game) *Game {
	game := New(deck)
	for _, option := range options {
		option(&game)
	}
	return &game
}

func Players(count int) func(*Game) *Game {
	return func(g *Game) *Game {
		g.PlayerCount = count
		for i := 0; i < count; i++ {
			g.Players = append(g.Players, NewPlayer())
		}
		g.Dealer = NewPlayer()
		return g
	}
}

func gamePoints(rank deck.Rank) int {
	switch rank {
	case deck.A:
		reader := bufio.NewReader(os.Stdin)
		for {
			fmt.Print("Select points for Rank A : \n1. [1]\n2. [11]\n Select [1/2]: ")
			text, _ := reader.ReadString('\n')
			opt, err := strconv.ParseInt(strings.TrimSpace(text), 10, 64)
			if err != nil {
				continue
			}
			if opt == 1 {
				return 1
			}
			return 11
		}
	case deck.Two:
		return 2
	case deck.Three:
		return 3
	case deck.Four:
		return 4
	case deck.Five:
		return 5
	case deck.Six:
		return 6
	case deck.Seven:
		return 7
	case deck.Eight:
		return 8
	case deck.Nine:
		return 9
	case deck.Ten:
		return 10
	case deck.Jack:
		return 10
	case deck.Queen:
		return 10
	case deck.King:
		return 10
	}
	return 0
}

func selectCardFromDeck(d *deck.Cards, index int) (*deck.Card, error) {
	// TODO : Handle error
	c := d.Cards[index]
	d.Cards[index] = d.Cards[len(d.Cards)-1]
	d.Cards = d.Cards[:len(d.Cards)-1]
	return &c, nil
}

func (g *Game) DistributeCards() {
	//distribute 2 cards to human player
	for _, p := range g.Players {
		for i := 0; i < 2; i++ {
			d := g.Deck
			c, _ := selectCardFromDeck(d, 0)
			p.Cards = append(p.Cards, *c)

		}
	}
	// distribute 1 Card to dealer
	c, _ := selectCardFromDeck(g.Deck, 0)
	g.Dealer.Cards = append(g.Dealer.Cards, *c)
}

func (g *Game) String() string {
	var allplayerstatus string
	for idx, p := range g.Players {
		playerstatus := fmt.Sprintf("player %d : %s\n", idx, p.Cards)
		allplayerstatus += playerstatus
	}
	dealerstatus := fmt.Sprintf("dealer : %+v\n", g.Dealer.Cards)
	allplayerstatus += dealerstatus

	return fmt.Sprintf(allplayerstatus)
}

type ConstPlayerOption string

const (
	CONST_PLAYER_OPTION_HIT   ConstPlayerOption = "CONST_PLAYER_OPTION_HIT"
	CONST_PLAYER_OPTION_STAND ConstPlayerOption = "CONST_PLAYER_OPTION_STAND"
)

type Player struct {
	Cards  []deck.Card
	Score  int
	Option ConstPlayerOption
}

func (p *Player) TotalScore() int {
	var pts int
	for _, c := range p.Cards {
		pts += gamePoints(c.Rank)
	}
	return pts
}

func NewPlayer() *Player {
	return &Player{Option: CONST_PLAYER_OPTION_HIT}
}

func userInput() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	return input, err

}

func Error(err error) {
	if err != nil {
		panic(err)
	}
}

func validateInput(expected []string, input string) bool {

	for _, e := range expected {
		if strings.ToLower(input) == e {
			return true
		}
	}
	return false
}

func (g *Game) Play() {
	for idx, p := range g.Players {
		for {
			fmt.Printf("Player %d: %+v\n", idx, g.Players[idx])
			fmt.Printf("Enter Hit/Stand [h/s]: ")
			input, err := userInput()
			if err != nil {
				continue
			}
			expectedInput := []string{"h", "s", "hit", "stand"}
			trimmedinput := strings.TrimSpace(input)
			res := validateInput(expectedInput, trimmedinput)
			if !res {
				fmt.Printf("Enter valid input\n")
				continue
			}
			if trimmedinput == "stand" || trimmedinput == "s" {
				break
			}
			d := g.Deck
			c, _ := selectCardFromDeck(d, 0)
			p.Cards = append(p.Cards, *c)
		}
		p.Score = p.TotalScore()
	}
	for {
		fmt.Printf("Dealer: %+v\n", g.Dealer)
		score := g.Dealer.TotalScore()
		if score > 17 {
			break
		}
		d := g.Deck
		c, _ := selectCardFromDeck(d, 0)
		g.Dealer.Cards = append(g.Dealer.Cards, *c)
	}

}
func (g *Game) Winner() {
	mindiff := 21
	var playeridx int
	for idx, p := range g.Players {
		score := p.TotalScore()
		if score > 21 {
			fmt.Printf("Player %d Lost", idx)
			continue
		}
		diff := 21 - score
		if diff < mindiff {
			playeridx = idx
			mindiff = diff
		}
	}
	ts := g.Dealer.TotalScore()
	if ts <= 21 {
		diff := 21 - ts
		if diff < mindiff {
			playeridx = -2
			mindiff = diff
		}
	}
	fmt.Printf("Winner is player %d, score is %d", playeridx, 21-mindiff)

}
func main() {
	d := deck.New(deck.Shuffle)
	g := gameInit(d, Players(2))
	g.DistributeCards()
	g.Play()
	g.Winner()
}
