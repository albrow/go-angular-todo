package models

import (
	"errors"
	"fmt"
	"github.com/stephenalexbrowne/zoom"
)

type Item struct {
	Id      string `redis:"-" json:"id"`
	Content string `redis:"content" json:"content"`
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

func FindAllItems() ([]*Item, error) {
	results, err := zoom.FindAll("item")
	if err != nil {
		return nil, err
	}
	items := make([]*Item, len(results))
	for index, result := range results {
		i, ok := result.(*Item)
		if !ok {
			msg := fmt.Sprintf("Couldn't convert %+v to *Item", result)
			return nil, errors.New(msg)
		}
		items[index] = i
	}
	return items, nil
}
