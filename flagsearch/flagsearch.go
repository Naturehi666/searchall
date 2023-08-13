package flagsearch

import (
	"flag"
	"fmt"

	"searchall3.5/search"
	"searchall3.5/tuozhan/liulanqi"
	"searchall3.5/tuozhan/liulanqi/browser"
)

func Banner() {
	fmt.Println(`
 __                        
(_  _  _  __ _ |_  _  |  | 
__)(/_(_| | (_ | |(_| |  |
        verson:3.5.4
                     `)

}

func FlagSearchall() {

	Banner()

	searchPath := flag.String("p", "", "The path to search for files"+
		"\nexample: searchall3.5.exe -p  C:\\Users\\")
	zipFlag := flag.Bool("z", false, "compress browsers result to zip"+
		"\nexample: searchall3.5.exe -b all -z")
	browserFlag := flag.String("b", "", "available browsers: all|"+browser.Names()+
		"\nexample: searchall3.5.exe -b all\n"+"360极速浏览器不支持360speedX版本\n")

	flag.Parse()

	if *searchPath != "" && *browserFlag == "" {
		search.Searchall(*searchPath)
	} else if *searchPath == "" && *browserFlag != "" {
		liulanqi.Chromeall(*browserFlag)
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
