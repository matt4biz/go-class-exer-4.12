package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

// we don't actually use the struct to unmarshall the JSON
// coming back from the server; we just dump it out as text

func getOne(i int) []byte {
	url := fmt.Sprintf("https://xkcd.com/%d/info.0.json", i)
	resp, err := http.Get(url)

	if err != nil {
		fmt.Fprintf(os.Stderr, "stopped reading: %s\n", err)
		os.Exit(-1)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		// easter egg: #404 returns HTTP 404 - not found

		fmt.Fprintf(os.Stderr, "skipping %d: got %d\n", i, resp.StatusCode)
		return nil
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "bad body: %s\n", err)
		os.Exit(-1)
	}

	return body
}

func main() {
	var (
		output io.WriteCloser = os.Stdout
		err    error
		cnt    int
		fails  int
		data   []byte
	)

	if len(os.Args) > 0 {
		output, err = os.Create(os.Args[1])

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(-1)
		}

		defer output.Close()
	}

	// the output will be in the form of a JSON array,
	// so add the brackets before and after

	fmt.Fprint(output, "[")
	defer fmt.Fprint(output, "]")

	// some code here

	fmt.Fprintf(os.Stderr, "read %d comics\n", cnt)
}
