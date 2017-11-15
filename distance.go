package lavenshtein

func Distance(s, t string) int {
	a, b := []rune(s), []rune(t)
	u, v := vectors(len(b))
	for i, A := range a {
		v[0] = i + 1
		for j, B := range b {
			v[j+1] = min3(v[j]+1, u[j+1]+1, u[j]+cost(A, B))
		}
		u, v = v, u
	}
	return u[len(u)-1]
}

func cost(a, b rune) int {
	if a != b {
		return 1
	}
	return 0
}

func vectors(n int) ([]int, []int) {
	u, v := make([]int, n+1), make([]int, n+1)
	for i := 0; i <= n; i++ {
		u[i] = i
	}
	return u, v
}

func min3(u, v, w int) int {
	return min(u, min(v, w))
}

func min(u, v int) int {
	if u < v {
		return u
	}
	return v
}
