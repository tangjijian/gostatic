package recursion

func Test(num int) int {

	if num < 2 {
		return num
	} else {
		return Test(num-2) + Test(num-1)
	}
}
