package main

import (
	"fmt"
)

const(
	eatFood = 30
	buyCoat = 350
	earnMoney = 150
	behappyHusband = 20
	behappyWife = 60
	happless = 10
	hungry = 10
	buyFood = 30
	moneyforbuyFood = 30
	dirty = 5
	leveldirty = 90
	cleanHouse = 100
	year = 1
	days = 365
)

type person struct{
	nameperson string
	dagreehungry int32
	dagreehappiness int32
}

type homeCommon struct{
	allmoney int32
	allfood int32
	alldirty int32
}

type statistic struct{
	earnedMoney int32
	counteatFood int32
	countbougthCoat int32
}

func (p *person) eatFood(h *homeCommon){
	p.dagreehungry += eatFood
	p.dagreehappiness -= happless
	h.allfood -= eatFood
}

func (p *person) gameComp(){
	p.dagreehappiness += behappyHusband
	p.dagreehungry -= hungry
}

func (p *person) goWork(h *homeCommon){
	h.allmoney += earnMoney
	p.dagreehungry -= hungry
	p.dagreehappiness -= happless
}

func (p *person) buyFood(h *homeCommon){
	h.allfood += buyFood
	h.allmoney -= moneyforbuyFood
	h.alldirty += dirty
	p.dagreehungry -= hungry
	p.dagreehappiness -= happless
}

func (p *person) buyCoat(h *homeCommon){
	h.allmoney -= buyCoat
	p.dagreehungry -= hungry
	p.dagreehappiness += behappyWife
	h.alldirty += dirty
	if h.alldirty > 90 {
		p.dagreehappiness -= happless
	}
	if p.dagreehappiness >= 100 {
		p.dagreehappiness = 100
	}	
}

func (p *person) cleanHouse(h *homeCommon){
	h.alldirty -= cleanHouse
	p.dagreehungry -= hungry
	p.dagreehappiness -= happless
	if h.alldirty < 0 {
		h.alldirty = 0
	}
}

func (p *person) doActionsWife(h *homeCommon,stat *statistic){
	switch {
	case p.dagreehungry <= 10 && p.dagreehappiness > 10:
		if h.allfood >= eatFood {
			p.eatFood(h)
			h.alldirty += dirty
			if p.dagreehungry >= 100{
				p.dagreehungry = 100
			}
			if h.alldirty > 90 {
				p.dagreehappiness -= 10
			}
			stat.counteatFood++
		}else if h.allfood  > 0 && h.allfood < eatFood{
			p.dagreehungry += h.allfood
			p.dagreehappiness -= happless
			h.alldirty += dirty
			h.allfood -= 0
			if p.dagreehungry >= 100{
				p.dagreehungry = 100
			}
			if h.alldirty > 90 {
				p.dagreehappiness -= happless
			}
			stat.counteatFood++
		}
	case p.dagreehungry > 10 && p.dagreehappiness > 10:
		if h.alldirty >=90 {
			p.cleanHouse(h)
		}	
		if h.allmoney > 30 && h.allfood < 210 {
			p.buyFood(h)
		}
	case p.dagreehungry >= 10 && p.dagreehappiness <= 10:
		if h.allmoney >= 350 {
			p.buyCoat(h)	
			stat.countbougthCoat++
		}
	}
}

func (p *person) doActionsHusband(home *homeCommon, stat *statistic){
	switch {	
	case p.dagreehungry <= 10 && p.dagreehappiness > 10 && home.allfood >= 30:   
		p.eatFood(home)
		stat.counteatFood++
	case p.dagreehungry > 10 && p.dagreehappiness > 10:
		p.goWork(home)
		stat.earnedMoney++
	case p.dagreehappiness <= 10 && p.dagreehungry > 0:
		p.gameComp()
	}
}

func doLogic (){
	personHusband := person{
		nameperson: "John",
		dagreehungry: 30,
		dagreehappiness: 100,
	}
	personWife := person{
		nameperson: "Maria",
		dagreehungry: 30,
		dagreehappiness: 100,
	}
	home := homeCommon{
		allmoney: 100,
		allfood:  50,
		alldirty: 0,
	}
	stat := statistic{
		earnedMoney: 0,
		counteatFood: 0,
		countbougthCoat: 0,
	}
	for i := 1; i < year*days; i++ {
		personHusband.doActionsHusband(&home,&stat)
		personWife.doActionsWife(&home,&stat)
		if personHusband.dagreehappiness < 10 || personHusband.dagreehungry < 0 {
			fmt.Println("Husband Died :( ")
			return
		}
		if personWife.dagreehappiness < 10 || personWife.dagreehungry < 0 {
			fmt.Println("Wife Died :( ")
			return
		}
	}
	fmt.Println("Itogi jizni: ","\nBilo zarabotano deneg:",stat.earnedMoney*150,"\nBilo syedeno edi:",stat.counteatFood*30,"\nBilo kupleno shub:", stat.countbougthCoat)
}

func main() {
		doLogic()
}


