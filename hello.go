package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	// fmt.Println("hello world!")
	// var num1 int = 20
	// var num2 = 21
	// num3 := 22
	// fmt.Printf("num1 is %d, num2 is %d, num3 is %d\n", num1, num2, num3)

	// var (
	// 	name = "maple"
	// 	sex  = "male"
	// 	age  = 20
	// )
	// fmt.Printf("name is %s, sex is %s, age is %d\n", name, sex, age)

	// const (
	// 	CONSTA = "aa"
	// 	CONSTB = 20
	// )
	// fmt.Printf("consta is %s, constb is %d\n", CONSTA, CONSTB)

	// // 数值类型
	// // bool
	// boolA := true
	// boolB := false
	// fmt.Println(boolA, boolB)
	// // int8 uint8
	// var intA int8 = -128
	// var uintA uint8 = 255
	// fmt.Println(intA, uintA)
	num1 := rand.Int()
	fmt.Printf("num1 is %d\n", num1)
	// t := time.Now()
	// tFloat := t.Unix()
	rand.Seed(time.Now().Unix())
	for i := 0; i < 10; i++ {
		randNum := rand.Intn(10)
		fmt.Println(randNum)
	}

	var arrayA [5]int
	for key := range arrayA {
		arrayA[key] = key
	}
	for key, value := range arrayA {
		fmt.Println("key is: ", key, "value is: ", value)
	}
	fmt.Println("length of arrayA is: ", len(arrayA), "capacity of arrayA is: ", cap(arrayA))
	arrayB := [5]int{}
	fmt.Println(arrayB)
	for key, value := range arrayB {
		arrayB[key] = key + value
	}
	fmt.Println(arrayB)
	sliceA := arrayB[1 : len(arrayB)-1]
	fmt.Println(sliceA)
	sliceB := make([]int, 5)
	sliceC := append(sliceA, 5, 5, 6, 78688)
	sliceC[2] = 50
	sliceA[1] = 3
	fmt.Println(arrayB, &arrayB[0])
	fmt.Println(sliceA, &sliceA[0])
	fmt.Println(sliceB, &sliceB[0])
	fmt.Println(sliceC, &sliceC[0])

	mapA := make(map[string]string)
	mapA["a"] = "aaa"
	mapA["b"] = "bbb"
	fmt.Println(mapA)
	value, ok := mapA["a"]
	if ok{
		fmt.Println(value)
	}else{
		fmt.Println("cannot get value")
	}

	delete(mapA, "a")
	fmt.Println(mapA)

	mapA["c"] = "ccc"
	for key, value := range mapA {
		fmt.Println(key, ":", value)
	}

	stringA := "my world"
	fmt.Println(strings.Contains(stringA, "a"))
	fmt.Println(strings.Count(stringA, "a"))

	funcTmp := func (){
		fmt.Println("匿名函数")
	}
	funcTmp()
	pension1 := Persion{
		name: "张三",
		age:  20,
	}
	fmt.Println(pension1.name, pension1.age)

	my := struct {
		name string
		age int
	}{
		"llisi",
		21,
	}
	fmt.Println(my.name, my.age)

}

type Persion struct {
	name string
	age int
}

type USB interface {
	create()
	delete()
	make()
}

func (p Persion) create() {
	fmt.Println("aaa")
}

funcTmp()