package models

import (
	"github.com/stephenalexbrowne/zoom"
)

type Item struct {
	Id      string
	Content string
}

func (p *Item) GetId() string {
	return p.Id
}

func (p *Item) SetId(id string) {
	p.Id = id
}

func NewItem(content string) *Item {
	p := &Item{Content: content}
	return p
}

func FindItemById(id string) (*Item, error) {
	result, err := zoom.FindById("item", id)
	if err != nil {
		return nil, err
	}
	p := result.(*Item)
	return p, nil
}
