package flagsearch

import (
	"flag"
	"fmt"
	"searchall3.5/search"
	"searchall3.5/tuozhan/liulanqi"
	"searchall3.5/tuozhan/liulanqi/browser"
	"strings"
)

func Banner() {
	fmt.Println(`
 __                        
(_  _  _  __ _ |_  _  |  | 
__)(/_(_| | (_ | |(_| |  |
        verson:3.5.6
                     `)

}

func FlagSearchall() {
	Banner()

	searchPath := flag.String("p", "", "The path to search for files"+
		"\nexample: searchall.exe -p  C:\\\\")
	zipFlag := flag.Bool("z", false, "compress browsers result to zip"+
		"\nexample: searchall.exe -b all -z")
	browserFlag := flag.String("b", "", "available browsers: all|"+browser.Names()+
		"\nexample: searchall.exe -b all\n"+"360极速浏览器不支持360speedX版本\n")
	userRegexes := flag.String("r", "", "自定义正则"+
		"\nexample: searchall.exe -p C:\\\\ -r 正则1,正则2,正则3")
	stringRegexes := flag.String("s", "", "自定义字符串(可预编译成正则)"+
		"\nexample: searchall.exe -p C:\\\\ -s 字符串，字符串2，字符串3")
	userOnlyFlag := flag.Bool("u", false, "只使用自定义正则和字符串进行检索"+
		"\nexample: searchall.exe -p C:\\\\ -r 正则1,正则2,正则3 -u  ")

	flag.Parse()

	var userRegexList []string
	if *userRegexes != "" {
		inputs := strings.Split(*userRegexes, ",")
		userRegexList = processUserRegexes(inputs)
	}
	if *stringRegexes != "" {
		inputs := strings.Split(*stringRegexes, ",")
		userRegexList = processUserRegexesString(inputs)
	}

	if *searchPath != "" && *browserFlag == "" && *userRegexes == "" && *stringRegexes == "" {
		search.Searchall(*searchPath, nil, *userOnlyFlag)
	} else if *searchPath != "" && *browserFlag == "" && *userRegexes != "" && *stringRegexes == "" {
		search.Searchall(*searchPath, userRegexList, *userOnlyFlag)
	} else if *searchPath != "" && *browserFlag == "" && *userRegexes == "" && *stringRegexes != "" {
		search.Searchall(*searchPath, userRegexList, *userOnlyFlag)
	} else if *searchPath == "" && *browserFlag != "" {
		liulanqi.Execute(*browserFlag)
		if *zipFlag {
			err := liulanqi.CompressResult()
			if err != nil {
				fmt.Println(err)
			}
		}
	} else {
		flag.PrintDefaults()
		return
	}
}
