package controllers

import (
	"fmt"
	"time"
)

func showFmt() {
	// var arr [3]int
	// arr[0] = 1
	// arr[1] = 2
	// arr[2] = 3
	// fmt.Println(arr[0])
	// fmt.Println(arr[1])
	// fmt.Println(arr[2])

	// var splie = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	// fmt.Println(splie[len(splie)-1:])

	// var numbers map[string]int
	// numbers = make(map[string]int)
	// numbers["one"] = 1
	// numbers["two"] = 2
	// numbers["three"] = 3
	// fmt.Println("第三个数字是：", numbers["three"])
	// fmt.Println(len(numbers))

	//fmt.Println(add(3, 5))

	//fmt.Println(swap("Hello", "Wored"))

	//fmt.Println(split(17))

	// v := 42 // change me!
	// fmt.Printf("v is of type %T\n", v)

	// fmt.Print("Gor runs on ")
	// switch os := runtime.GOOS; os {
	// case "darwin":
	// 	fmt.Println("OS X")
	// case "linux":
	// 	fmt.Println("Linux")
	// default:
	// 	fmt.Println(os)
	// }

	// today := time.Now()
	// fmt.Println(today.Weekday())

	// i, j := 42, 2701

	// p := &i         // point to i
	// fmt.Println(*p) // read i through the pointer
	// *p = 21         // set i through the pointer
	// fmt.Println(i)  // see the new value of i

	// p = &j         // point to j
	// *p = *p / 37   // divide j through the pointer
	// fmt.Println(j) // see the new value of j
	// V := Vertex{1, 2}
	// fmt.Println(V)
	// //V.x = 4
	// fmt.Println(V.x)

	// a := 100
	// b := 200
	// fmt.Println(max(a, b))

	// num := []int{3, 4, 5}
	// sum := 0
	// for _, val := range num {
	// 	sum += val
	// }
	// fmt.Println(sum)

	//sum(1, 2, 3)
	// nums := []int{1, 2, 3, 4, 5}
	// sum(nums...)
	// v := Vertex{x: 2, y: 4}
	// cs(v)

	// arg, err := f2(42)
	// fmt.Println(arg)
	// fmt.Println(err)

	// f("direct")

	// // To invoke this function in a goroutine, use
	// // `go f(s)`. This new goroutine will execute
	// // concurrently with the calling one.
	// go f("goroutine")

	// // You can also start a goroutine for an anonymous
	// // function call.
	// go func(msg string) {
	// 	fmt.Println(msg)
	// }("going")

	// // Our two function calls are running asynchronously in
	// // separate goroutines now, so execution falls through
	// // to here. This `Scanln` code requires we press a key
	// // before the program exits.
	// var input string
	// fmt.Scanln(&input)
	// fmt.Println("done")

	// message := make(chan string)
	// go func() { message <- "cscscs" }()
	// msg := <-message
	// fmt.Println(msg)

	// done := make(chan bool, 1)
	// go worker(done)
	// <-done

	// str := []string{"c", "a", "b"}
	// sort.Strings(str)
	// fmt.Println(str)

	// num := []int{3, 6, 1}
	// sort.Ints(num)
	// fmt.Println(num)

	// cs := Circel{2, 3, 7}
	// fmt.Println(cs.area())

	// cs := Circel{2, 3, 7}
	// str := show(cs)
	// fmt.Println(str)

	fmt.Println("test")
}

//Shape is cs
type Shape interface {
	area() float64
}

//Circel is
type Circel struct {
	x, y, redius float64
}

func (circel Circel) area() float64 {
	return circel.x * circel.y
}

func show(shape Shape) float64 {
	return shape.area()
}

func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d-%s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

//Vertex 测试注释。
type Vertex struct {
	x int
	y int
}

type measure interface {
	numadd() int
}

func (v Vertex) numadd() int {
	return v.x + v.y
}

func cs(c measure) {
	fmt.Println(c.numadd())
}

func sum(nums ...int) {
	fmt.Println(nums, " ")
	totle := 0
	for _, num := range nums {
		totle += num
	}
	fmt.Println(totle)
}

//max 比较两个值的大小
func max(num1, num2 int) int {
	var result int
	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

/*相加函数*/
func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}
