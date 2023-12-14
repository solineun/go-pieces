package main

import "sync"

func main() {
	
}

type SlowParser interface {
	Parse(string) string
}

var parser SlowParser
var once sync.Once

func Parse(data string) string {
	once.Do(func ()  {
		parser = initParser()
	})
	return parser.Parse(data)
}

func initParser() SlowParser {
	// configuration and loading
}