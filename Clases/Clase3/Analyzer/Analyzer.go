package Analyzer

import (
	"flag"
	"fmt"
	"regexp"
	"strings"
)

func AnalyzeType() {
	// define flags
	size := flag.Int("size", 0, "Tamaño")
	fit := flag.String("fit", "f", "Ajuste")
	unit := flag.String("unit", "m", "Unidad")

	// Parse the command line into the defined flags. 
	flag.Parse()

	// Command line input "-size=3000 -unit=\"K a\""
	input := "-size=3000 -unit=K -fit=\"BF\"" 

	// Proccess the input string and set the values of the flags
	processInput(input, size, fit, unit)


	// validate fit equals to b/w/f
	if *fit != "b" && *fit != "w" && *fit != "f" {
		fmt.Println("Error: Fit must be b, w or f")
		return
	}

	// validate size > 0
	if *size <= 0 {
		fmt.Println("Error: Size must be greater than 0")
		return
	}

	// validate unit equals to k/m
	if *unit != "k" && *unit != "m" {
		fmt.Println("Error: Unit must be k or m")
		return
	}

	// Print the values of the flags
	fmt.Println("Size:", *size)
	fmt.Println("Fit:", *fit)
	fmt.Println("Unit:", *unit)

	// Luego puedes llamar a tu función mkdisk y pasarle los valores necesarios
	// mkdisk(*size, *path, *fit, *unit)

}

func processInput(input string, size *int, fit *string, unit *string) {
	re := regexp.MustCompile(`-(\w+)=("[^"]+"|\S+)`)

	matches := re.FindAllStringSubmatch(input, -1)

	for _, match := range matches {
		flagName := match[1]
		flagValue := match[2]

		// Delete quotes if they are present in the value
		flagValue = strings.Trim(flagValue, "\"")

		switch flagName {
		case "size":
			sizeValue := 0
			fmt.Sscanf(flagValue, "%d", &sizeValue)
			*size = sizeValue
		case "fit":
			flagValue = flagValue[:1]
			flagValue = strings.ToLower(flagValue)
			*fit = flagValue
		case "unit":
			flagValue = strings.ToLower(flagValue)
			*unit = flagValue
		default:
			fmt.Println("Error: Flag not found")
		}
	}
}