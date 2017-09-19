/**
The original interview question uses Java and is as
follows:

Write a class TempTracker with these methods:
1. insert()— records a new temperature
2. getMax()— returns the highest temp we've seen so far
3. getMin()— returns the lowest temp we've seen so far
4. getMean()— returns the mean of all temps we've seen so far
5. getMode()— returns a mode of all temps we've seen so far

This will need to be modified slightly since Go does not have the concept
of classes.
 */
package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
	The structure used to represent the
	thermometer.
 */
type Thermometer struct {
	max int
	min int
	mean float32
	mode int
	total int
	count[50] int
}

/**
	The main function.
 */
func main() {
	var thermometer Thermometer
	initThermometer(&thermometer)
	populateThermometer(&thermometer)
	printValues(&thermometer)
}

/**
	Initialise the thermometer.
 */
func initThermometer(thermometer *Thermometer) {
	thermometer.max = 0
	thermometer.min = 100
	thermometer.mean = 0
	for i := 0; i < len(thermometer.count); i++ {
		thermometer.count[i] = 0
	}
}

/**
	Send values into the thermometer and update
	the structs contents.
 */
func populateThermometer(thermometer *Thermometer) {

	for i := 0; i < len(thermometer.count); i++ {
		currentTemp := random(0, 50)
		thermometer.total += currentTemp

		fmt.Println("Current index: ", i, " Current Temperature: ", currentTemp)
		updateMax(thermometer, currentTemp)
		updateMin(thermometer, currentTemp)
		updateMode(thermometer, currentTemp)
	}
}

/*
	Update the thermometer for max temp.
 */
func updateMax(thermometer *Thermometer, currentTemp int) {
	if currentTemp > thermometer.max {
		thermometer.max = currentTemp
	}
}

/**
	Update the thermometer for min temp.
 */
func updateMin(thermometer *Thermometer, currentTemp int) {
	if currentTemp < thermometer.min {
		thermometer.min = currentTemp
	}
}

/**
	Update the mode.
 */
func updateMode(thermometer *Thermometer, currentTemp int) {
	thermometer.count[currentTemp]++
	if thermometer.count[currentTemp] > thermometer.mode {
		thermometer.mode = thermometer.count[currentTemp]
	}
}

/**
	Print the results stored in the thermometer
 */
func printValues(thermometer *Thermometer) {
	fmt.Println("Thermometer Readings: ")
	fmt.Println("Maximum Temperature: ", thermometer.max)
	fmt.Println("Minimum Temperature: ", thermometer.min)
	fmt.Println("Mean Temperature: ", thermometer.total / 100)
	fmt.Println("Mode: ", thermometer.mode)
}

/**
	Generate a random value between min-max
 */
func random(min, max int) int {
	rand.NewSource(time.Now().UnixNano())
	return rand.Intn(max - min) + min
}
