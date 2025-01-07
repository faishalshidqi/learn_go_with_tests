package poker

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

const (
	PlayerPrompt         = "Please enter the number of players: "
	BadPlayerAmountInput = "Bad value received for number of players, please try again with a number"
	BadWinnerInputMsg    = "invalid winner input, expect format of 'PlayerName wins'"
)

type CLI struct {
	in   *bufio.Scanner
	out  io.Writer
	game Game
}

func NewCLI(in io.Reader, out io.Writer, game Game) *CLI {
	return &CLI{
		in:   bufio.NewScanner(in),
		out:  out,
		game: game,
	}
}

func (cli *CLI) readLine() string {
	cli.in.Scan()
	return cli.in.Text()
}

func (cli *CLI) PlayPoker() {
	fmt.Fprint(cli.out, PlayerPrompt)
	numberOfPlayersIn := cli.readLine()
	numberOfPlayers, err := strconv.Atoi(strings.Trim(numberOfPlayersIn, "\n"))
	if err != nil {
		fmt.Fprintf(cli.out, BadPlayerAmountInput)
		return
	}
	cli.game.Start(numberOfPlayers, cli.out)
	winner := cli.readLine()
	winner, err = extractWinnerName(winner)
	if err != nil {
		fmt.Fprintf(cli.out, BadWinnerInputMsg)
		return
	}
	cli.game.Finish(winner)
}

func extractWinnerName(input string) (string, error) {
	if !strings.Contains(input, " wins") {
		return "", errors.New(BadWinnerInputMsg)
	}
	return strings.Replace(input, " wins", "", 1), nil
}
