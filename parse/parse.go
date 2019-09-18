package parse

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Snippet struct {
	Name string
	Code string
}

type stackElement struct {
	startIndex int
	snippet    *Snippet
}

const markBegin = "@snippet_begin"
const markEnd = "@snippet_end"

func Snippets(file string) (r []*Snippet, err error) {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	content := string(b)
	if strings.Index(content, markBegin) < 0 {
		return
	}
	lines := strings.Split(content, "\n")

	var stack []*stackElement

	for i, c := range lines {

		if name, ok := snippetName(c); ok {
			stack = append(stack, &stackElement{
				startIndex: i,
				snippet:    &Snippet{Name: name},
			})
		}

		last := peek(stack)
		if snippetEnd(c) {
			if last != nil {
				last.snippet.Code = strings.Join(
					removeIndent(
						cleanInner(
							lines[last.startIndex:i],
						),
					),
					"\n",
				)
				r = append(r, last.snippet)
				stack = stack[0 : len(stack)-1]
			} else {
				err = fmt.Errorf("@snippet_begin and @snipped_end not matched at %d", i)
				return
			}
		}
	}

	if len(stack) > 0 {
		err = fmt.Errorf("@snippet_begin and @snipped_end not matched at %d", stack[0].startIndex)
		return
	}

	return
}

func peek(stack []*stackElement) *stackElement {
	if len(stack) == 0 {
		return nil
	}
	return stack[len(stack)-1]
}

func removeIndent(code []string) (r []string) {
	l1 := strings.TrimSpace(code[0])
	trimIndex := strings.Index(code[0], l1)
	for _, c := range code {
		if len(c) < trimIndex {
			r = append(r, c)
			continue
		}
		if len(strings.TrimSpace(c[0:trimIndex])) == 0 {
			r = append(r, c[trimIndex:])
			continue
		}

		panic(fmt.Sprintf("%s can't be trim with %d", c, trimIndex))
	}

	return
}

func cleanInner(lines []string) (r []string) {
	for _, l := range lines {
		if _, ok := snippetName(l); ok {
			continue
		}
		if snippetEnd(l) {
			continue
		}
		r = append(r, l)
	}
	return
}

func snippetEnd(line string) bool {
	if strings.Index(line, markEnd) < 0 {
		return false
	}
	return true
}

func snippetName(line string) (name string, isSnippet bool) {
	if strings.Index(line, markBegin) < 0 {
		isSnippet = false
		return
	}

	start := strings.Index(line, "(")
	end := strings.LastIndex(line, ")")
	if end <= start {
		isSnippet = false
		return
	}

	return line[start+1 : end], true
}
