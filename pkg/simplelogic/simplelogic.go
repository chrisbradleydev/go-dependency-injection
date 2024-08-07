package simplelogic

import (
	"errors"
	"fmt"

	s "github.com/chrisbradleydev/go-dependency-injection/pkg/datastore"
	l "github.com/chrisbradleydev/go-dependency-injection/pkg/logger"
)

type SimpleLogic struct {
	l  l.Logger
	ds s.DataStore
}

func (sl SimpleLogic) SayHello(userID string) (string, error) {
	sl.l.Log(fmt.Sprintf("SimpleLogic SayHello for userID: %s", userID))
	name, ok := sl.ds.UserNameForID(userID)
	if !ok {
		return incorrectUserID(userID)
	}
	return fmt.Sprintf("Hello, %s", name), nil
}

func (sl SimpleLogic) SayGoodbye(userID string) (string, error) {
	sl.l.Log(fmt.Sprintf("SimpleLogic SayGoodbye for userID: %s", userID))
	name, ok := sl.ds.UserNameForID(userID)
	if !ok {
		return incorrectUserID(userID)
	}
	return fmt.Sprintf("Goodbye, %s", name), nil
}

func NewSimpleLogic(l l.Logger, ds s.DataStore) SimpleLogic {
	return SimpleLogic{
		l:  l,
		ds: ds,
	}
}

func incorrectUserID(userID string) (string, error) {
	if userID != "" {
		return "", fmt.Errorf("unknown userID: %s", userID)
	}
	return "", errors.New("please provide user_id query string param")
}

type Logic interface {
	SayHello(userID string) (string, error)
}
