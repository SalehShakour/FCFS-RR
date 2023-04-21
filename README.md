# FCFS-RR Scheduling Algorithm

This is an implementation of the First-Come, First-Served (FCFS) and Round Robin (RR) scheduling algorithms in Go.

## Algorithms

### FCFS

The FCFS algorithm schedules processes in the order that they arrive. The implementation of the FCFS algorithm in this project assumes that all processes arrive at time 0.

### RR

The RR algorithm schedules processes in a round-robin fashion, giving each process a fixed time slice, or quantum, and then moving on to the next process. If a process does not complete in its time slice, it is moved to the end of the queue. 

## How to Use

1. Clone the repository.
2. Navigate to the root directory of the repository.
3. Run `go build` to build the executable.
4. Run the executable.
5. Enter the number of processes you would like to schedule.
6. Enter the details of each process as a string in the format `"process-name execution-time arrival-time"`. For example, `"P1 10 0"` represents a process named P1 with an execution time of 10 and an arrival time of 0.
7. If you want to use the FCFS algorithm, comment out the call to the `RR()` function in the `main()` function of the `main.go` file.
8. If you want to use the RR algorithm, comment out the call to the `FCFS()` function in the `main()` function of the `main.go` file.
9. If you are using the RR algorithm, you can change the time slice by modifying the `quantum` variable in the `rrScheduling()` function of the `scheduler.go` file.

## License

This project is licensed under the MIT License. See the `LICENSE` file for more information.

