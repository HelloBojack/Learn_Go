package main

import "fmt"

type Human struct {
	sex string
}

func (t *Human) Eat() {
	fmt.Println("Eating ...")
}

type Hero struct {
	Human
	Name string
	Lv   int
}

func (t *Hero) getName() string {
	return t.Name
}
func (t *Hero) setName(newName string) {
	t.Name = newName
}

func main() {
	h1 := Hero{
		Human{"male"},
		"bojack",
		999,
	}
	h1Name := h1.getName()

	fmt.Printf("%v\n", h1Name)

	h1.setName("xlk")

	fmt.Printf("%v\n", h1)

	h1.Eat()

}
