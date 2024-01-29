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
