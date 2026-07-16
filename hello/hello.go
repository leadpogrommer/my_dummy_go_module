package hello

import "fmt"

func GetHello(name string) string {
	return fmt.Sprintf("Hello, %s (v1.0.0)", name)
}