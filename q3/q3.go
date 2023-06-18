package q3

import "fmt"

//Você está subindo uma escada. Leva n passos para chegar ao topo.
//
//A cada vez, você pode subir 1 ou 2 degraus. De quantas maneiras distintas você pode subir até o topo?

func ClimbStairs(n int) int {
	if n <= 2 {
		return n
	}
	degraus := make(map[int]int, n+1)
	for i := 3; i <= n; i++ {
		degraus[i] = degraus[i-1] + degraus[i-2]
	}
	fmt.Println(degraus)
	return degraus[n]
}
