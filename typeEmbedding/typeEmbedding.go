package main

import "fmt"

type Player struct {
	Name string
	Age  int
}

type Cricketer struct {
	Player
	kind string
}

func (P Player) PrintName() {
	fmt.Println(P.Name)
}

func (P *Player) SetAge(age int) {
	P.Age = age
}

func main() {
	var X = Cricketer{Player: Player{"Sachin Tendulkar", 40}, kind: "batsmen"}

	X.PrintName()
	X.Name = "Virat Kohli"
	X.SetAge(31)
	X.PrintName()
	fmt.Printf("%#v\n", X)

	return

}
