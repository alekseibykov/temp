package main

import "fmt"

func main() {
	fmt.Println("qwe")

	var slc []int

	for i := 1; i <= 100; i++ {
		slc = append(slc, i)
	}

	slcFirst10 := slc[0:10]
	slc = append(slcFirst10, slc[90:]...)

	var res []int
	for i := 1; i <= len(slc); i++ {
		res = append(res, slc[len(slc)-i])
	}
	fmt.Println(res)

	str := []rune("йцукен")
	resStr := ""
	for i := 1; i <= len(str); i++ {
		resStr += string(str[len(str)-i])
	}
	fmt.Println(resStr)
}
