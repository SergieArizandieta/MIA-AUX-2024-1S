package main 

import (
	"bufio"
	"os"
	"fmt"
 )
 
 func main(){
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() // use `for scanner.Scan()` to keep reading
	line := scanner.Text()
	fmt.Println("captured:",line)
 }