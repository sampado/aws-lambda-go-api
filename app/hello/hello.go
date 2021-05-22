package hello

import "fmt"

func HelloMessage(name string) string {
	if name == "" {
		return "Hello World from Go"
	}

	return fmt.Sprintf("Hello %s from Go", name)
}
