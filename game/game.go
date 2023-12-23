package main

import "fmt"

/*
Struct
pointer receiver or value receiver
sturct methods
return stuct data types

No inheritance only embedding of structures
Implicit definition of varibales in Interfaces
*/
func main() {

	//how to define the structures
	//Option 1
	var i1 Item
	fmt.Printf("i1: %#v\n", i1)

	//Option 2
	i2 := Item{10, 20}
	fmt.Printf("i2: %#v\n", i2)

	//Option 3
	i3 := Item{
		Y: 230,
		// X: 134,
	}
	fmt.Printf("i3: %#v\n", i3)
	fmt.Println(NewItem(23, 43))
	fmt.Println(NewItem(23, -43))

	i3.Move(23, 55)
	fmt.Printf("i3 move %#v\n", i3)

	player1 := Player{
		Name: "gaya",
		Item: i2,
	}
	fmt.Printf("player %#v\n", player1)
	fmt.Printf("player %#v\n",player1.X) //you can also directly invoke the embed Item attributes
	player1.Move(43,66)
	fmt.Printf("player %#v",player1)

	// consider i want all of the Items, players to be moved to one thing
	// in Go interfaces is something wnat we want, not what we provide

	ms :=[]mover{
		&i2,
		&player1,
		&i3,
	}
	for _, m := range ms {
		fmt.Println(m)
	}
	MoveAll(ms,0,0)
	for _, m := range ms {
		fmt.Println(m)
	}
}

type mover interface{
	Move(x int, y int)
}

func MoveAll(ms []mover, x, y int){
	for _, m := range ms{
		m.Move(x,y)
	}

}
type Player struct {
	Name string
	Item //Embed Item
}

type Item struct {
	X int
	Y int
}

const (
	maxX = 100
	maxY = 60
)

// func NewItem(x int, y int) Item{
// func NewItem(x int, y int) *Item{
// func NewItem(x int, y int) (Item, error){
func NewItem(x int, y int) (*Item, error) {
	if x < 0 || x > maxX || y < 0 || y > maxY {
		return nil, fmt.Errorf("%d/%d out of range indices %d/%d", x, y, maxX, maxY)
	}
	i := Item{x, y}
	//Go does the "escape analysis" and stores it in heap
	//go run -gcflags=-m
	return &i, nil
}

// you can decide whetehr you wanna pass by reference or value
// consider the item will not change if you pas by value
// i is the "receiver"
func (i *Item) Move(x int, y int) {
	i.X = x
	i.Y = y
}
