package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("welcome to Terminal_chat")
	reader := bufio.NewReader(os.Stdin)
	juicy,_,_ := reader.ReadLine()
	fmt.Printf("%s\n",juicy)
}
