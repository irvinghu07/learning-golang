package greetings

import (
	"fmt"
	"errors"
	"math/rand"
	"time"
)

func Hello(name string) (string, error){
	if "" == name{
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf(randomFormat(), name)
	// toggle this one in order to trigger an error
	// message := fmt.Sprintf(randomFormat())
	return message, nil
}

func Hellos(names []string) (map[string]string, error){
	messages := make(map[string]string)
	for _, name := range names{
		msg, err := Hello(name)
		if nil != err{
			return nil, err
		}
		messages[name] = msg
	}
	return messages, nil
}


func init(){
	rand.Seed(time.Now().UnixNano())
}

func randomFormat()string{
	formats := []string{
		"Greeting from %v",
		"Hello world from %v",
		"Hi there, %v",
	}
	return formats[rand.Intn(len(formats))]
}