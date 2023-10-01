package logic

import (
	"errors"
)

type DataStore interface {
	UserNameForId(id string) (string, bool)
}

type Logger interface {
	Log(message string)
}

type SimleLogic struct {
	l Logger
	ds DataStore
}

func (sl SimleLogic) SayHello(id string) (string, error) {
	sl.l.Log("In Logic SayHello for id = " + id)
	name, ok := sl.ds.UserNameForId(id)
	if !ok {
		return "", errors.New("unknown user")
	}
	return "Hello, " + name, nil
}

func (sl SimleLogic) SayGoodbye(id string) (string, error) {
	sl.l.Log("in SayGoodbye for " + id)
	name, ok := sl.ds.UserNameForId(id)
	if !ok {
		return "", errors.New("unknown user")
	}
	return "Goodbye, " + name, nil
}

func NewSimpleLogic(l Logger, ds DataStore) SimleLogic {
	return SimleLogic{
		l: l,
		ds: ds,
	}
}