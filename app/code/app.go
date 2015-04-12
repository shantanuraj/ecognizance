package main

import (
	"github.com/codegangsta/negroni"
)

func main() {
	n := negroni.Classic()
	n.UseHandler(GetHandlers())
	n.Run(":8000")
}
