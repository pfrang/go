package main

func main() {
	src := []int{1, 2, 3, 4}
	dest := make([]int, len(src))
	copy(dest, src)
	println(dest[2])
}
