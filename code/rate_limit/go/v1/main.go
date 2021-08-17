package main

import (
	"fmt"
	"net"
	"os"
	"container/ring"
	"sync/atomic"
	"sync"
	"time"
)


var (
	limitCount int = 10
	limitBucket int = 6
	curCount  int32 = 0
	head *ring.Ring
)


func main(){
	tcpAddr, err := net.ResolveTCPAddr("tcp4", "0.0.0.0:9000")
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	defer listener.Close()

	head = ring.New(limitBucket)
	for i := 0; i < limitBucket; i ++ {
		head.Value = 0
		head = head.Next()
	}

	go func(){
		timer := time.NewTicker(time.Second * 1)
		for range timer.C {
			subCount := int32(0 - head.Value.(int))
			newCount := atomic.AddInt32(&curCount, subCount)

			arr := [6]int{}
			for i := 0; i < limitBucket; i++ {
				arr[i] = head.Value.(int)
				head = head.Next()
			}

			fmt.Println("move subCount,newCount,arr", subCount, newCount, arr)
			head.Value = 0
			head = head.Next()
		}
	}()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handle(&conn)
	}
}


func handle(conn *net.Conn){
	defer (*conn).Close()
	n :=  atomic.AddInt32(&curCount, 1)

	//fmt.Println("handler n: ", n)
	if n > int32(limitCount){
		atomic.AddInt32(&curCount, -1)
		(*conn).Write([]byte("HTTP/1.1 404 NOT FOUND\r\n\r\nError, too many request, please try again."))
	}else{
		mu := sync.Mutex{}
		mu.Lock()
		pos := head.Prev()
		val := pos.Value.(int)
		val ++
		pos.Value = val
		mu.Unlock()
		time.Sleep(1 * time.Second)
		(*conn).Write([]byte("HTTP/1.1 200 OK\r\n\r\nI can change the world!"))
	}
}

func checkError(err error){
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}



