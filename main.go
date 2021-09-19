package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
)

func ping(out chan string, urlList []string) {

	listLength := len(urlList)


	//This loop iterates through the list of urls which were passed in by the user, and running ping each time.
	//Note that this works for the Mac OS shell, and is untested for Windows and Linux.
	for i := 0; i < listLength; i++ {
		go func(i int) {
			tempUrl := urlList[i]
			result, err := exec.Command("ping", "-c 5", tempUrl).Output() //uses shell to ping a url 5 times
			if err != nil {
				fmt.Println(err.Error())
			} else {
				out <- string(result)
			}
		}(i)
	}
}

func main() {
	runtime.GOMAXPROCS(0)//default number of cores
	urlList := os.Args[1:]//list of urls defined by user inputted args

	out := make(chan string, len(urlList))
	start := time.Now()

	go ping(out, urlList)

	for range urlList {
		fmt.Println(<-out)
	}

	duration := time.Since(start)
	fmt.Println("time", duration)
}