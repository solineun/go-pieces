package mymap

func WordCount(data []string) map[string]int {
	words := make(map[string]int)

	for _, w := range data {
		if _, ok := words[w]; ok {
			words[w] += 1
		} else {
			words[w] = 1
		}
	}
	return words
}