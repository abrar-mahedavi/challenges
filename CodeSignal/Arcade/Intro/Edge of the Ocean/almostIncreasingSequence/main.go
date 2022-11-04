package main

func almostIncreasingSequence(sequence []int) bool {
	var last, lastPrev, c = -0x186a0, -0x186a0, 0
	for _, value := range sequence {
		if value <= last {
			c++
			if value > lastPrev {
				last = value
				// continue
			}
			continue
		}
		lastPrev = last
		last = value
	}
	return c <= 1
}


func main() {
	fmt.Printf("%v", almostIncreasingSequence(3))
}
