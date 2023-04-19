package main

type Process struct {
	name          string
	executionTime float64
	arrivalTime   float64
}
type ByArrivalTime []Process

func (a ByArrivalTime) Len() int           { return len(a) }
func (a ByArrivalTime) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByArrivalTime) Less(i, j int) bool { return a[i].arrivalTime < a[j].arrivalTime }

func main() {
	// Test the algorithms one by one. Comment the algorithm you don't need.

	//FCFS()
	RR()
}
