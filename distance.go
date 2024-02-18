package levenshtein

// Distance between two strings is the number of deletions, insertions, or
// substitutions required to transform source string into target string.
func Distance(s, t string) int {
	a, b := []rune(s), []rune(t)
	u, v := vectors(len(b))
	for i, A := range a {
		v[0] = i + 1
		for j, B := range b {
			v[j+1] = min(v[j]+1, u[j+1]+1, u[j]+cost(A, B))
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
