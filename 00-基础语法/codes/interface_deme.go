package main

import (
	"fmt"
	"time"
)

func main() {
	var a = time.Time{}
	fmt.Println(a)
	locationStart, _ := time.ParseInLocation("2006-01-02 15:04:05", "2021-02-14 00:00:00", time.UTC)
	fmt.Println(locationStart)
}

// 定义struct
type Human struct {
	name  string
	age   int
	phone string
}
type Student struct {
	Human  // 匿名字段
	school string
	loan   float32
}
type Employee struct {
	Human   // 匿名字段
	company string
	money   float32
}

// Human对象实现SayHi()方法
func (h Human) SayHi() {
	fmt.Printf("Hi, I am %s, you can call me on %s\n", h.name, h.phone)
}

// Human对象实现Sing()方法
func (h Human) Sing(lyrics string) {
	fmt.Println("La la la...", lyrics)
}

// Human对象实现Guzzle()方法
func (h Human) Guzzle(beerStein string) {
	fmt.Println("Guzzle Guzzle Guzzle...", beerStein)
}

// Employee对象重写SayHi()方法
func (e Employee) SayHi() {
	fmt.Printf("Hi I am %s, I work at %s. Call me on %s\n", e.name, e.company, e.phone)
}

// Student对象实现BorrowMoney()方法
func (s Student) BorrowMoney(amount float32) {
	s.loan += amount
}

// Employee对象实现SpendSalary()方法
func (e Employee) SpendSalary(amount float32) {
	e.money -= amount
}

// 定义interface，interface是一组method签名的组合
// interface可以被任意对象实现，一个对象也可以实现多个interface
// 任意类型都实现了空interface（也就是包含0个method的interface）
// 空interface可以存储任意类型的值
// interface Men的3个method被Human,Student,Employee实现，也就是这3个对象都实现了interface Men。即：
// interface定义了一组方法，如果某个对象实现了某个接口的所有方法，则此对象就实现了此接口。
type Men interface {
	SayHi()
	Sing(lyrice string)
	Guzzle(beerStein string)
}

// interface YoungChap的BorrowMoney() method只被Student对象实现，也就是只有Student实现了YoungChap
type YoungChap interface {
	SayHi()
	Sing(song string)
	BorrowMoney(amount float32)
}

// interface ElderlyGent的SpendSalary() method只被Employee对象实现，也就是只有Employee实现了ElderlyGent
type ElderlyGent interface {
	SayHi()
	Sing(song string)
	SpendSalary(amount float32)
}
