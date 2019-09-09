package fixtures

import (
	"fmt"
)

func error2() {
	fmt.Println("Hello")
	// @snippet_begin(Helloworld)
	for _, c := range "Hello" {
		// @snippet_begin(Helloworldprint)
		fmt.Println(c)
		// @snippet_end
	}
}
