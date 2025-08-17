package main

import (
	"effectiveGo/test"
	"fmt"
	"io"
	"os"
	"sort"
)

func main() {
	student := test.Student()
	student.Age = 1
	student.Name = "a"
	test.SetStudent(student)
	fmt.Println(student)
	//
	//f, err := os.Open("test.txt")
	//d, err := f.Stat()
	//fmt.Println(d, err)

	var oldMap map[string]string = make(map[string]string)
	var newMap map[string]string = make(map[string]string)
	oldMap["a"] = "1"
	oldMap["b"] = "2"
	for k, v := range oldMap {
		newMap[k] = v
	}
	fmt.Println(newMap)

	arr := []int{1, 2, 3}
	for k, v := range arr {
		fmt.Println(k, v)
	}

	for pos, char := range "日本\x80語" {
		fmt.Printf("character %#U starts at byte position %d\n", char, pos)
	}

	var c bool = true
	for n := 0; n < 3; n++ {
		switch {
		case arr[n] > 1:
			if c {
				break
			}
		default:
			fmt.Println("hhh")
		}
	}

	var t interface{}
	t = arr
	switch t := t.(type) {
	case bool:
		fmt.Println("bool", t)
	case []int:
		fmt.Println("[]int", t)
	default:
		fmt.Println("default", t)
	}

	bytes := []byte{'a', 'b', '3', 'b', '5', 'a'}
	b := bytes[0:5]
	x := 0
	for i := 0; i < len(b); {
		x, i = nextInt(b, i)
		fmt.Println(x, i)
	}

	_, err := Contents("aaa")
	fmt.Println(err)

	bb(i1)

	studentp := new(test.Demo)
	studentp.Name = "aa"
	fmt.Println(*studentp)

	var studentp1 *test.Demo = &test.Demo{Name: "a", Age: 1}
	var studentp2 *test.Demo = &test.Demo{Name: "a"}
	var studentp3 *test.Demo = new(test.Demo)
	fmt.Println(studentp1, studentp2, studentp3)

	var ap *[]int = new([]int)
	fmt.Println(*ap)

	var farr = []float64{1.0, 3.5, 4.3, 2.7, 5}
	fmt.Println(Sum(&farr))
	fmt.Println(farr[0])

	picture := make([][]uint8, 3)
	for i := 0; i < len(picture); i++ {
		picture[i] = make([]uint8, i+1)
	}
	fmt.Println(picture)

	picture = make([][]uint8, 4)
	pixels := make([]byte, 4*3)
	pixels = []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	for i := range picture {
		picture[i], pixels = pixels[:3], pixels[3:]
	}
	fmt.Println(picture)
	fmt.Println(ByteSize(1).String())

	var sq = Sequence{3, 2, 1}
	println(sq.String())
}

type Sequence []int

func (s Sequence) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s Sequence) Len() int {
	return len(s)
}

func (s Sequence) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Sequence) String() string {
	sort.Sort(s)
	str := "["
	for i, v := range s {
		if i > 0 {
			str += " "
		}
		str += fmt.Sprint(v)
	}
	return str + "]"
}

type ByteSize float64

const (
	_           = iota
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
)

func (b ByteSize) String() string {
	switch {
	case b >= TB:
		return fmt.Sprintf("%.2fTB", b/TB)
	case b >= GB:
		return fmt.Sprintf("%.2fGB", b/GB)
	case b >= MB:
		return fmt.Sprintf("%.2fMB", b/MB)
	case b >= KB:
		return fmt.Sprintf("%.2fKB", b/KB)
	}
	return fmt.Sprintf("%.2fB", b)
}

func Sum(a *[]float64) (sum float64) {
	for _, v := range *a {
		sum += v
	}
	(*a)[0] = 0
	return
}

func nextInt(b []byte, i int) (int, int) {
	for ; i < len(b) && !isDigit(b[i]); i++ {
	}
	x := 0
	for ; i < len(b) && isDigit(b[i]); i++ {
		x = x*10 + int(b[i]) - '0'
	}
	return x, i
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

// Contents returns the file's contents as a string.
func Contents(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		defer fmt.Println("defer")
		return "", err
	}
	defer f.Close() // f.Close will run when we're finished.

	var result []byte
	buf := make([]byte, 100)
	for {
		n, err := f.Read(buf[0:])
		result = append(result, buf[0:n]...) // append is discussed later.
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err // f will be closed if we return here.
		}
	}
	return string(result), nil // f will be closed if we return here.
}

var i1 = 0

func trace(s int) int {
	fmt.Printf("trace: %d\n", s)
	return s
}
func un(s int) {
	fmt.Printf("leaving:%d\n", s)
}
func aa(i int) {
	defer un(trace(i))
	fmt.Println("in a")
}
func bb(i int) {
	defer un(trace(i))
	fmt.Println("in b")
	i++
	aa(i)
}
