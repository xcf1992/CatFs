package main

import (
	"code.google.com/p/go-uuid/uuid"
	"github.com/proj-223/CatFs/client"
	"time"
)

func bandWidthBench(args []string) {
	t1 := time.Now()
	fi, err := client.Create(uuid.New())
	if err != nil {
		printError(err)
		return
	}
	for i := 0; i < 100*(1<<10); i++ {
		buf := make([]byte, 1<<10)
		_, err := fi.Write(buf)
		if err != nil {
			printError(err)
			return
		}
	}
	fi.Close()
	t2 := time.Now()
	td := t2.UnixNano() - t1.UnixNano()
	println(td)
}

func printError(err error) {
	println("Error :", err.Error())
}