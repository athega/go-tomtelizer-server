package main

import (
	"log"
	"strconv"
	"time"
)

func timestampString() string {
	return strconv.FormatInt(time.Now().UnixNano(), 10)
}

func puts(a ...interface{}) {
	log.Println(a...)
}

func fatal(a ...interface{}) {
	log.Fatal(a...)
}

func atoi(s string) (n int) {
	n, _ = strconv.Atoi(s)

	return
}
