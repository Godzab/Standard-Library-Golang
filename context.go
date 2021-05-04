package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	minLatency = 10
	maxLatency = 5000
	timeout = 3000
)

func main(){
	// Little program that searches flight routes
	// We are going to use a mock backend database
	// The purpose of this is to show how context can be used to propergate cancellation signals across go routines and processes.
	rootCtx := context.Background()
	ctx, cancel := context.WithTimeout(rootCtx, time.Duration(timeout) * time.Millisecond)
	defer cancel()

	//Listen for interrupt signal - When user manually cancels operation.
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sig
		fmt.Println("Aborting due to interrupt...")
		cancel()
		os.Exit(0)
	}()
	fmt.Println("Starting to search...")
	res, err := Search(ctx, "London", "Brisban")
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(res)

}

func Search(ctx context.Context,from, to string)([]string, error){
	//SlowSearch here
	// We need to watch for context.Done is closed.
	res := make(chan []string)
	go func() {
		res <- SlowSearch(from, to)
		close(res)
	}()

	// Wait for 2 events: either one will be result
	for {
		select{
			case dst := <-res:
				return dst, nil
			case <- ctx.Done():
				return []string{}, ctx.Err()
		}
	}
}

// SlowSearch is a very slow function that goes through a series of operations.
func SlowSearch(from, to string)[]string{
	rand.Seed(time.Now().Unix())
	latency := rand.Intn(maxLatency - minLatency + 1) - minLatency
	fmt.Printf("Started to search from %s to %s takes %dms..", from, to, latency)
	time.Sleep(time.Duration(latency) * time.Millisecond)
	return []string{from + "-" + to +"-British airways-11AM", from + "-" + to +"-Virgin airways-11AM"}
}