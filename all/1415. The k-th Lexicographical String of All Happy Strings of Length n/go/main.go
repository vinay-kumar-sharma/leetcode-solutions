func getHappyString(n int, k int) string {
    chars := []rune{'a', 'b', 'c'}
	var solve func(i int, s string) string

	solve = func(i int, s string) string {
		if i == n {
			k--
			if k == 0 {
				return s
			}
			return ""
		}

		for _, c := range chars {
			if len(s) == 0 || rune(s[len(s)-1]) != c {
				rc := solve(i+1, s+string(c))
				if rc != "" {
					return rc
				}
			}
		}
		return ""
	}

	return solve(0, "")
}