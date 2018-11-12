package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	//	_ "github.com/lib/pq"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/", hello)
	router.HandleFunc("/hello", hello)
	router.HandleFunc("/fibonacci/{id}", getFibonacci).Methods("GET")
	router.HandleFunc("/fibonacci2/{id}", getFibonacci2).Methods("GET")
	router.HandleFunc("/fibonacci3/{id}", getFibonacci3).Methods("GET")
	router.HandleFunc("/countries", getCountries)

	log.Fatal(http.ListenAndServe(":8000", router))
}

func FibonacciRecursive(first int, second int, rank int) int {
	if rank == 1 {
		return first
	}

	if rank == 2 {
		return second
	}
	return FibonacciRecursive(first, second, rank-1) + FibonacciRecursive(first, second, rank-2)
}

func FibonacciGenerator(first int, second int) chan int {
	c := make(chan int)

	go func() {
		for i, j := first, second; ; j, i = i+j, j {
			c <- i
		}
	}()

	return c
}

func hello(w http.ResponseWriter, r *http.Request) {
	//input : nothing
	//output: Hello World
	//hello -  9.213µs on my MacBook pro.
	start := time.Now()
	json.NewEncoder(w).Encode("Hello World")
	fmt.Println("hello - ", time.Since(start))
}

func getFibonacci(w http.ResponseWriter, r *http.Request) {
	//input : the nth fibonnachi desired (ex : 10)
	//output: the reprsenting number of the nth fibonnachi (answer : 34)
	//loop version - good old procedural.  Sems to bust at 94th (wrong number / uint64 limit?)
	start := time.Now()
	params := mux.Vars(r)
	rank, _ := strconv.Atoi(params["id"])
	first := uint64(0)
	second := uint64(1)
	total := uint64(0)

	if rank == 1 || rank == 0 {
		total = 0
	}

	if rank == 2 {
		total = 1
	}

	for i := 3; i <= rank; i++ {
		total = first + second
		first = second
		second = total
		//debug
		//fmt.Println("loop fibonacci nth - ", i, " equals: ", second, " ", reflect.TypeOf(second), " ", time.Since(start))
	}

	json.NewEncoder(w).Encode(&total)
	fmt.Println("fibonacci - ", time.Since(start))
}

func getFibonacci2(w http.ResponseWriter, r *http.Request) {
	//input : the nth fibonnachi desired (ex : 10)
	//output: the reprsenting number of the nth fibonnachi (answer : 34)
	//recursive version - costly compared to loop hard for my MacBook pro to go over 45.
	//@48 fibonacci  -  10.033µs
	//@48 fibonacci2 -  10.606042737s  (loop 1000x faster than recurive on my McBook pro)
	start := time.Now()
	params := mux.Vars(r)
	rank, _ := strconv.Atoi(params["id"])
	total := 0

	if rank == 1 || rank == 0 {
		total = 0
		json.NewEncoder(w).Encode(total)
	} else if rank == 2 {
		total = 1
		json.NewEncoder(w).Encode(total)
	} else {
		json.NewEncoder(w).Encode(FibonacciRecursive(1, 2, rank-2)) //Pile away!
	}

	//json.NewEncoder(w).Encode(total)
	fmt.Println("fibonacci2 - ", time.Since(start))
}

func getFibonacci3(w http.ResponseWriter, r *http.Request) {
	//input : the nth fibonnachi desired (ex : 10)
	//output: the reprsenting number of the nth fibonnachi (answer : 34)
	//go routine version (multi processor)
	//inspird by : https://www.jonathan-petitcolas.com/2014/08/18/fibonacci-generator-in-go.html
	//@48 fibonacci  -  10.033µs
	//@48 fibonacci3 -  30.684µs (loop version 3x faster than recurive on my McBook pro).
	// i.e. thre is a loop to launch the go routines.
	// May be godd if more demanding like loop of loop (On2 'Order on n square' and up?)
	start := time.Now()
	params := mux.Vars(r)
	rank, _ := strconv.Atoi(params["id"])
	total := 0

	if rank == 1 || rank == 0 {
		total = 0
		json.NewEncoder(w).Encode(total)
	} else if rank == 2 {
		total = 1
		json.NewEncoder(w).Encode(total)
	} else {
		generator := FibonacciGenerator(1, 1)
		for i := 0; i < rank-1; i++ {
			total = (<-generator)

		}
	}
	json.NewEncoder(w).Encode(total)
	fmt.Println("fibonacci3 - ", time.Since(start))
}

func getCountries(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	json.NewEncoder(w).Encode(params)
}
