package main

import (
    "runtime"
)

func main() {

    runtime.GOMAXPROCS(runtime.NumCPU())
    
    for i := 0; i < runtime.NumCPU(); i++ {
        go func() {
            for {

            }
        }()
    }

    for {

    }
}