package interactions

import "fmt"

func Curse(name string) string {
	if name == "" {
		name = "default"
	}
	return fmt.Sprintf("Go to hell, %s!", name)
}
