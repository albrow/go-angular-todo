package models

import (
	"github.com/stephenalexbrowne/zoom"
)

func Initialize() error {
	zoom.Init()

	zoom.Register(new(Item), "item")

	return nil
}
