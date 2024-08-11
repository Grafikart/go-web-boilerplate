package main

func removeItem[K comparable](s []K, item K) []K {
	pos := -1
	for k, v := range s {
		if v == item {
			pos = k
			continue
		}
	}
	if pos == -1 {
		return s
	}
	return removeAt(s, pos)
}

func removeAt[K comparable](s []K, i int) []K {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
