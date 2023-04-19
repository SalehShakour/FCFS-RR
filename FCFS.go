package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"sync"
)

func FCFS() {

	processChan := make(chan string)
	inputChan := make(chan []Process)

	var averageWaitingTime, averageTurnaroundTime float64

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		reader := bufio.NewReader(os.Stdin)

		// Read the counter from the user
		fmt.Print("Enter the number of processes: ")
		var count int
		_, err := fmt.Scanln(&count)
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}

		var processes []Process
		for i := 0; i < count; i++ {
			//fmt.Printf("Enter process %d of %d: ", i+1, count)
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
			//name := parts[0]
			execTimeStr := parts[1]
			arrivalTimeStr := parts[2]

			var xValue float64 = 0
			var yValue float64 = 0
			fmt.Sscanf(execTimeStr, "%f", &xValue)
			fmt.Sscanf(arrivalTimeStr, "%f", &yValue)

			processes = append(processes, Process{executionTime: xValue, arrivalTime: yValue})
		}
		sort.Sort(ByArrivalTime(processes))
		inputChan <- processes
		close(inputChan)
		wg.Done()

	}()

	go func() {
		for processes := range inputChan {
			//fmt.Println(processes)
			var waitingTime, turnaroundTime float64 = 0, 0
			var finishingTime float64 = 0
			for i, item := range processes {
				finishingTime += item.executionTime
				i = i
				//fmt.Println(i, item.executionTime, item.arrivalTime, finishingTime)
				turnaroundTime += finishingTime - item.arrivalTime
				waitingTime += finishingTime - item.executionTime - item.arrivalTime
			}
			averageWaitingTime, averageTurnaroundTime = waitingTime/float64(len(processes)), turnaroundTime/float64(len(processes))
			result := fmt.Sprintf("FCFS average WT and ATA: %0.3f, %0.3f", averageWaitingTime, averageTurnaroundTime)
			processChan <- result
			close(processChan)
			wg.Done()

		}

	}()

	for result := range processChan {
		fmt.Println(result)
	}

}
