package inMemory

import "buildingHTTPServer/src"

type MemoryPlayerStore struct {
	store map[string]int
}

func NewInMemoryPlayerStore() *MemoryPlayerStore {
	return &MemoryPlayerStore{make(map[string]int)}
}

func (i *MemoryPlayerStore) GetPlayerScore(name string) int {
	return i.store[name]
}

func (i *MemoryPlayerStore) RecordWin(name string) {
	i.store[name]++
}

func (i *MemoryPlayerStore) GetLeague(order bool) src.League {
	league := make([]src.Player, 0)
	for name, wins := range i.store {
		league = append(league, src.Player{Name: name, Wins: wins})
	}
	return league
}
