package system

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func isChromeRunning() (bool, error) {
	var cmd *exec.Cmd
	var output []byte
	var err error

	switch os := runtime.GOOS; os {
	case "windows":
		cmd = exec.Command("tasklist", "/FI", "IMAGENAME eq chrome.exe")
	case "linux":
		cmd = exec.Command("pgrep", "chrome")
	case "darwin":
		cmd = exec.Command("pgrep", "Google Chrome")
	default:

		return false, fmt.Errorf("不支持的操作系统")
	}

	if output, err = cmd.Output(); err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok && exitErr.ExitCode() == 1 {
			return false, nil // 命令执行成功但没有找到匹配的进程
		}
		return false, err // 命令执行出错
	}
	if strings.Contains(string(output), "chrome.exe") || strings.Contains(string(output), "Google Chrome") {
		return true, nil
	}
	return false, nil

}
func Run() {
	if running, err := isChromeRunning(); err == nil {
		if running {
			fmt.Println("检测到 Chrome 正在运行，关闭浏览器后重新运行Searchall")
			os.Exit(0)
		} else {

			// 继续其他操作
		}
	} else {

		// 错误处理逻辑
	}
}
