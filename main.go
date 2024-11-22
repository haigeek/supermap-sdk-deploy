package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	// 文件夹路径
	dirPath := "D:\\data\\Downloads\\11.2.0"

	// Maven部署的基本命令模板
	mvnCmdTemplate := "mvn deploy:deploy-file -DgroupId=com.supermap.iobjects -DartifactId=%s -Dversion=11.2.0 -Dpackaging=jar -Dfile=%s -Durl=http://127.0.0.1:8080/artifactory/libs-release -DrepositoryId=central"

	// 遍历文件夹
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// 检查是否为文件且扩展名为.jar
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".jar") {
			// 获取文件名（不包括扩展名）
			artifactID := strings.TrimSuffix(info.Name(), ".jar")

			// 构建Maven命令参数
			filePath := path

			// 替换模板中的变量
			cmdStr := fmt.Sprintf(mvnCmdTemplate, artifactID, filePath)
			fmt.Println(cmdStr)

			// 执行Maven命令
			// 根据操作系统调整
			cmd := exec.Command("cmd", "/C", cmdStr)
			output, err := cmd.CombinedOutput()
			if err != nil {
				fmt.Printf("Error executing command for file %s: %v\n", filePath, err)
				fmt.Println(string(output))
			} else {
				fmt.Printf("Successfully deployed file %s\n", filePath)
			}
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error walking the path %s: %v\n", dirPath, err)
	}
}
