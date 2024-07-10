package mod

import (
	"fmt"

	"github.com/HaraldWik/go-game-2/scr/abus"
)

type Physics struct {
}

func (p *Physics) Update() {
	fmt.Print(abus.SceneManager.Scenes[0].CurObj, " \n")
}
