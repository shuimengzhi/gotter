package main

import (
	"gotter/logo"
	"gotter/router"
)

func main() {
	r := router.NewRouter()
	logo.Print()
	r.Run(":4000")
}
