package main

import (
	"fmt"
	"math/big"
)

var level_max int = 0
var level int = 0

func persistance(n big.Int) {
	zero := big.NewInt(0)
	ten := big.NewInt(10)
	checking_if_exists := new(big.Int).Set(&n)
	if exists_0(*checking_if_exists) == true {
		fmt.Println("returned")
		return
	}

	persistance_num := big.NewInt(1)
	for n.Cmp(zero) == 1 {
		digit := new(big.Int).Mod(&n, ten)
		persistance_num.Mul(persistance_num, digit)
		n.Div(&n, ten)
	}

	fmt.Println(persistance_num)
	//temp := persistance_num
	if persistance_num.Cmp(ten) == 1 {
		level++
		if level_max < level {
			level_max = level
		}
		persistance(*persistance_num)
	} else {

		fmt.Println("done")

	}

}

func exists_5(n big.Int) bool {
	five := big.NewInt(5)
	zero := big.NewInt(0)
	ten := big.NewInt(10)

	for n.Cmp(zero) == 1 {
		digit := new(big.Int).Mod(&n, ten)
		if digit.Cmp(five) == 0 {
			return true
		}

		n.Div(&n, ten)
	}

	return false
}

func exists_0(n big.Int) bool {
	zero := big.NewInt(0)
	ten := big.NewInt(10)

	for n.Cmp(zero) == 1 {
		digit := new(big.Int).Mod(&n, ten)
		if digit.Cmp(zero) == 0 {
			return true
		}
		n.Div(&n, ten)
	}

	return false
}

func main() {
	num := big.NewInt(277777788888899)
	temp := new(big.Int).Set(num)
	checking_if_5_exists := new(big.Int).Set(num)
	one := big.NewInt(1)

	for level_max <= 12 {
		level = 0

		fmt.Println(temp)
		persistance(*temp)
		fmt.Println("Levels deep ", level+1)
		fmt.Println("\n")

		num.Add(num, one)

		temp.Set(num)
		checking_if_5_exists.Set(num)
	}

	fmt.Println("Levels deep ", level_max+1)

}
