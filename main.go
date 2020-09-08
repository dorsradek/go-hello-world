package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"
)

func main() {
	var x = 3
	fmt.Println(x)
	y := 4
	fmt.Println(y)
	z := 8

	if y > x {
		yString := strconv.Itoa(y)
		fmt.Println(strconv.FormatInt(int64(z), 2))
		fmt.Println(strconv.FormatInt(int64(z), 10))
		fmt.Printf("x=%d", x)
		fmt.Printf("y=%s", yString)
	}

	var list [4]string
	list[3] = "three"
	fmt.Println(list)
	list2 := [3]string{"onu", "dos", "tres"}
	fmt.Println(list2)

	var radekDors = Person{
		firstName: "Radek",
		lastName:  "Dors",
	}
	fmt.Println(radekDors)

	now := time.Now()
	fmt.Println(now)
	fmt.Println(time.Now().Location())
	fmt.Println(time.Now().Clock())
	fmt.Println(time.Now().Date())
	fmt.Println(time.Now().ISOWeek())
	fmt.Println(time.Now().UTC())

	fmt.Println("Hello world")
	checkTime()
	writeToFile()
	writeToFileIoUtils()
	maps()
	loopFor()
	loopWhile()
	loopRange()
	sum := sum(1, 2)
	fmt.Printf("sum=%d\n", sum)
	multiple1, err := multiply(2, 3)
	fmt.Printf("mutiply 2*3=%d error=%s\n", multiple1, err)
	multiple2, err2 := multiply(-1, 3)
	fmt.Printf("mutiply 2*3=%d error=%s\n", multiple2, err2)
	structs()
	memoryAddress()
	fmt.Printf("CURRENCY=%s\n", CURRENCY)
	switches()
	matrix()
	slices()
	sumSlice()
	rangeString()

	variadicFunction(3, 2, 1)
	nums := []int{6, 7, 8}
	variadicFunction(nums...)
	returnFunction := returnFunction()(2)
	fmt.Println(returnFunction)

	closures()
	recursion()
	pointers()
	structs2()
	tom := animal{name: "tom", breed: "cat"}
	fmt.Println(tom.title())
	fmt.Println(tom.name)
	fmt.Println(tom.titlePointer())
	fmt.Println(tom.name)
	interfaces()
	errores(5)
	errores(-3)
	goroutines()
	waitGroups()
	waitGroup2()
	goMaxProcs()
}

// don't create goroutines in the libraries
// be careful to close goroutine because it can be memory leak
// to detect race condition: go run -race main.go
// `go` keyword on front of the function call
// when using anonymous function, pass data as a parameter because it is safer and cleaner
// synchronisation:
//   sync.WaitGroup{}
//     wg.Add(1) - number of goroutines
//     wg.Done() - signal that it is done for 1 goroutine
//     wg.Wait() - block
//   sync.RWMutex{}
//     RLock() RUnlock() - read lock
//     Lock() Unlock() - write lock

// runtime.GOMAXPROCS(-1) - number of available threads

func goMaxProcs() {
	fmt.Printf("GOMAXPROCS=%d\n", runtime.GOMAXPROCS(-1))
	fmt.Printf("NumCPU=%d\n", runtime.NumCPU())
	fmt.Printf("NumGoroutine=%d\n", runtime.NumGoroutine())
	fmt.Printf("NumCPU=%d\n", runtime.Version())
}

var wg = sync.WaitGroup{}
var m = sync.RWMutex{}
var counter = 0

func waitGroup2() { // no synchronisation between goroutines
	for i := 0; i < 10; i++ {
		wg.Add(2)
		m.RLock()
		go sayHola(i)
		m.Lock()
		go inc()
	}
	wg.Wait()
}

func inc() {
	counter++
	m.Unlock()
	wg.Done()
}

func sayHola(i int) {
	fmt.Printf("Hola %d %d\n", i, counter)
	m.RUnlock()
	wg.Done()
}

func waitGroups() {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(1)
	function := func() {
		fmt.Println("Hello")
		waitGroup.Done()
	}
	go function()
	waitGroup.Wait()
}

func sayHello() {
	fmt.Println("Say Hello!")
}

func goroutines() {
	sayHello()
	go sayHello()
}

func returnTypeAndError(i int) (int, error) {
	if i < 0 {
		return -1, errors.New(fmt.Sprintf("below zero=%d", i))
	}
	return 10 * i, nil
}

func errores(i int) {
	result, err := returnTypeAndError(i)
	if err != nil {
		fmt.Printf("error=%s\n", err)
	}
	fmt.Printf("result=%d\n", result)
}

type geometry interface {
	sum() int
	print()
}

type square struct {
	a int
}

func (s square) sum() int {
	return 4 * s.a
}

func (s square) print() {
	fmt.Println("asdadsada")
}

func calcSum(g geometry) {
	sum := g.sum()
	fmt.Println(sum)
	g.print()
}

func interfaces() {
	square := square{a: 3}
	calcSum(square)
}

type animal struct {
	name  string
	breed string
}

func (a animal) title() string {
	a.name = "mr " + a.name
	return a.name
}

func (a *animal) titlePointer() string {
	a.name = "mr " + a.name
	return a.name
}

