package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

// { "month":      "4",
//   "day":        "20"
//   "year":       "2009",
//   "num":        571,
//   . . .
//   "transcript": "[[Someone is in bed, . . . long int.",
//   "img":        "https://imgs.xkcd.com/comics/cant_sleep.png",
//   "title":      "Can't Sleep",
// }

type xkcd struct {
	Num        int    `json:"num"`
	Day        string `json:"day"`
	Month      string `json:"month"`
	Year       string `json:"year"`
	Title      string `json:"title"`
	Transcript string `json:"transcript"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "no file given")
		os.Exit(-1)
	}

	fn := os.Args[1]

	if input, err = os.Open(fn); err != nil {
		fmt.Fprintf(os.Stderr, err)
		os.Exit(-1)
	}

	var (
		input io.ReadCloser
		items []xkcd
		terms []string
		cnt   int
		err   error
	)

	for _, t := range os.Args[2:] {
		terms = append(terms, strings.ToLower(t))
	}

	if len(terms) < 1 {
		fmt.Fprintln(os.Stderr, "no search terms")
		os.Exit(-1)
	}

	// some code here

	fmt.Fprintf(os.Stderr, "found %d comics\n", cnt)
}
