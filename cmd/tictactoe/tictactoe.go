package tictactoe

import (
	"github.com/JesseyBland/project-0/gameio/ai"
	"github.com/JesseyBland/project-0/gameio/two"
)

func main() {
	switch {
	case tp == true:
		two.Twoplayer()
	case av == true:
		ai.AivsAi()
	default:
		ai.Aiplayer()
	}

}
