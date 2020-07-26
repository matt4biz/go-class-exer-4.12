[![Run on Repl.it](https://repl.it/badge/github/matt4biz/go-class-exer-4.12)](https://repl.it/github/matt4biz/go-class-exer-4.12)

# Go class: Exercise 4.12 from GoPL
These programs make up the answer to exercise 4.12 in [The Go Programming Language](http://www.gopl.io).

`xkcd-load.go` reads the xkcd.com JSON database until it gets two 404 responses (there is no comic #404, that's an easter egg returning an HTTP 404 reponse :-). Note that it may take a while and the resulting file will be at least 2 MB in size, all on one line.

```shell
$ go run ./load xkcd.json
read 2318 comics
```

`xkcd-find.go` takes some words to find from the command line (after the DB file) and finds comics whose title _or_ transcript matches all the words:

```shell
$ go run ./find xkcd.json someone bed sleep
read 2318 comics
https://xkcd.com/571/ 4/20/2009  "Can't Sleep"
found 1 comics
```

(This is a slight variation on the actual exercise description.)
