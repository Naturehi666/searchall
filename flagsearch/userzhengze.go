package flagsearch

func processUserInput(input string) string {
	return input
}

func processUserRegexes(inputList []string) []string {
	var regexList []string

	for _, input := range inputList {
		regex := processUserInput(input)
		regexList = append(regexList, regex)
	}

	return regexList
}
