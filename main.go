package main

import (
	"fmt"

)

type person struct {
	
	name   string
	hungry int
	happy  int
}

type homeCommon struct {
	money int
	food  int
	dirty int
}

func (p *person) doActionsFv1(h *homeCommon)(int,int){
	counteatF :=0
	countCoat :=0
	switch {
	case p.hungry <= 10 && p.happy > 10:
		if h.food >= 30 {
			p.hungry += 30
			p.happy -= 10
			h.dirty += 5
			h.food -= 30
			if p.hungry >= 100{
				p.hungry = 100
			}
			if h.dirty > 90 {
				p.happy -= 10
			}
			if p.happy >= 100 {
				p.happy = 100
			//	p.hungry -= 10
			}	
			counteatF++
			return counteatF, countCoat
		}else if h.food  > 0 {
			p.hungry += h.food
			p.happy -= 10
			h.dirty += 5
			h.food -= 0
			if p.hungry >= 100{
				p.hungry = 100
			}
			if h.dirty > 90 {
				p.happy -= 10
			}
			if p.happy >= 100 {
				p.happy = 100
			//	p.hungry -= 10
			}	
			counteatF++
			return counteatF, countCoat
		}
		fallthrough
	case p.hungry > 10 && p.happy > 10:
		if h.dirty >=90 {
			p.hungry -= 10
			p.happy -= 10
			h.dirty = 0
			return counteatF, countCoat
		}	
		if h.money > 30 && h.food < 210 {
			p.happy -= 10
			p.hungry -= 10
			h.food += 30
			h.dirty += 5
			h.money -= 30
			return counteatF, countCoat
		}
		if p.hungry > 0 && p.happy > 0 && h.food > 30 && h.food <= 230{
			p.hungry += 30
			p.happy -= 10
			h.dirty += 5
			h.food -= 30
			if p.hungry >= 100{
				p.hungry = 100
			}
			if h.dirty > 90 {
				p.happy -= 10
			}
			if p.happy >= 100 {
				p.happy = 100
			//	p.hungry -= 10
			}	
			counteatF++
			return counteatF, countCoat
		}
		fallthrough
	case p.hungry >= 10 && p.happy <= 10:
		if h.money >= 350 {
			p.hungry -= 10
			p.happy += 60
			h.money -= 350
			h.dirty += 5
			if h.dirty > 90 {
				p.happy -= 10
			}
			if p.happy >= 100 {
				p.happy = 100
			}	
			countCoat++
			return counteatF, countCoat
		}
	}
	return counteatF, countCoat
}


func (p *person) doActionsM(h *homeCommon) (int, int){
	counteat := 0
	countmoney := 0
	switch {	
	
	case p.hungry <= 10 && p.happy > 10 && h.food >= 30:   
		p.hungry += 30
		p.happy -= 10
		h.food -= 30
		counteat++
		return counteat, countmoney		
	case p.hungry > 10 && p.happy > 10:
		p.hungry -= 10
		p.happy -= 10
		h.money += 150
		countmoney++
		return counteat, countmoney
	case p.happy <= 10 && p.hungry > 0:
		p.hungry -= 10
		p.happy += 20
		return counteat, countmoney
					
	}
	return counteat, countmoney
}


func main() {

	 personM := person{name: "John", hungry: 50, happy: 100}
	 personF := person{name: "Maria", hungry: 50, happy: 100}
	 home := homeCommon{money: 150, food: 50, dirty: 0}

	counteatM :=0
	counteatF :=0
	countMoney :=0
	countCoat :=0

	for i := 0; i<355; i++{
		per3, per4 := personF.doActionsFv1(&home)
	   	per1, per2 := personM.doActionsM(&home)
		if personF.hungry < 0 || personF.happy  == 0  {
			fmt.Println("Chel F umer")
		}
		if personM.hungry < 0 || personM.happy == 0 {
			fmt.Println("Chel M umer")
		}
		counteatM += per1
		countMoney += per2
		counteatF += per3
		countCoat += per4
	}

	fmt.Println("Itogi jizni: ","\nBilo zarabotano deneg:",countMoney*150,"\nBilo syedeno edi:",(counteatF+counteatM)*30,"\nBilo kupleno shub:", countCoat)




}


