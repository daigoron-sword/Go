package main

import (
	"fmt"
)

// func main() {
// 	num01 := 123
// 	var num02 int = 1234567890
// 	num03 := 1.23
// 	var num04 float64 = 1.23456789
// 	fmt.Println(reflect.TypeOf(num01))
// 	fmt.Println(reflect.TypeOf(num02))
// 	fmt.Println(reflect.TypeOf(num03))
// 	fmt.Println(reflect.TypeOf(num04))
// }

// func main() {
// 	a := 10
// 	b := 1
// 	num_bool := a < b
// 	fmt.Println(num_bool)
// 	fmt.Println(reflect.TypeOf(num_bool))
// }

// 配列
// func main() {
// a := [3]string{"sato", "suzuki", "takahashi"}
// 要素数の省略
// a := [...]string{"sato", "suzuki", "takahashi"}
// a[1] = "tnaka"

// fmt.Println(a[0])
// fmt.Println(a[1])
// fmt.Println(a[2])

// a := [2][2]string{{"sato", "suzuki"}, {"takahashi", "tanaka"}}

// fmt.Println(a[0][0])
// fmt.Println(a[0][1])
// fmt.Println(a[1][0])
// fmt.Println(a[1][1])
//

// 算術演算子
// func main() {
// 	x := 10
// 	y := 2

// 	fmt.Println(x + y)
// 	fmt.Println(x - y)
// 	fmt.Println(x * y)
// 	fmt.Println(x / y)
// 	fmt.Println(x % y)
// }

// // 関係演算子
// func main() {
// 	x := 10
// 	y := 2

// 	fmt.Println(x > y)
// 	fmt.Println(x == y)
// 	fmt.Println(x != y)
// }

// 論理演算子
// func main() {
// 	x := 8
// 	y := 3

// 	// fmt.Println(x >= 5 && x <= 10)
// 	// fmt.Println(y >= 5 && y <= 10)
// 	fmt.Println(x == 5 || x == 10)
// 	fmt.Println(y == 5 || y == 10)
// }

// 複合代入演算子
// func main() {
// 	x := 8
// 	y := 12
// 	z := 20
// 	x += 10
// 	z += y

// 	fmt.Println(x)
// 	fmt.Println(z)
// }

// インクリメント＆デクリメント
// func main() {
// 	x := 8
// 	y := 8

// 	x++
// 	y--

// 	fmt.Println(x)
// 	fmt.Println(y)
// }

// if 文
// func main() {
// 	age := 22

// 	if age >= 20 {
// 		fmt.Println("adult")
// 	} else {
// 		fmt.Println("child")
// 	}
// }

// // else if
// func main() {
// 	age := 0

// 	if age >= 20 {
// 		fmt.Println("adult")
// 	} else if age == 0 {
// 		fmt.Println("baby")
// 	} else {
// 		fmt.Println("child")
// 	}
// }

// 簡易文if
// func main() {
// 	x := 10
// 	y := 12

// 	if age := x + y; age >= 20 {
// 		fmt.Println("adult")
// 	} else if age == 0 {
// 		fmt.Println("baby")
// 	} else {
// 		fmt.Println("child")
// 	}
// }

// 5以上10未満のif文
func main() {
	x := 10

	if x >= 5 && x < 10 {
		fmt.Println("5以上10未満")
	} else {
		fmt.Println("それ以外")
	}
}
