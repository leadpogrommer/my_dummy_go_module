package hello

import (
	"fmt"

	"github.com/samber/lo"
)

func GetHello(name string) string {

	fmt.Println(lo.Map([]byte(name), func(b byte, _ int) byte { return b + 1 }))
	return fmt.Sprintf("Hello, %s (v1.3.0)", name)
}
