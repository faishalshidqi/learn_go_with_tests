package reflection

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{"struct with one string field",
			struct {
				Name string
			}{"Chris"},
			[]string{"Chris"},
		},
		{
			"struct with two string fields",
			struct {
				Name string
				City string
			}{"Chris", "London"},
			[]string{"Chris", "London"},
		},
		{
			"struct with non string field",
			struct {
				Name string
				Age  int
			}{"Chris", 33},
			[]string{"Chris"},
		},
		{
			"nested fields",
			struct {
				Name    string
				Profile struct {
					Age  int
					City string
				}
			}{"Chris", Profile{33, "London"}},
			[]string{"Chris", "London"},
		},
		{
			"pointers to things",
			&Person{
				"Chris",
				Profile{33, "London"},
			},
			[]string{"Chris", "London"},
		},
		{
			"slices",
			[]Profile{
				{33, "London"},
				{34, "Jakarta"},
			},
			[]string{"London", "Jakarta"},
		},
		{
			"arrays",
			[2]Profile{
				{33, "London"},
				{35, "Kuala Lumpur"},
			},
			[]string{"London", "Kuala Lumpur"},
		},
		{
			"maps",
			map[string]string{
				"Cow":   "Moo",
				"Sheep": "Baa",
			},
			[]string{"Moo", "Baa"},
		},
		{
			"channels",
			//make(chan struct{}),
			nil,
			[]string{"London", "Kuala Lumpur"},
		},
		{
			"functions",
			func() (Profile, Profile) {
				return Profile{33, "London"}, Profile{34, "Jakarta"}
			},
			[]string{"London", "Jakarta"},
		},
	}

	for _, test := range cases {
		switch test.Name {
		case "maps":
			t.Run(test.Name, func(t *testing.T) {
				got := make([]string, 0)
				walk(test.Input, func(input string) {
					got = append(got, input)
				})
				assertContains(t, got, "Moo")
				assertContains(t, got, "Baa")
			})
		case "channels":
			t.Run(test.Name, func(t *testing.T) {
				aChannel := make(chan Profile)
				go func() {
					aChannel <- Profile{33, "London"}
					aChannel <- Profile{35, "Kuala Lumpur"}
					close(aChannel)
				}()
				got := make([]string, 0)
				walk(aChannel, func(input string) {
					got = append(got, input)
				})
				assertEqual(t, test.ExpectedCalls, got)
			})
		default:
			t.Run(test.Name, func(t *testing.T) {
				got := make([]string, 0)
				walk(test.Input, func(input string) {
					got = append(got, input)
				})
				assertEqual(t, test.ExpectedCalls, got)
			})

		}
	}
}

func assertEqual(t *testing.T, expected interface{}, got interface{}) {
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("want %v, got %v", expected, got)
	}
}

func assertContains(t *testing.T, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, hay := range haystack {
		if hay == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("%s not found", needle)
	}
}
