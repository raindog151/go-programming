package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {

	ridgeline := truck {
		vehicle: vehicle {
			doors: 2,
			color: "silver",
		},
		fourWheel: true,
	}

	fmt.Println(ridgeline)

	accord := sedan {
		vehicle: vehicle {
			doors: 4,
			color: "midnight blue",
		},
		luxury: false,
	}

	fmt.Printf("Honda Ridgeline: \nColor: %s\nDoors: %d\nFour Wheel: %v\n\n", ridgeline.color, ridgeline.doors, ridgeline.fourWheel)
	fmt.Printf("Honda Accord: \nColor: %s\nDoors: %d\nLuxury: %v\n", accord.color, accord.doors, accord.luxury)

}
