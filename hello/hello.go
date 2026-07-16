package hello

import "fmt"

func GetHello(name string) string {
	return fmt.Sprintf("Hello, %s (v1.2.0-alpha.1)", name)
}
