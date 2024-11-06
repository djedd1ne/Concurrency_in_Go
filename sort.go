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

	for _, str := range s {
		num, _ := strconv.Atoi(str)
		a = append(a, num)
	}

	const par = 4
	n := int(math.Max(math.Ceil(float64(len(s))/float64(par)), 1))
}
