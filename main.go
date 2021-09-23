package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

//maxCPU is a functions that finds the highest possible gomaxprocs value by comparing the number of cpus.
func maxCPU() int {
	maxProcs := runtime.GOMAXPROCS(0) //default value
	numCPU := runtime.NumCPU()
	if maxProcs < numCPU {
		return maxProcs
	}
	return numCPU
}

//does the pinging as well as formats the result so that only the url and stats come out as the result
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

//generates a scatter plot from the values given by XY pairs
func plotGraph(values plotter.XYs){
	//defining a new plot using the gonum package
	p := plot.New()
	p.Add(plotter.NewGrid())

	//plotting of the values into a scatter plot
	s, err := plotter.NewScatter(values)
	if err != nil{
		panic(err)
	}
	p.Add(s)

	//saves the image of the plot as a png to the local folder
	if err := p.Save(3*vg.Inch, 3*vg.Inch, "graph.png"); err != nil{
		panic(err)
	}
}

//parallelizes ping command using go routine
func main() {
	urlList := os.Args[1:]//list of urls defined by user inputted args

	out := make(chan string, len(urlList))

	fmt.Println("GOMAXPROCS max value", maxCPU())

	//this tests runtime of the program with default value of GOMAXPROCS up to max value of GOMAXPROCS
	gmpMax := maxCPU()
	var durationList = make(map[int]int64)//list to store and display the duration of the pings according to the value of GOMAXPROCS
	for i := 0; i <= gmpMax; i++ {
		fmt.Println("\nTesting GOMAXPROCS value = ", i)
		runtime.GOMAXPROCS(i)
		start := time.Now()
		for range urlList {
			go ping(out, urlList)
		}
		duration := time.Since(start)
		durationList[i] = duration.Microseconds()
		fmt.Println("time: ", duration)
	}
	fmt.Println("\nFinal list of durations in microseconds for the code according to the value of GOMAXPROCS")
	fmt.Println(durationList)
	// creating a list of XY pairs to feed as data points into the scatter plot
	values := make(plotter.XYs, len(durationList))
	var i = 0
	// separating the key-value pairs into XY pairs from durationList
	for k, v := range durationList {
		values[i].X = float64(k)
		values[i].Y = float64(v)
		fmt.Println(values[i])
		i++
	}
	plotGraph(values)
}
