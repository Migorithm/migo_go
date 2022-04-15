package main

import "fmt"

type People interface {
	Field()
	Salary()
}

type Developer struct {
	Age    int
	Sex    string
	Name   string
	Area   string
	Income int
}

type HR struct {
	Age    int
	Name   string
	Area   string
	Income int
}

type Company struct {
	Department []*Department
}

type Department struct {
	Name   string
	Member []People
}

func (r *Developer) Field() {
	fmt.Printf("I have a forte at %s a month!\n", r.Area)
}

func (r *Developer) Salary() {
	fmt.Printf("My Salary is %d a month!\n", r.Income)
}

func (r *HR) Field() {
	fmt.Printf("I have been here forever man! :in %s\n", r.Area)
}

func (r *HR) Salary() {
	fmt.Printf("My Salary is %d a month!\n", r.Income)
}

func main() {
	var Dev1 = &Developer{Age: 20, Sex: "Male", Name: "Jjun", Area: "BackEnd", Income: 10000}
	var Hr1 = &HR{Age: 20, Name: "Jjun", Area: "WemakePrice HR team", Income: 5000}

	var Dep *Department = &Department{Name: "IT", Member: []People{Dev1, Hr1}}
	var WeMakePrice *Company = &Company{}
	WeMakePrice.Department = append(WeMakePrice.Department, Dep)
	for _, val := range WeMakePrice.Department[0].Member {
		val.Field()
		val.Salary()
	}
}

func Worker(person People) {
	person.Field()
	person.Salary()
}

// Person  - 나이, 성별, 직업, 이름
// Developer  - 나이, 성별, 직업, 이름, 분야, 연봉

// Company - Department

// Department - IT, HR
