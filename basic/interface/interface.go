package main

import "fmt"

// Interface Men is implemented by Human, Student and Employee
// because it contains methods implemented by them.
type Men interface {
	SayHi()
	Sing(lyrics string)
}

type Human struct {
	name  string
	age   int
	phone string
}

//通过嵌入类型，与内部类型相关联的所有字段、方法、标志符等等所有，都会被外包类型所拥有，就像外部类型自己的一样
//这就达到了代码快捷复用组合的目的，而且定义非常简单，只需声明这个类型的名字就可以了。
type Student struct {
	Human  //an anonymous field of type Human
	school string
	loan   float32
}

type Employee struct {
	Human   //an anonymous field of type Human
	company string
	money   float32
}

//human实现了men接口中的SayHi()方法
func (h Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

//human实现了men接口中的Sing()方法
func (h Human) Sing(lyrics string) {
	fmt.Println("La la la la...", lyrics)
}

// Employees覆盖human方法中的name, age,phone
// 嵌入类型的强大，体现如果内部类型实现了某个接口，那么外部类型也被认为实现了这个接口。本例中可以认为Employee实现了men接口
// 对于内部类型的属性和方法访问上，我们可以用外部类型直接访问，也可以通过内部类型进行访问，例如这里的e.name
// 但是我们为外部类型新增的方法属性字段，只能使用外部类型访问，因为内部类型没有这些。
func (e Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
		e.company, e.phone) //Yes you can split into 2 lines here.
}

func main() {
	mike := Student{Human{"Mike", 25, "222-222-XXX"}, "MIT", 0.00}
	paul := Student{Human{"Paul", 26, "111-222-XXX"}, "Harvard", 100}
	sam := Employee{Human{"Sam", 36, "444-222-XXX"}, "Golang Inc.", 1000}
	Tom := Employee{Human{"Sam", 36, "444-222-XXX"}, "Things Ltd.", 5000}

	//声明一个接口类型
	var i Men

	//接口存储Student类型
	i = mike
	fmt.Println("This is Mike, a Student:")
	i.SayHi()
	i.Sing("November rain")

	//接口存储Employee类型
	i = Tom
	fmt.Println("This is Tom, an Employee:")
	i.SayHi()
	i.Sing("Born to be wild")

	//a slice of Men
	fmt.Println("Let's use a slice of Men and see what happens")
	x := make([]Men, 3)
	//These elements are of different types that satisfy the Men interface
	x[0], x[1], x[2] = paul, sam, mike

	for _, value := range x {
		value.SayHi()
	}
}
