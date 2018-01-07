package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

// 3
/// \
//9  20
///  \
//15  7

//Depth-first-search https://en.wikipedia.org/wiki/Depth-first_search
func averageOfLevels_1(root *TreeNode) []float64 {
	var res []float64
	queue := Queue{}
	queue = append(queue, *root)

	for !queue.Empty() {
		temp := Queue{}
		var sum, count int = 0, 0

		for !queue.Empty() {
			n := queue.Dequeue()
			sum += n.Val
			count++
			if n.Left != nil {
				temp.Enqueue(*n.Left)
			}
			if n.Right != nil {
				temp.Enqueue(*n.Right)
			}

		}

		queue = temp
		res = append(res, (float64(sum)*1.0)/float64(count))
	}
	return res
}

func averageOfLevels_2(root *TreeNode) []float64 {
	var res []float64
	var count []int

	avrg(root, 0, &res, &count)
	for i := 0; i < len(res); i++ {
		res = append(res, res[i]/float64(count[i]))
	}

	return res
}

func avrg(node *TreeNode, i int, sum *[]float64, count *[]int) {
	if node == nil {
		return
	}

	if i < len(*sum) {
		*sum = append(*sum, (*sum)[i]+float64(node.Val))
		*count = append(*count, (*count)[i]+1)
	} else {
		*sum = append(*sum, 1.0*float64(node.Val))
		*count = append(*count, 1)
	}
	avrg(node.Left, i+1, sum, count)
	avrg(node.Right, i+1, sum, count)
}
