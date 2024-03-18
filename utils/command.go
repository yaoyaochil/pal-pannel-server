package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"os/exec"
)

// ExecuteCommand 函数执行终端命令并返回执行结果和可能的错误
func ExecuteCommand(command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)

	// 执行命令并获取输出
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("执行命令失败: %v", err)
	}

	return string(output), nil
}

// ExecuteCommandWithOutputToFile 函数执行终端命令并将输出写入到文件中
func ExecuteCommandWithOutputToFile(command, outputFile string) error {
	// 创建文件以保存命令输出
	file, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("无法创建输出文件: %v", err)
	}
	defer file.Close()

	// 设置命令的输出为文件
	cmd := exec.Command("bash", "-c", command)
	cmd.Stdout = file

	// 执行命令
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("执行命令失败: %v", err)
	}

	return nil
}

// StreamLogFile 函数将文件的内容流式输出到http响应中
func StreamLogFile(filePath string, c *gin.Context) {
	file, err := os.Open(filePath)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "无法打开文件",
		})
		return
	}
	defer file.Close()

	fmt.Println(file.Name())

	c.Stream(func(w io.Writer) bool {
		_, err := io.Copy(w, file)
		if err != nil {
			return false
		}
		return true
	})
}
