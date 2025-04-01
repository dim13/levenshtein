package levenshtein

// Distance between two strings is the number of deletions, insertions, or
// substitutions required to transform source string into target string.
func Distance(a, b string) int {
	return DistanceSlices([]rune(a), []rune(b))
}

// DistanceSlices between two slices is the number of deletions, insertions, or
// substitutions required to transform source slice into target slice.
// func DistanceSlices[S ~[]E, E comparable](a, b S) int {
func DistanceSlices[S ~[]E, E comparable](a, b S) int {
	n := len(b) + 1
	u, v := make([]int, n), make([]int, n)
	for i := range n {
		u[i] = i
	}
	for i, A := range a {
		v[0] = i + 1
		for j, B := range b {
			var cost int
			if A != B {
				cost = 1
			}
			v[j+1] = min(v[j]+1, u[j+1]+1, u[j]+cost)
		}
		u, v = v, u
	}
	return u[len(u)-1]
}
