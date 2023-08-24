package flagsearch

import (
	"fmt"
	"regexp"
)

func processUserInputString(input string) string {
	// 判断是否为字符串，如果是，则将其转换为正则表达式格式

	return fmt.Sprintf("(?i)(?:%s?\\s*[=:])\\s*([\\S]+)", regexp.QuoteMeta(input))

}

func processUserRegexesString(inputList []string) []string {
	var regexList []string

	for _, input := range inputList {
		regex := processUserInputString(input)
		regexList = append(regexList, regex)
	}

	return regexList
}
