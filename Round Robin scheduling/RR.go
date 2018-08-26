package main

func main() {

}

func findWaitingTime(processes []int, n int, bt []int, wt []int, quantum int) {
	remBt := make([]int, n)

	for i := 0; i < n; i++ {
		remBt[i] = bt[i]
	}

	t := 0

	for {
		done := true
		for i := 0; i < n; i++ {
			if remBt[i] > 0 {
				done = false
				if remBt[i] > quantum {
					t += quantum
					remBt[i] -= quantum
				}
			} else {
				t = t + remBt[i]
				wt[i] = t - bt[i]
				remBt[i] = 0
			}
		}

		if done == true {
			break
		}
	}
}

func findTurnAroundTime(processes []int, n int, bt []int, wt []int, quantum int) {

}
