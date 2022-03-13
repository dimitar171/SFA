package main

import (
	"fmt"
	"log"
)

type Action func() error

func SafeExec(action Action) Action {

	defer func() {
		if r := recover(); r != nil {
			log.Fatalf("there was an error: %v !", r)
		} else {
			fmt.Println("couldn't recover")
		}
	}()

	if action() != nil {
		return func() error {
			return fmt.Errorf("safe exec:  %w", action())
		}
	}

	return nil
}

func main() {

	/* FuncError := func() error {
		return errors.New("function with error")  //try function with error
	}
	err := SafeExec(FuncError)  */

	// FuncWithotError := func() error { //try function with no error
	// 	return nil
	// }
	// err := SafeExec(FuncWithotError)

	FuncWithPanic := func() error { // try function with panic msg
		panic("Panic error msg")
	}
	err := SafeExec(FuncWithPanic)

	// FuncWithPanic := func() error { // try function with panic nil msg
	// 	panic(nil)
	// }
	// err := SafeExec(FuncWithPanic)

	if err != nil {
		log.Fatalf("there was an error: %v !", err())
	} else {
		fmt.Println("there is no error!")
	}
}
