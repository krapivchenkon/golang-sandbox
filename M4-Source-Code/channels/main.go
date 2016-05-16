package main

import (
	"fmt"
	"strconv"
	"time"
)

var frameId = 0
var frameName = ""
var assemblyArrangement [3]string

func main() {
	assemblyArrangement[0] = "frame"
	assemblyArrangement[1] = "body"
	assemblyArrangement[2] = "interior"

	numberOfCars := 3

	framesToCreate := len(assemblyArrangement)
	frameInfoChan := make(chan string)

	for i := 1; i < numberOfCars+1; i++ {
		for stageNumber := 0; stageNumber < framesToCreate; stageNumber++ {
			go assemblyStage(frameInfoChan, assemblyArrangement[stageNumber], stageNumber, framesToCreate, i)
			time.Sleep(time.Millisecond * 1000)
		}
	}

	/*for i := 0; i < framesToCreate; i++ {
		go assembleFrame(frameInfoChan)
		go addBody(frameInfoChan)
		go addInterior(frameInfoChan)
		time.Sleep(time.Millisecond * 1000)
	}*/
}

func assemblyStage(frameInfoChan chan string, stage string, stageNumber int, framesToCreate int, carId int) {
	nextStage := "paint"
	if stageNumber < framesToCreate {
		frameName = "Frame ID" + strconv.Itoa(carId)
		if stageNumber != framesToCreate-1 {
			nextStage = assemblyArrangement[stageNumber+1]
		}
	}
	fmt.Println("Add", stage, "to", frameName, "and proceed to", nextStage)
	frameInfoChan <- frameName
	time.Sleep(time.Millisecond * 10)
}

func assembleFrame(frameInfoChan chan string) {
	frameId++
	frameName = "Frame ID" + strconv.Itoa(frameId)
	fmt.Println("Frame assembly complete.", frameName, "Proceed to body")
	frameInfoChan <- frameName
	time.Sleep(time.Millisecond * 5)
}

func addBody(frameInfoChan chan string) {
	body := <-frameInfoChan
	fmt.Println("Add body to", body, "and proceed to interior.")
	frameInfoChan <- frameName
	time.Sleep(time.Millisecond * 5)
}

func addInterior(frameInfoChan chan string) {
	interior := <-frameInfoChan
	fmt.Println("Add interior to", interior, "and proceed to paint.")
	time.Sleep(time.Millisecond * 5)

}
