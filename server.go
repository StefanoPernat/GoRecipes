package main

import (
	//"fmt"
	"github.com/go-martini/martini"
	//"io/ioutil"
)

func main() {
	m := martini.Classic()

	m.Get("/top", func() string {
		return "saranno i top 10"
	})

	m.Run()
}
