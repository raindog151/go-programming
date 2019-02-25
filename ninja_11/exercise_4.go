package main

import (
	"fmt"
	"log"
)

type sqrtError struct {
	lat  string
	long string
	err  error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
}

func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {

		myErr := sqrtError{
			lat: "50.2289 N",
			long: "99.4656 W",
			err: fmt.Errorf("Can not use non positive value %v to sqrt()!", f),
		}

		return 0, myErr
	}
	return 42, nil
}
