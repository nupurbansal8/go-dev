package main

import (
	"bufio"
	"fmt"
	"os"
	//"strings"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	x := scanner.Text()
	scanner.Scan()
	y, _ := strconv.Atoi(scanner.Text())
	//z := strings.Split(y, ",")
	fmt.Println(x)
	fmt.Println(y)
}
