package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/sunfmin/gogen"
	"github.com/sunfmin/snippetgo/parse"
)

var pkg = flag.String("pkg", "generated", "generated package name")

func main() {
	flag.Parse()

	gf := gogen.File("f.go").Package(*pkg)

	err := filepath.Walk(".", func(path string, f os.FileInfo, err error) error {
		if strings.Index(path, "node_modules") >= 0 {
			return filepath.SkipDir
		}

		if strings.Index(path, ".git") >= 0 {
			return filepath.SkipDir
		}

		if !strings.HasSuffix(f.Name(), ".go") {
			return nil
		}

		snippets, err := parse.Snippets(path)
		if err != nil {
			fmt.Fprint(os.Stderr, err)
		}

		for _, s := range snippets {
			gf.BodySnippet("var $NAME = string($BYTE)", "$NAME", s.Name, "$BYTE", fmt.Sprintf("%#+v", []byte(s.Code)))
		}

		return nil
	})

	if err != nil {
		panic(err)
	}

	err = gf.Fprint(os.Stdout, context.TODO())
	if err != nil {
		panic(err)
	}
}