func createDog(name string) animal {
	return animal{name: name, breed: "dog"}
}

func createDogPointer(name string) *animal {
	return &animal{name: name, breed: "dog"}
}

func structs2() {
	dog := createDog("tom")
	fmt.Println(dog)
	fmt.Println(&dog)
	dogPointer := createDogPointer("ben")
	fmt.Println(dogPointer)
	fmt.Println(*dogPointer)
	fmt.Println(&dogPointer)
}

func pointers() {
	i := 2
	println(i)
	withoutPointer(i)
	println(i)
	println(&i)
	withPointer(&i) // &i - gives memory address
	println(i)
}

func withoutPointer(i int) {
	println(i)
	i += 10
}

func withPointer(i *int) { // *int pointer
	println(i)
	*i += 10 // *i - dereference pointer from its memory address to the current value at that address
}

func recursion() {
	result := factorial(3)
	println(result)
}

func factorial(i int) int {
	if i == 0 {
		return 1
	}
	return i * factorial(i-1)
}

func closures() {
	function := func(a int) string {
		return "string " + strconv.Itoa(a)
	}

	result := function(45)
	println(result)
}

func returnFunction() func(int) string {
	return func(i int) string {
		return test(i)
	}
}

func test(number int) string {
	return "Asd " + strconv.Itoa(number)
}

func variadicFunction(nums ...int) {
	for index, value := range nums {
		fmt.Println(index, value)
	}
}

func rangeString() {
	println("rangeString")
	s := "test string"
	for index, value := range s {
		fmt.Println(index, strconv.QuoteToASCII(string(value)))
	}
}

func slices() {
	sl := make([]string, 3)
	sl[0] = "first"
	fmt.Printf("slice=%v\n", sl)
}

func sumSlice() {
	sl := []int{1, 2, 3, 4}
	var sum int
	for index, value := range sl {
		fmt.Println(index)
		sum += value
	}
	fmt.Printf("sum=%d\n", sum)
}

func matrix() {
	matrix := [2][3]int{}
	matrix[1][1] = 666
	fmt.Printf("matrix=%v\n", matrix)
}

func switches() {
	i := 5

	switch i {
	case 1, 2, 3:
		fmt.Println("case 1, 2, 3")
	case 4, 5, 6:
		fmt.Println("case 4, 5, 6")
	default:
		println("default")
	}
}

const CURRENCY = "GBP"

func memoryAddress() {
	a := 30
	fmt.Println(a)
	fmt.Println(&a)
	incrementCopy(a)
	fmt.Println(a)
	fmt.Println(&a)
	incrementPointer(&a)
	fmt.Println(a)
	fmt.Println(&a)
}

func incrementCopy(number int) { // passing copy
	number++
}

func incrementPointer(number *int) { // passing pointer
	*number++
}

func structs() {
	p := person{name: "Radek", age: 30}
	fmt.Printf("person=%v\n", p)
}

type person struct {
	name string
	age  int
}

func multiply(a int, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, errors.New("negative ")
	}
	return a * b, nil
}

func sum(a int, b int) int {
	return a + b
}

func loopRange() {
	arr := []string{"a", "b", "c"}
	for index, value := range arr {
		fmt.Printf("index=%d value=%s\n", index, value)
	}

	m := make(map[string]int)
	m["a"] = 1
	m["b"] = 2
	for key, value := range m {
		fmt.Printf("key=%s value=%v\n", key, value)
	}
}

func loopWhile() {
	i := 0
	for i < 3 {
		fmt.Printf("i=%d\n", i)
		i++
	}
}

func loopFor() {
	for i := 0; i < 3; i++ {
		fmt.Printf("i=%d\n", i)
	}
}

type Dupa map[string]int

func maps() {
	mapa := make(Dupa)
	mapa["a"] = 2
	mapa["b"] = 3
	delete(mapa, "a")

	mapa2 := make(map[string]int)
	mapa2["a"] = 2
	mapa2["b"] = 3
	delete(mapa2, "b")

	fmt.Printf("mapa=%v\n", mapa)
	fmt.Printf("mapa2=%v\n", mapa2)
}

func (dupa Dupa) String() string {
	return "asdasd"
}

func writeToFile() {
	file, err := os.Create("/Users/Radoslaw.Dors/workspace/go-hello-world/test.txt")
	if err != nil {
		fmt.Printf("err=%s\n", err)
	}
	fmt.Fprintf(file, "asdasd")
}

func writeToFileIoUtils() {
	text := "first line\nsecond line"
	textByteSlice := []byte(text)
	err := ioutil.WriteFile("/Users/Radoslaw.Dors/workspace/go-hello-world/test-utils.txt", textByteSlice, 0666)
	if err != nil {
		fmt.Printf("err=%s\n", err)
	}
}

type Person struct {
	firstName string
	lastName  string
}

func (person Person) String() string {
	return fmt.Sprintf("firstName=%s lastName=%s age=%d", person.firstName, person.lastName, person.age())
}

func (person Person) age() int {
	return 10
}

func checkTime() {
	now := time.Now()
	fmt.Printf("Time since start=%s\n", time.Since(now))
}
