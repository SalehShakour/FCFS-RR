# FCFS-RR Scheduling Algorithm

This is an implementation of the First-Come, First-Served (FCFS) and Round Robin (RR) scheduling algorithms in Go.

## Algorithms

### FCFS

The FCFS algorithm is the simplest CPU scheduling algorithm. It schedules processes in the order in which they arrive in the ready queue. When the CPU becomes available, the process at the front of the queue is selected for execution and is removed from the queue. The execution continues until the process completes or blocks for I/O. Then, the next process in the queue is selected for execution.

In FCFS, the waiting time for a process is simply the sum of the execution times of all the processes that arrived before it. The turnaround time is the sum of the waiting time and the execution time.

### RR

The RR algorithm is a preemptive scheduling algorithm that is designed for time-sharing systems. In RR, each process is assigned a fixed time slice, called a time quantum, and is scheduled for execution for that amount of time. After the time quantum expires, the process is preempted and moved to the back of the ready queue. The next process in the queue is then selected for execution, and the process continues until it completes or blocks for I/O.

The RR algorithm allows each process to run for a short period of time, providing good response time for interactive applications. However, if the time quantum is too small, there can be a lot of context switching overhead, which can slow down the system.

In RR, the waiting time for a process is the total time it spends in the ready queue before it gets to run. The turnaround time is the sum of the waiting time and the execution time. 

## How to Use

1. Clone the repository.
2. Navigate to the root directory of the repository.
3. Run `go build` to build the executable.
4. Run the executable.
5. Enter the number of processes you would like to schedule.
6. Enter the details of each process as a string in the format `"process-name execution-time arrival-time"`. For example, `"P1 10 0"` represents a process named P1 with an execution time of 10 and an arrival time of 0.
7. If you want to use the FCFS algorithm, comment out the call to the `RR()` function in the `main()` function of the `main.go` file.
8. If you want to use the RR algorithm, comment out the call to the `FCFS()` function in the `main()` function of the `main.go` file.
9. If you are using the RR algorithm, you can change the time slice by modifying the `quantum` variable in the `RR()` function of the `RR.go` file.



