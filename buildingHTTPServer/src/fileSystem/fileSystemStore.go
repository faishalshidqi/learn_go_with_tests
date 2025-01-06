package fileSystem

import (
	"buildingHTTPServer/src"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"
)

type FSPlayerStore struct {
	Database *json.Encoder
	league   src.League
}

func NewFSPlayerStore(file *os.File) (*FSPlayerStore, error) {
	err := playerDBFileInit(file)
	if err != nil {
		return nil, err
	}
	league, err := src.NewLeague(file)
	if err != nil {
		return nil, fmt.Errorf("problem loading player store from file %s, %v", file.Name(), err.Error())
	}
	return &FSPlayerStore{
		json.NewEncoder(&src.Tape{File: file}),
		league,
	}, nil
}

func playerDBFileInit(file *os.File) error {
	file.Seek(0, io.SeekStart)
	info, err := file.Stat()
	if err != nil {
		return fmt.Errorf("problem getting file info from file %s: %v", file.Name(), err)
	}
	if info.Size() == 0 {
		file.Write([]byte("[]"))
		file.Seek(0, io.SeekStart)
	}
	return nil
}

func (f *FSPlayerStore) GetLeague(order bool) (players src.League) {
	if order {
		sort.Slice(f.league, func(i, j int) bool {
			return f.league[i].Wins > f.league[j].Wins
		})
	}
	players = f.league
	return
}

func (f *FSPlayerStore) GetPlayerScore(name string) int {
	player := f.league.FindByName(name)
	if player != nil {
		return player.Wins
	}
	return 0
}

func (f *FSPlayerStore) RecordWin(name string) {
	player := f.league.FindByName(name)
	if player != nil {
		player.Wins++
	} else {
		f.league = append(f.league, src.Player{Name: name, Wins: 1})
	}
	f.Database.Encode(f.league)
}
