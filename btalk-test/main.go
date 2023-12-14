package main

import (
	"fmt"
	"time"

	"github.com/beanstalkd/go-beanstalk"
) 

var c, _ = beanstalk.Dial("tcp", ":11300")

func main() {
	go Produce()
	for {
		id, job, err := c.Reserve(0)
		if err == nil {
			Work(job)
			c.Delete(id)
		}		
	}
}

func Produce() {
	c.Put([]byte("hi"), 1, 0, 1 * time.Second)
}

func Work(job []byte) {
	time.Sleep(2 * time.Second)
	fmt.Println(string(job))
}