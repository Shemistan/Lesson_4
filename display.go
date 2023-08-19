package main

import (
	"fmt"
)

// Возвращает тип отчета. Годовой, полугодовой или квартальный.
func display() int {
	repeat := true
	var reportType string
	fmt.Println("\nВыберите вид отчета:\n1. Годовой\n2. Полугодовой.\n3. Квартальный")
	for repeat {
		_, err := fmt.Scan(&reportType)
		if err != nil {
			return 0
		}
		if reportType == "1" || reportType == "2" || reportType == "3" {
			repeat = false
		} else {
			fmt.Println("Введите цифру 1, 2, или 3")
		}
	}
	if reportType == "1" {
		return year
	}
	if reportType == "2" {
		return semiyear
	}
	return quarter
}

// Возвращает "Да" или "НЕТ" (нужен ли пользователью новый отчет).
func repeatReport() int {
	reset()
	prepareRecord := true
	var returnval int
	var yesNo string
	fmt.Println("\nПриготовить новый отчет? Y/N")
	for prepareRecord {
		_, err := fmt.Scan(&yesNo)
		if err != nil {
			return 0
		}
		if yesNo == "y" || yesNo == "Y" {
			prepareRecord = false
			returnval = 1
		} else if yesNo == "n" || yesNo == "N" {
			prepareRecord = false
			returnval = 2
		} else {
			fmt.Println("Введите N или Y")
		}
	}
	return returnval
}
