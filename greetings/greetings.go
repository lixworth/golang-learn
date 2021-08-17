package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Hello(name string) (string,error) {
	if name == "" {
		return "",errors.New("empty name")
	}
	//message := fmt.Sprintf("Hi, %v. Welcome! \n",name)
	message := fmt.Sprintf(randomFormat(),name)
	return message,nil
}


func Hellos(names []string) (map[string]string,error) {
	messages := make(map[string]string)
	for _,name := range names {
		message,err := Hello(name)
		if err != nil {
			return nil, err
		}

		messages[name] = message
	}

	return messages,nil
}
func init() {
	//fmt.Printf("init run")

	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	formats := []string {
		"Hi, %v. Welcome! \n",
		"Great to see you, %v! \n",
		"Hail, %v! Well met! \n",
	}

	return formats[rand.Intn(len(formats))]
}
