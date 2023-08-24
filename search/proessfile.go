package search

import (
	"bufio"
	"fmt"
	"os"
	"searchall3.5/tuozhan/xirangrikui"
	"strings"
)

var foundDockerOverlay2 bool

func ProcessFile(info os.FileInfo, path string, absPath string, resultChan chan []string, fastCodeHistoryChan chan string, errchan chan error) {
	if info.Name() == "config.ini" && strings.Contains(path, "SunloginClient") {
		fmt.Println("\n本系统安装了向日葵，配置路径为：", path)
		err := xirangrikui.ProcessFastCodeHistory(path, fastCodeHistoryChan)
		if err != nil {

			errchan <- err
			return
		}
		/*else if info.Name() == "passwd" && strings.Contains(absPath, "etc") {
			fileContent, err := ioutil.ReadFile(absPath)
			if err != nil {
				errchan <- err
				return
			}
			// 加入到结果中
			resultChan <- []string{fmt.Sprintf("File: %s\n%s\n", absPath, fileContent)}
			fmt.Printf("\n读取File: %s\n", absPath)
		} else if info.Name() == "shadow" && strings.Contains(absPath, "etc") {
			fileContent, err := ioutil.ReadFile(absPath)
			if err != nil {
				errchan <- err
				return
			}
			// 加入到结果中
			resultChan <- []string{fmt.Sprintf("File: %s\n%s\n", absPath, fileContent)}
			fmt.Printf("\n读取File: %s\n", absPath)*/
	} else if info.Name() == "docker" && strings.Contains(absPath, "overlay2") {
		overlay2Index := strings.Index(absPath, "overlay2")
		if overlay2Index != -1 && !foundDockerOverlay2 {
			dockerOverlay2Path := absPath[:overlay2Index+len("overlay2")]
			fmt.Printf("\n本系统安装了docker，路径为：%s\n", dockerOverlay2Path)
			resultChan <- []string{fmt.Sprintf("docker path: %s\n", dockerOverlay2Path)}

			foundDockerOverlay2 = true
		}

	} else if info.Name() == "secure" && strings.Contains(absPath, "var/log") {
		file, err := os.Open(absPath)
		if err != nil {
			errchan <- err
			return
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		successPattern := "Accepted password for"
		var result strings.Builder
		successCount := 0

		for scanner.Scan() {
			line := scanner.Text()
			if strings.Contains(line, successPattern) {
				successCount++
				result.WriteString(line)
				result.WriteRune('\n')
			}
		}

		if err := scanner.Err(); err != nil {
			errchan <- err
			return
		}

		if successCount > 0 {

			resultChan <- []string{fmt.Sprintf("File: %s\n%s\n", absPath, result.String())}
			fmt.Printf("\n读取File: %s, 成功登录次数: %d\n", absPath, successCount)
		}
	}
}
