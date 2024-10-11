package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Args struct {
	Limit int
}

type PrimeNumberServer int64

func (t *PrimeNumberServer) GetPrimeNumbers(args *Args, reply *[]int64) error {
	var primeList []int64
	for num := 2; num <= args.Limit; num++ {
		isPrime := true

		for i := 2; i*i <= num; i++ {
			if num%i == 0 {
				isPrime = false
				break
			}
		}

		if isPrime {
			primeList = append(primeList, int64(num))
		}
	}

	*reply = primeList
	return nil
}

func main() {
	primeNumberService := new(PrimeNumberServer)

	rpc.Register(primeNumberService)
	rpc.HandleHTTP()

	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listen error", err)
	}

	fmt.Println("[*] Server is running on port 1234...")
	http.Serve(l, nil)
}
