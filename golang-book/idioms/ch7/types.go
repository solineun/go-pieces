package ch7

import (
	"fmt"
	"time"
)

type Score int
type Converter func(string)Score
type TeamScores map[string]Score

type Counter struct {
	total int
	lastUpdated time.Time	
}

func (c *Counter) Increment() {
	c.total++
	c.lastUpdated = time.Now()
}

func (c Counter) String() string {
	return fmt.Sprintf("total: %d, lasat updated: %v", c.total, c.lastUpdated)
}

type Adder struct {
	Start int
}

func (a Adder) AddTo(val int) int {
	return a.Start + val
}

type MailCategory int 

const (
	Uncategorized MailCategory = iota
	Personal
	Spam
	Social
	Advertisements
)