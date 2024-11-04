package main

import (
	"fmt"
  "bufio"
  
)

func main() {
	fmt.Print("Enter integers split by space")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	s := strings.Split(strings.TrimSpace(input), " ")
	a := make([]int, 0, len(s))
}
