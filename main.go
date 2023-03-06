package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	c "github.com/terminatingcode/martians/controller"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Error: no file path input")
		os.Exit(1)
	}
	fileName := os.Args[1]
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// todo: need to watch out for buffer overload
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	gridString := scanner.Text()
	// todo: ensure this is valid string
	gridCoordinates := strings.Split(gridString, " ")
	x, _ := strconv.Atoi(gridCoordinates[0])
	y, _ := strconv.Atoi(gridCoordinates[1])

	controller, err := c.Create(x, y)
	if err != nil {
		fmt.Println(err)
	}

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}

		// todo: handle potential newlines between robot initialisation and directions
		robotString := strings.Split(line, " ")
		x, _ := strconv.Atoi(robotString[0])
		y, _ := strconv.Atoi(robotString[1])
		robot, err := controller.ConnectRobot(x, y, robotString[2])
		if err != nil {
			fmt.Println(err)
		}

		scanner.Scan()
		directionsLine := scanner.Text()
		directionsString := strings.Split(directionsLine, "")
		for _, direction := range directionsString {
			controller.DirectRobot(robot, direction)

		}
	}

	fmt.Println(controller.ToString())
}
