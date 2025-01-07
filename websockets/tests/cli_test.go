package tests

import (
	"buildingHTTPServer/poker"
	"bytes"
	"fmt"
	"io"
	"strings"
	"testing"
	"time"
)

type GameSpy struct {
	StartedWith  int
	FinishedWith string
	StartCalled  bool
	FinishCalled bool
}

func (g *GameSpy) Start(numberOfPlayers int, to io.Writer) {
	g.StartedWith = numberOfPlayers
	g.StartCalled = true
}

func (g *GameSpy) Finish(winner string) {
	g.FinishedWith = winner
	g.FinishCalled = true
}

func userSends(messages ...string) io.Reader {
	return strings.NewReader(strings.Join(messages, "\n"))
}

type scheduledAlert struct {
	scheduledAt time.Duration
	amount      int
}

func (s scheduledAlert) String() string {
	return fmt.Sprintf("%d chips at %v", s.amount, s.scheduledAt)
}

type SpyBlindAlerter struct {
	alerts []scheduledAlert
}

func (s *SpyBlindAlerter) ScheduleAlertAt(duration time.Duration, amount int, to io.Writer) {
	s.alerts = append(s.alerts, scheduledAlert{scheduledAt: duration, amount: amount})
}

var (
	dummyStdout      = new(bytes.Buffer)
	dummySpyAlerter  = &SpyBlindAlerter{}
	dummyPlayerStore = &StubPlayerStore{}
)

func TestCLI(t *testing.T) {
	t.Run("record Chris win from user input", func(t *testing.T) {
		game := &GameSpy{}
		in := userSends("3", "Chris wins")
		cli := poker.NewCLI(in, dummyStdout, game)
		cli.PlayPoker()

		assertMessagesSentToUser(t, dummyStdout, poker.PlayerPrompt)
		assertGameStartedWith(t, game, 3)
		assertFinishCalledWith(t, game, "Chris")
	})
	t.Run("record Cleo win from user input", func(t *testing.T) {
		in := userSends("3", "Cleo wins")
		playerStore := &StubPlayerStore{}
		dummySpyAlerter := &SpyBlindAlerter{}
		game := poker.NewGame(dummySpyAlerter, playerStore)
		cli := poker.NewCLI(in, dummyStdout, game)
		cli.PlayPoker()

		assertPlayerWin(t, playerStore, "Cleo")
	})
	t.Run("it prints an error when a non numeric value is entered and does not start the game", func(t *testing.T) {
		game := &GameSpy{}

		stdout := &bytes.Buffer{}
		in := userSends("pies")

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		assertEqual(t, game.StartCalled, false)
		assertMessagesSentToUser(t, stdout, poker.PlayerPrompt, poker.BadPlayerAmountInput)
	})
	t.Run("it prints an error when the winner is declared incorrectly", func(t *testing.T) {
		game := &GameSpy{}
		stdout := &bytes.Buffer{}

		in := userSends("8", "Lloyd is a killer")
		cli := poker.NewCLI(in, stdout, game)

		cli.PlayPoker()

		assertEqual(t, game.FinishCalled, false)
		assertMessagesSentToUser(t, stdout, poker.PlayerPrompt, poker.BadWinnerInputMsg)
	})
}

func TestStartGame(t *testing.T) {
	t.Run("schedules alerts on game start for 5 players", func(t *testing.T) {
		blindAlerter := &SpyBlindAlerter{}
		game := poker.NewGame(blindAlerter, dummyPlayerStore)
		game.Start(5, io.Discard)
		cases := []scheduledAlert{
			{0 * time.Second, 100},
			{10 * time.Minute, 200},
			{20 * time.Minute, 300},
			{30 * time.Minute, 400},
			{40 * time.Minute, 500},
			{50 * time.Minute, 600},
			{60 * time.Minute, 800},
			{70 * time.Minute, 1000},
			{80 * time.Minute, 2000},
			{90 * time.Minute, 4000},
			{100 * time.Minute, 8000},
		}
		checkSchedulingCases(t, cases, blindAlerter)
	})
	t.Run("schedules alerts on game start for 7 players", func(t *testing.T) {
		blindAlerter := &SpyBlindAlerter{}
		game := poker.NewGame(blindAlerter, dummyPlayerStore)
		game.Start(7, io.Discard)
		cases := []scheduledAlert{
			{0 * time.Second, 100},
			{12 * time.Minute, 200},
			{24 * time.Minute, 300},
			{36 * time.Minute, 400},
		}
		checkSchedulingCases(t, cases, blindAlerter)
	})
	t.Run("prints an error when a non numeric value is entered and does not start the game", func(t *testing.T) {
		stdout := new(bytes.Buffer)
		in := strings.NewReader("Pies\n")
		game := &GameSpy{}
		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()
		assertEqual(t, game.StartCalled, false)
	})
}

func TestFinishGame(t *testing.T) {
	store := &StubPlayerStore{}
	game := poker.NewGame(dummySpyAlerter, store)
	winner := "Ruth"
	game.Finish(winner)
	assertPlayerWin(t, store, winner)
}

func assertFinishCalledWith(t *testing.T, game *GameSpy, player string) {
	t.Helper()
	if game.FinishedWith != player {
		t.Errorf("finished with %v; want %v", game.FinishedWith, player)
	}
}

func assertMessagesSentToUser(t testing.TB, stdout *bytes.Buffer, messages ...string) {
	t.Helper()
	want := strings.Join(messages, "")
	got := stdout.String()
	if got != want {
		t.Errorf("got %v; want %v", got, want)
	}
}

func assertGameStartedWith(t testing.TB, game *GameSpy, numPlayer int) {
	t.Helper()
	if game.StartedWith != numPlayer {
		t.Errorf("start called %v; want %v", game.StartCalled, numPlayer)
	}
}

func checkSchedulingCases(t *testing.T, cases []scheduledAlert, blindAlerter *SpyBlindAlerter) {
	for i, want := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			if len(blindAlerter.alerts) <= i {
				t.Fatalf("alert %d was not scheduled %v", i, blindAlerter.alerts)
			}
			got := blindAlerter.alerts[i]
			assertScheduledAlert(t, got, want)
		})
	}
}

func assertScheduledAlert(t *testing.T, alert scheduledAlert, want scheduledAlert) {
	amountGot := alert.amount
	if amountGot != want.amount {
		t.Errorf("Expected amount to be %d, got %d", want.amount, amountGot)
	}
	gotScheduleTime := alert.scheduledAt
	if gotScheduleTime != want.scheduledAt {
		t.Errorf("Expected scheduled time to be %d, got %d", want.scheduledAt, gotScheduleTime)
	}
}
