package parse_test

import (
	"strings"
	"testing"

	"github.com/sunfmin/snippetgo/parse"
	"github.com/theplant/testingutils"
)

func TestAll(t *testing.T) {

	var cases = []struct {
		name     string
		file     string
		expected []*parse.Snippet
		err      string
	}{
		{
			name: "hello",
			file: "./fixtures/simple.go",
			expected: []*parse.Snippet{
				{
					Name: "Helloworldprint",
					Code: `fmt.Println(c)`,
				},
				{
					Name: "Helloworld",
					Code: `fmt.Println("Hello")
for _, c := range "Hello" {
	fmt.Println(c)
}`,
				},
			},
		},
		{
			name: "hello js",
			file: "./fixtures/hello.js",
			expected: []*parse.Snippet{

				{
					Name: "HelloworldJS",
					Code: `components: {
    EditorContent,
    EditorMenuBubble,
    Icon,
},

props: ['value'],`,
				},
			},
		},
		{
			name: "error1",
			file: "./fixtures/error1.go",
			err:  "not matched",
		},
		{
			name: "error2",
			file: "./fixtures/error2.go",
			err:  "not matched",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			snips, err := parse.Snippets(c.file)
			if err != nil {
				if len(c.err) > 0 {
					if !strings.Contains(err.Error(), c.err) {
						t.Error(err)
					}
				} else {
					t.Error(err)
				}
				return
			}
			diff := testingutils.PrettyJsonDiff(c.expected, snips)
			if len(diff) > 0 {
				t.Error(diff)
			}
		})
	}
}
