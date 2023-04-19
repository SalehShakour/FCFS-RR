package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
	"sync"
)

func RR() {
	var quantum float64 = 2

	processChan := make(chan string)
	inputChan := make(chan []Process)
	var wg sync.WaitGroup
	wg.Add(2)

	var maxFinishingTime = map[string]float64{}
	var arrTime = map[string]float64{}
	var exTime = map[string]float64{}

	go func() {
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Enter the number of processes: ")
		var count int
		_, err := fmt.Scanln(&count)
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}
		var processes []Process
		for i := 0; i < count; i++ {
			line, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Error reading input:", err)
				return
			}
			line = strings.TrimSuffix(line, "\n")
			parts := strings.Split(line, " ")
			if len(parts) != 3 {
				fmt.Println("Invalid input")
				return
			}

			name := parts[0]
			execTimeStr := parts[1]
			arrivalTimeStr := parts[2]

			var xValue float64 = 0
			var yValue float64 = 0
			fmt.Sscanf(execTimeStr, "%f", &xValue)
			fmt.Sscanf(arrivalTimeStr, "%f", &yValue)

			processes = append(processes, Process{name: name, executionTime: xValue, arrivalTime: yValue})
		}
		sort.Sort(ByArrivalTime(processes))
		inputChan <- processes
		close(inputChan)
		wg.Done()

	}()
	go func() {
		for tuple := range inputChan {
			processLen := len(tuple)
			var waitingTime, turnaroundTime float64 = 0, 0
			var finishingTime float64 = 0
			i := 0
			var timer float64 = 0
			for {
				if i == len(tuple) {
					break
				}

				if tuple[i].executionTime > quantum {
					var temp = tuple[i]
					temp.executionTime = tuple[i].executionTime - quantum
					tuple[i].executionTime = quantum
					timer += quantum
					var index int = -1
					j := 0
					for j = i + 1; j < len(tuple); j++ {
						if tuple[j].arrivalTime > timer {
							index = j
							break
						}
					}
					if index == -1 {
						tuple = append(tuple, temp)
					} else {
						tuple = append(tuple[:index+1], tuple[index:]...)
						tuple[index] = temp
					}
				}
				i += 1

			}
			for i, item := range tuple {
				i = i
				finishingTime += item.executionTime
				value, key := maxFinishingTime[item.name]
				if key {
					maxFinishingTime[item.name] = math.Max(value, finishingTime)
					arrTime[item.name] = item.arrivalTime
					exTime[item.name] += item.executionTime
				} else {
					maxFinishingTime[item.name] = finishingTime
					arrTime[item.name] = item.arrivalTime
					exTime[item.name] = item.executionTime
				}
			}

			for key, finish := range maxFinishingTime {
				var arr = arrTime[key]
				turnaroundTime += finish - arr
				waitingTime += finish - exTime[key] - arr

			}

			var averageWaitingTime, averageTurnaroundTime float64 = waitingTime / float64(processLen), turnaroundTime / float64(processLen)
			//fmt.Printf("RR average WT and ATA (quantum = %f):\n", quantum)
			//
			//fmt.Println(averageWaitingTime, averageTurnaroundTime)
			result := fmt.Sprintf("RR average WT and TAT (quantum = %f): %0.3f, %0.3f", quantum, averageWaitingTime, averageTurnaroundTime)
			processChan <- result
			close(processChan)
			wg.Done()

		}
	}()

	for result := range processChan {
		fmt.Println(result)
	}

}
