package main

func main() {

  var data = []int{}

	f := 0
	fuelOfFuel := 0

	for _, d := range data  {
		f += fuel(d)
		fuelOfFuel += fuelOfFuel(d)
	}
}

func fuelOfFuel(input int) int {
	if input == 0 {
		return input
	}

	n := input / 3
	if n < 2 {
		n = 0
	} else {
		n = n-2
	}

	input = n
	
	return input + fuelOfFuel(input)
}

func fuel(input int) int {
	return (input / 3) - 2
}
