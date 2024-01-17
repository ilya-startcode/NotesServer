package main

import (
	"flag"
	"main/NotesServer/controller/httpserver"
	"main/NotesServer/gates/storage"
	"main/NotesServer/gates/storage/list"
	"main/NotesServer/gates/storage/mp"
)

func main() {
	var st storage.Storage
	var useList bool
	flag.BoolVar(&useList, "list", false, "A boolean flag")
	flag.Parse()

	if useList {
		st = list.NewList()
	} else {
		st = mp.NewMap()
	}
	hs := httpserver.NewHttpServer(":8080", st)
	hs.Start()
}
