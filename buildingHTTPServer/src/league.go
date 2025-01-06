package src

import (
	"encoding/json"
	"fmt"
	"io"
)

type League []Player

func (l League) FindByName(name string) *Player {
	for i, player := range l {
		if player.Name == name {
			return &l[i]
		}
	}
	return nil
}

func NewLeague(readr io.Reader) (players []Player, err error) {
	err = json.NewDecoder(readr).Decode(&players)
	if err != nil {
		err = fmt.Errorf("problem parsing league: %s", err.Error())
	}
	return

}
