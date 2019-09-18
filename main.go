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

var skipDirs = []string{
	"node_modules/",
	".git/",
	"dist/",
}

func main() {
	flag.Parse()

	gf := gogen.File("f.go").Package(*pkg)

	err := filepath.Walk(".", func(path string, f os.FileInfo, err error) error {

		for _, dir := range skipDirs {
			if strings.Index(path, dir) >= 0 {
				//fmt.Println("skipping dir", path)
				return filepath.SkipDir
			}
		}

		if f.IsDir() {
			//fmt.Println("is dir", path)
			return nil
		}

		// to support other source files like js, ts, json
		// if !strings.HasSuffix(f.Name(), ".go") {
		//	 return nil
		// }

		//fmt.Println("is file", path)
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
