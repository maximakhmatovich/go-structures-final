package actioninfo

import (
	"errors"
	"fmt"
	"log"
)

var errInvalidInput = errors.New("некорректные входные данные")

type DataParser interface {
	Parse(string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	if len(dataset) == 0 {
		log.Println(errInvalidInput)
	}

	for _, t := range dataset {
		err := dp.Parse(t)
		if err != nil {
			log.Println(err)
			continue
		}

		result, err := dp.ActionInfo()
		if err != nil {
			log.Println(err)
		}

		fmt.Println(result)
	}
}
