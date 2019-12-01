package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var totalFuelRequired = 0

	if modules, err := readLines("input.txt"); err == nil {
		for _, moduleMass := range modules {
			totalFuelRequired += calculateRequiredFuel(moduleMass)
		}
	} else {
		log.Fatalf("readLines: %s", err)
	}

	fmt.Println("Total fuel required:", totalFuelRequired)

}

func calculateRequiredFuel(mass int) int {
	requiredFuel := mass/3 - 2
	if requiredFuel > 0 {
		return requiredFuel + calculateRequiredFuel(requiredFuel)
	}
	return 0
}

func readLines(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if mass, err := strconv.Atoi(scanner.Text()); err == nil {
			lines = append(lines, mass)
		}
	}

	return lines, scanner.Err()
}
