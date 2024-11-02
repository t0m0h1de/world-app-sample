package errhandler

import "log"

func ErrHandler(err error) {
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}
