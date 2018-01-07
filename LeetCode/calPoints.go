package main

import "strconv"

func main() {
}

func calPoints(ops []string) int {

	var res int
	st := Stack{}

	for _, ch := range ops {
		switch ch {
		case "+":
			res1 := st.Pop()
			res2 := st.Pop()

			st.Push(res2)
			st.Push(res1)
			st.Push(res1 + res2)

		case "D":
			res := st.Pop()
			st.Push(res)
			st.Push(res * 2)
		case "C":
			st.Pop()
		default:
			num, _ := strconv.Atoi(ch)
			st.Push(num)
		}
	}

	for _, num := range st {
		res += num
	}

	return res

}
