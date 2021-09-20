# MP0 - Gary Yip

# Goal of the Project
This project serves as an introduction into using Go. Most importantly, it highlights the use of go channels and go routines to achieve things sychronously and in parallel.
There are two main goals for this program. First is to ping ip addresses in parallel and to display and observe the statistics outputted from the ping. Second is to compare the results using GOMAXPROCS and see how it affects performance of the program.

# How to run the program
This program can be run in the terminal. Most importantly, it is set up to run on the Mac OS, but is untested for Windows or Linux systems. It can be run using go run main.go as shown below.

![Screen Shot 2021-09-19 at 6 08 01 PM](https://user-images.githubusercontent.com/70530925/133944551-cddb7107-1d2a-435c-9bfb-d63ebf04d39c.png)

The program takes in user input through the command line, simply put the urls after the run command separated by a space.

![Screen Shot 2021-09-19 at 6 10 26 PM](https://user-images.githubusercontent.com/70530925/133944594-153eb52f-ac53-4ec8-8cc3-69bfe1fa3eda.png)

# The Code

The program utilizes the os/exec golang package, which allows the use of shell commands in our program. Because this code was written on a Mac, it may use some arguments to the ping command that may not exist in Windows or Linux, specifically the -c argument to define how many times to ping.

The code also makes use of go-routines to accomplish the task in parallel. All of the results are printed after the ping has concluded and the program is finished running.

# Resources Used

This code was written with the help of the following resources.

Finding max parallelism https://gist.github.com/peterhellberg/5848304

Use of goroutines https://gobyexample.com/goroutines

Using shell commands to run ping https://stackoverflow.com/questions/6182369/exec-a-shell-command-in-go

String functions in go https://gobyexample.com/string-functions
