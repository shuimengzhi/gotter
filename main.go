package main

import (
	"gotter/core"
	"gotter/logo"
	"gotter/router"
)

func main() {
	r := router.NewRouter()
	core.LoadCore()
	logo.Print()
	r.Run(":4000")
}
