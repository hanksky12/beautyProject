package mouseAction

import (
	"beautyProject/internal/pkg/enum"
)

type MouseAction struct {
	*enum.Base
}

var (
	Move   = &MouseAction{&enum.Base{Number: 1, Name: "move", ChineseName: "移動"}}
	Click  = &MouseAction{&enum.Base{Number: 2, Name: "click", ChineseName: "點擊"}}
	Scroll = &MouseAction{&enum.Base{Number: 3, Name: "scroll", ChineseName: "滾動"}}
)

var Map = map[string]*MouseAction{
	Move.Name:   Move,
	Click.Name:  Click,
	Scroll.Name: Scroll,
}
