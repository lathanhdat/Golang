package main

import "fmt"

func main() {
	type vehicle struct {
		doors int
		color string
	}

	type truck struct {
		vehicle   vehicle
		fourWheel bool
	}

	type sedan struct {
		vehicle vehicle
		luxury  bool
	}

	truck1 := truck{
		vehicle: vehicle{
			doors: 4,
			color: "Blue",
		},
		fourWheel: true,
	}

	sedan1 := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "Red",
		},
		luxury: true,
	}

	fmt.Println(truck1, sedan1)
}
