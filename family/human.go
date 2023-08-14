package family

import "fmt"

type Human struct {
	name      string
	fullnes   int16
	happiness int16
}

func (human *Human) eat() bool {
	fmt.Println("food has been eaten.")
	return true
}
