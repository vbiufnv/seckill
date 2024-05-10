package model

var Done chan int

func InitChannel() {
	Done = make(chan int, 1)

	Done <- 1
}
