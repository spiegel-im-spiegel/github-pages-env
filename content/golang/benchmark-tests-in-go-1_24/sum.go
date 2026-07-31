package sample

// func Sum(data []int) int {
// 	total := 0
// 	for _, value := range data {
// 		total += value
// 	}
// 	return total
// }

func Sum(data []int) int {
	total := 0
	l := len(data)
	h := l / 2
	for i := range h {
		total += data[i] + data[l-i-1]
	}
	if h*2 != l {
		total += data[h]
	}
	return total
}

func Add(a, b int) int {
	return a + b
}
