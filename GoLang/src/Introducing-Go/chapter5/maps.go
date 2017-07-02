package main

import "fmt"

func main() {
	simplemap()
	elementmap()
	otherwaydefinemap()
}

func simplemap()  {
	x := make(map[string]int)
	x["key"] = 10
	fmt.Println(x["key"])

	y := make(map[int]int)
	y[1] = 9
	fmt.Println(y[1])
}

func elementmap()  {
	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"
	fmt.Println(elements["Li"])
	// such way nothin returned
	fmt.Println(elements["Un"])
	// If element does not exist, then false
	// otherwise Beryllium true
	name, ok := elements["Un"]
	fmt.Println(name, ok)

	if name, ok := elements["Be"]; ok {
		fmt.Println(name, ok)
	}
}

func otherwaydefinemap()  {
	elements := map[string]string{
		"H":  "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
		"Be": "Beryllium",
		"B":  "Boron",
		"C":  "Carbon",
		"N":  "Nitrogen",
		"O":  "Oxygen",
		"F":  "Fluorine",
		"Ne": "Neon",
	}

	fmt.Println(elements["O"])
}