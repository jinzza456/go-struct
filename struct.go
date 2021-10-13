package main

import "fmt"

// type Person struct {
// 	name string
// 	age  int
// }

// func (p Person) PrintName() { // Person 이라는 객체에 PrintName이라는 기능을 추가한다.
// 	fmt.Print(p.name)
// }

// func main() {
// 	var p Person

// 	p.name = "Smith"
// 	p.age = 24

// 	fmt.Println(p)

// 	p.PrintName()

type Student struct {
	name  string
	class int

	grade Grade // grade에 Grade라는 타입을 설정해줌
}

type Grade struct {
	name  string
	grade string
}

func (s Student) ViewGrade() { // Student의 기능 ViewGrade로 성적 조회가 가능한 기능
	fmt.Println(s.grade)
}

func (s Student) Studentinfo() {
	fmt.Println(s.name, s.class, "반")
}

func main() {

	var s Student
	s.name = "철수"
	s.class = 1

	s.grade.name = "수학"
	s.grade.grade = "C"

	s.Studentinfo()
	s.ViewGrade()
}
