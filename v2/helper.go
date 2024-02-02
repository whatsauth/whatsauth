package whatsauth

func removeDuplicateValues(stringSlice []string) (list []string) {
	keys := make(map[string]bool, len(stringSlice))
	list = make([]string, 0, len(stringSlice))

	for _, entry := range stringSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func DuplicateRemover[T comparable](list []T) (newList []T) {
	newMap := make(map[T]struct{}, len(list))

	for _, entry := range list {
		newMap[entry] = struct{}{}
	}

	for key := range newMap {
		newList = append(newList, key)
	}

	return newList
}
