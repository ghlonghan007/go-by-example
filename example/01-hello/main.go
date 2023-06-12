package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// fmt.Println("hellowolrd")
	name := flag.String("name", "world", "specify the name you want to say")
	flag.Parse()
	fmt.Println("os arg is:", os.Args)
	fmt.Println("input parameter is:", *name)
	fullString := fmt.Sprintf("hello %s form Go\n", *name)
	fmt.Println(fullString)
	//go run /workspaces/go-by-example/example/01-hello/main.go  -name longhan
	//os arg is: [/tmp/go-build2868223516/b001/exe/main -name longhan]
	// input parameter is: longhan
	// hello longhan form Go
}
