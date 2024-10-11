package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Args struct {
	Limit int64
}

func main() {
	var limit int64

	fmt.Println("Ini adalah program untuk menghasilkan list bilangan prima")
	fmt.Println("Silakan masukkan input di bawah sebagai batas terbesar angka")
	fmt.Printf("Input: ")
	fmt.Scan(&limit)

	var reply []int64
	args := Args{
		Limit: limit,
	}

	client, err := rpc.DialHTTP("tcp", "localhost"+":1234")
	if err != nil {
		log.Fatal("dialing", err)
	}

	err = client.Call("PrimeNumberServer.GetPrimeNumbers", args, &reply)
	if err != nil {
		log.Fatal("PrimeNumberServer error: ", err)
	}

	for _, val := range reply {
		fmt.Printf("%v ", val)
	}
}
