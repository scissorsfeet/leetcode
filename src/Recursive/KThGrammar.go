package Recursive

func kthGrammar(N int, K int) int {
	if N < 1 || K < 1 {
		return -1
	}
	//N = int(math.Log2(float64(K))) + 1
	K = transK(N) + K - 1
	return helper(N, K)
}

func transK(N int) int {
	K := 1
	for i:=1;i<=N-1;i++ {
		K = 2 * K
	}
	return K
}

func helper(N int, K int) int {
	if N == 1 {
		return 0
	}
	parentK := K / 2
	parentV := helper(N-1, parentK)
	var currentV int
	if parentV == 0 {
		if K % 2 == 0 {
			currentV = 0
		} else {
			currentV = 1
		}
	} else {
		if K % 2 == 0 {
			currentV = 1
		} else {
			currentV = 0
		}
	}
	return currentV
}