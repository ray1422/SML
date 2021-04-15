package utils

func RegexNamedGroupMap(matches []string, names []string) map[interface{}]string {
	if len(matches) == 0 || len(names) == 0 {
		return nil
	}
	result := make(map[interface{}]string)
	for i, name := range names {
		if i != 0 && name != "" {
			result[name] = matches[i]
		}
		result[i] = matches[i]
	}
	return result
}

func RegexNamedGroupIndexMap(idxes []int, names []string) map[interface{}][2]int {
	result := make(map[interface{}][2]int)
	for i := 0; i < len(idxes); i += 2 {
		name := names[int(i/2)]
		if i != 0 && name != "" {
			result[name] = [2]int{idxes[i], idxes[i+1]}
		}
		result[i] = [2]int{idxes[i], idxes[i+1]}
	}
	return result
}
