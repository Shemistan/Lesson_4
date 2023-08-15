package main

import "github.com/Shemistan/Lesson_4/entities"

func main() {
	var family entities.Family
	family.Init()

	family.Logic()
}