package main

import "fmt"
import "math"

func add(x int, y int) int {
	return x + y
}

func add2(x, y int) int {
	fmt.Printf("add2!")
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

// named result
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var x, y, z int = 1, 2, 3
var c, python, java = true, false, "no!"
const Pi = 3.14

// 숫자형 상수. 정밀한 값
const (
	Big = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int { return x * 10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

type Vertex struct {
	X int
	Y int
}

func adder() func(int) int {
    sum := 0
    return func(x int) int {
        sum += x
        return sum
    }
}

type Vertex3 struct {
	X float64
	Y float64
}
// 포인터 리시버: 1. 복사를 하지 않기위해, 2. 값을 수정하기 위해
func (v *Vertex3) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type Abser interface {
	Abs() float64;
}

type Example1 struct {}

func (ex1 *Example1) Start() {
	fmt.Printf("hello, world\n")
	fmt.Println(add2(42, 13))

	a, b := swap("hello", "world")
	fmt.Println(a,b)

	fmt.Println(split(17))
	a1, b1 := split(17)
	fmt.Println(a1, b1)

	fmt.Println(x, y, z, c, python, java)

	// 함수안에서 := 사용해서 var 와 명시적인 타입 생략
	c,d,e := 1,2,3
	fmt.Println(c,d,e)

	const World = "안녕"
	fmt.Println("hello", World)
	fmt.Println("happy", Pi, "Day")

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
	// fmt.Println(needInt(Big))

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}

	fmt.Println(sum)

	sum2 := 1
	for sum2 < 1000 {
		sum2 += sum2
	}

	fmt.Println(sum2)

	// 무한루프
	// for {
	// }

	fmt.Println(sqrt(2), sqrt(-4))

	fmt.Println(pow(3, 2, 10), pow(3, 3, 20))
	
	fmt.Println(Vertex{1, 2})

	p := Vertex{1, 2}
	q := &p
	q.X = 1e9
	fmt.Println(p)

	v := new(Vertex)
	fmt.Println(v)
	v.X, v.Y = 11, 9
	fmt.Println(v)

	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("s ==", s)

	for i := 0; i < len(s); i++ {
		fmt.Printf("s[%d] == %d\n", i, s[i])
	}

	fmt.Println("p[1:4] ==", s[1:4]) // [1, 4-1]

	fmt.Println("p[:3] ==", s[:3]) // [0, 3-1] 

	fmt.Println("p[4:] ==", s[4:]) // [4, len(s) -1] 

	// a := make([]int , 5)

	pow := []int{1,2,4,8,16,32,64,128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	type Vertex2 struct {
		Lat, Long float64
	}

	var m map[string]Vertex2

	m = make(map[string]Vertex2)
	m["Bell Labs"] = Vertex2 {
		40.68433, -74.39967,
	}

	var m2 = map[string]Vertex2{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}
	fmt.Println(m2["Bell Labs"])

	pos, neg := adder(), adder()
    for i := 0; i < 10; i++ {
        fmt.Println(
            pos(i),
            neg(-2*i),
        )
	}
	
	v3 := &Vertex3{3, 4}
	fmt.Println(v3.Abs())

	var ifa Abser
	ifa = v3

	fmt.Println(ifa.Abs())
}
// tour of go
// import ("name" "name")

// 첫문자가 대문자면 export name 이 됨.
// math.Pi

