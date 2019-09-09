package fixtures

import (
	"fmt"
)

func error1() {
	fmt.Println("Hello")
	for _, c := range "Hello" {
		// @snippet_begin(Helloworldprint)
		fmt.Println(c)
		// @snippet_end
	}
	// @snippet_end
}
