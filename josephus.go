package main

import (
	"fmt"
	"strconv"
)

func josephus(n int64) (int64, error) {
	temp := strconv.FormatInt(n, 2)
	bin := []rune(temp)
	noBin := string(bin[1:len(bin)]) + string(bin[0])
	no, err := strconv.ParseInt(noBin, 2, 62)
	if err != nil {
		return 0, err
	}
	return no, nil
}

func main() {
	var input int64
	fmt.Printf("Enter the no of soilders : ")
	fmt.Scanf("%d", &input)
	no, err := josephus(input)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Josephus place: %d\n", no)
}
