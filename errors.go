package main

import ( "fmt"
	"errors"
)

func errorTest(wrong bool) (result string, err error) {
	if wrong {
		return "Fault", errors.New("Something got fucked")
	}
	return "Success", nil
}

func main() {
	if r, err := errorTest(false); err != nil {
		fmt.Println(r, err)
	}
	if r, err := errorTest(true); err != nil {
		fmt.Println(r, err)
	}
}