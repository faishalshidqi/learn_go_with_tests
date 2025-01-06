package poker

import (
	"bufio"
	"io"
	"strings"
)

type CLI struct {
	playerStore PlayerStore
	in          *bufio.Scanner
}

func NewCLI(playerStore PlayerStore, in io.Reader) *CLI {
	return &CLI{
		playerStore: playerStore,
		in:          bufio.NewScanner(in),
	}
}

func (cli *CLI) readLine() string {
	cli.in.Scan()
	return cli.in.Text()
}

func (cli *CLI) PlayPoker() {
	input := cli.readLine()
	cli.playerStore.RecordWin(extractWinnerName(input))
}

func extractWinnerName(input string) string {
	return strings.Replace(input, " wins", "", 1)
}
