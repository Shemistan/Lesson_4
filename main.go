package main

import (
	"fmt"

	"github.com/kamolisrailov/Lesson_4.git/family"
)

func main() {

	var husband = family.Husband{
		human: family.Human{"Mike", 30, 100},
	}

	fmt.Println(husband.human.name)
}
