package xirangrikui

import (
	"bufio"
	"fmt"
	"os"
	"searchall3.5/jiexi"
	"strings"
)

func ProcessFastCodeHistory(path string, fastCodeHistoryChan chan<- string) error {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("无法打开文件:", err)
		return nil
	}
	defer file.Close()

	bufScanner := bufio.NewScanner(file)

	count := 0 //计数器初始为0
	for bufScanner.Scan() {
		line := bufScanner.Text()
		if strings.HasPrefix(line, "fastcodehistroy=") {
			count++

			content := strings.SplitN(line, "=", 2)[1]
			decodedValue, err := jiexi.ProcessFastCodeHistroy(content)
			if err != nil || decodedValue == "" {
				continue
			}
			fastCodeHistoryChan <- strings.TrimSpace(decodedValue)

		}
	}
	fmt.Printf("\n找到向日葵历史识别记录，保存在search.txt中（共%d次）\n", count)
	if err := bufScanner.Err(); err != nil {
		return nil
	}
	return nil
}
