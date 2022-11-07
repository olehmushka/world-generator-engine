package tools

func GetKeyWithGreatestValue[K string, V int](in map[K]V) K {
	var (
		key K
		max V
	)
	for k, v := range in {
		if v > max {
			key = k
			max = v
		}
	}

	return key
}
