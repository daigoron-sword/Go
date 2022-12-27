package main

import "fmt"

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
// func main() {
// 	x := 10

// 	if x >= 5 && x < 10 {
// 		fmt.Println("5以上10未満")
// 	} else {
// 		fmt.Println("それ以外")
// 	}
// }

// 繰り返し
// func main() {
// 	for i := 0; i <= 4; i++ {
// 		fmt.Println(i)
// 	}
// }

// break
// func main() {
// 	for i := 0; i <= 4; i++ {
// 		if i == 3 {
// 			break
// 		}
// 		fmt.Println(i)
// 	}
// }

// continue
// func main() {
// 	for i := 0; i <= 4; i++ {
// 		if i == 3 {
// 			continue
// 		}
// 		fmt.Println(i)
// 	}
// }

// nest
// func main() {
// 	for i := 0; i <= 4; i++ {
// 		for j := 0; j <= 2; j++ {
// 			fmt.Println(i, "-", j)
// 		}
// 	}
// }

// 配列
// func main() {
// 	arr := [...]int{2, 4, 6, 8, 10}

// 	for i := 0; i <= 4; i++ {
// 		fmt.Println(arr[i])
// 	}
// }

// for+sumで足し上げ
// func main() {
// 	arr := [...]int{2, 4, 6, 8, 10}
// 	sum := 0

// 	for i := 0; i <= 4; i++ {
// 		sum += arr[i]
// 	}
// 	fmt.Println(sum)
// }

// for文省略形
// func main() {
// 	i := 0
// 	for {
// 		fmt.Println(i)
// 		if i == 4 {
// 			break
// 		}
// 		i++
// 	}
// }

// for文省略形２
// func main() {
// 	i := 0
// 	for i <= 4 {
// 		fmt.Println(i)
// 		i++
// 	}
// }

// 1から始まり10で終わるfor文を、3通りの方法で記述
// func main() {
// 	for i := 1; i <= 10; i++ {
// 		fmt.Println(i)
// 	}
// }

// func main() {
// 	i := 1
// 	for i <= 10 {
// 		fmt.Println(i)
// 		i++
// 	}
// }

// func main() {
// 	i := 1
// 	for {
// 		if i == 11 {
// 			break
// 		}
// 		fmt.Println(i)
// 		i++
// 	}
// }

// 関数定義
// func sayHello(greeting string) {
// 	fmt.Println(greeting)
// }

// func main() {
// 	sayHello("Good morning")
// }

// 引数定義した関数
// func cal(x int) {
// 	fmt.Println(x * 3)
// }

// func main() {
// 	cal(6)
// }

// 複数定義した関数
// func cal(x, y int) {
// 	fmt.Println(x / y)
// }

// func main() {
// 	cal(6, 3)
// }

// 戻り値
// func cal(x, y int) int {
// 	return (x / y)
// }

// func main() {
// 	result := cal(6, 3)
// 	fmt.Println(result)
// }

// 戻り値複数
// func cal(x, y int) (int, int) {
// 	return (x / y), (x * y)
// }

// func main() {
// 	result01, result02 := cal(6, 3)
// 	fmt.Println(result01, result02)
// }

// 関数内で指定した変数を返り値に指定
// func cal(x, y int) (int, int) {
// 	a := x / y
// 	b := x * y
// 	return a, b
// }

// func main() {
// 	result01, result02 := cal(10, 5)
// 	fmt.Println(result01, result02)
// }

// 戻り値指定の省略
// func cal(x, y int) (a int, b int) {
// 	a = x / y
// 	b = x * y
// 	return
// }

// func main() {
// 	result01, result02 := cal(10, 5)
// 	fmt.Println(result01, result02)
// }

// 関数を変数に代入
// func main() {
// 	hello := func(greeting string) {
// 		fmt.Println(greeting)
// 	}

// 	hello("Good morinig")
// }

// 無名関数
// func main() {
// 	func(greeting string) {
// 		fmt.Println(greeting)
// 	}("Good evening")
// }

// 確認問題→正解
// func cal(x, y int) (sum int) {
// 	sum = x + y
// 	return
// }

// func main() {
// 	fmt.Println(cal(10, 5))
// }

// 構造体定義
// type Student struct {
// 	name          string
// 	math, english float64
// }

// func main() {
// 	var s Student
// 	s.name = "sato"
// 	s.math = 80
// 	s.english = 70

// 	fmt.Println(s)
// }

// 初期化と同時にフィールド値を代入
// type Student struct {
// 	name          string
// 	math, english float64
// }

// func main() {
// 	s := Student{name: "sato", math: 80, english: 70}
// 	fmt.Println(s)
// }

// 確認
// type User struct {
// 	gender string
// 	age    int
// }

// func main() {
// 	var u User
// 	u.gender = "male"
// 	u.age = 20
// 	fmt.Println(u)
// }

// func main() {
// 	u := User{gender: "male", age: 20}
// 	fmt.Println(u)
// }

// メソッド定義方法
// type Student struct {
// 	name          string
// 	math, english float64
// }

// 関数の中に構造体を記述することで、フィールドを使用可能
// (s Student)→レシーバと呼ぶ
// func (s Student) avg() {
// 	fmt.Println(s.name, (s.math+s.english)/2)
// }

// func main() {
// 	a001 := Student{"sato", 80, 70}
// 	a001.avg()
// }

// メソッドに引数を渡す場合
// type Student struct {
// 	name          string
// }

// func (s Student) avg(math, english float64) {
// 	fmt.Println((math + english) / 2)
// }

// func main() {
// 	a001 := Student{name: "sato"}
// 	a001.avg(80, 70)
// }

// 戻り値があるメソッドの定義
// type Student struct {
// 	name string
// }

// func (s Student) avg(math, english float64) (avgResult float64) {
// 	avgResult = (math + english) / 2
// 	return
// }

// func main() {
// 	a001 := Student{name: "sato"}
// 	fmt.Println(a001.avg(80, 70))
// }

// 確認問題→結果○
// type User struct {
// 	name string
// }

// func (u User) cal(weight, height float64) (result float64) {
// 	result = weight / height / height * 10000
// 	return
// }

// func main() {
// 	a001 := User{"daichi"}
// 	fmt.Println(a001.name, a001.cal(73, 172))
// }

type Student struct {
	name string
}

func (s Student) CalAvg(data []float64) (avgResult float64) {
	sum := 0.0
	for i := 0; i < len(data); i++ {
		sum += data[i]
	}
	avgResult = sum / float64(len(data))
	return
}

func (s Student) judge(avg float64) (judgeResult string) {
	if avg >= 60 {
		judgeResult = "passed"
	} else {
		judgeResult = "failed"
	}
	return
}

func main() {
	a001 := Student{"sato"}
	data := []float64{70, 65, 50, 10, 30}
	var avg float64 = a001.CalAvg(data)
	result := a001.judge(avg)
	fmt.Println(avg)
	fmt.Println(a001.name + " " + result)
}
