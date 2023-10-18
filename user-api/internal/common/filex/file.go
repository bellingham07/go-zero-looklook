package filex

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type Config struct {
	Ptuningv2 bool `yaml:"ptuningv2"`
}

// SaveUploadFileToLocal 上传文件
func SaveUploadFileToLocal(file *multipart.FileHeader, dst string) error {
	fmt.Println("++++++++++++++++++++++++++++", dst)
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	if err = os.MkdirAll(filepath.Dir(dst), 0750); err != nil {
		return err
	}

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}

// a-100,a-200,a-300,获取a-300
func getLastFolder() {
	// 指定包含checkpoint文件夹的目录
	directoryPath := "D:\\hibug\\test" // 将路径替换为实际路径
	//targetPath := "D:\\hibug\\test2"

	// 读取目录中的所有子目录
	subdirectories, err := os.ReadDir(directoryPath)
	if err != nil {
		fmt.Println("无法读取目录：", err)
		return
	}

	// 存储checkpoint文件夹的名称
	checkpoints := []string{}

	// 遍历子目录，筛选出checkpoint文件夹
	for _, subdir := range subdirectories {
		if subdir.IsDir() && strings.HasPrefix(subdir.Name(), "checkpoint-") {
			checkpoints = append(checkpoints, subdir.Name())
		}
	}

	// 按名称对checkpoint文件夹进行排序
	sort.Strings(checkpoints)

	// 检查是否有任何checkpoint文件夹
	if len(checkpoints) == 0 {
		fmt.Println("没有找到任何checkpoint文件夹")
		return
	}

	// 选择最新的checkpoint文件夹（最后一个）
	latestCheckpoint := checkpoints[len(checkpoints)-1]
	fmt.Printf("选择最新的checkpoint文件夹：%s\n", latestCheckpoint)

	//err = moveFile(targetPath, directoryPath+"\\"+latestCheckpoint+"\\pytorch_model.bin", directoryPath+"\\"+latestCheckpoint+"\\config.json")
	//if err != nil {
	//	fmt.Println("文件copy失败")
	//	return
	//}

	parseYaml2()
}

// 移动文件
func moveFile(targetPath string, file1, file2 string) error {
	srcFile1, err := os.Open(file1)
	if err != nil {
		return err
	}
	defer srcFile1.Close()

	srcFile2, err := os.Open(file2)
	if err != nil {
		return err
	}
	defer srcFile2.Close()

	destFile1, err := os.Create(targetPath + "\\pytorch_model.bin")
	if err != nil {
		return err
	}
	defer destFile1.Close()

	destFile2, err := os.Create(targetPath + "\\config.json")
	if err != nil {
		return err
	}
	defer destFile2.Close()

	_, err = io.Copy(destFile1, srcFile1)
	_, err = io.Copy(destFile2, srcFile2)
	if err != nil {
		return err
	}
	return nil
}

// 根据结构体解析
func parseYaml1() {
	path := "D:\\hibug\\test\\checkpoint-5000\\config.yaml"
	yamlData, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("无法读取文件：%v", err)
	}

	// 解析YAML数据到配置结构
	var config Config
	err = yaml.Unmarshal(yamlData, &config)
	if err != nil {
		log.Fatalf("无法解析YAML：%v", err)
	}

	// 修改配置项
	config.Ptuningv2 = true

	// 将修改后的配置重新编码为YAML
	modifiedYAML, err := yaml.Marshal(&config)
	if err != nil {
		log.Fatalf("无法编码为YAML：%v", err)
	}

	// 将修改后的YAML写回文件
	err = os.WriteFile(path, modifiedYAML, 0644)
	if err != nil {
		log.Fatalf("无法写回文件：%v", err)
	}
	fmt.Println("已修改配置项 'p' 为 'true'")
}

// 使用map进行解析
func parseYaml2() {
	path := "D:\\hibug\\test\\checkpoint-5000\\config.yaml"
	yamlData, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("无法读取文件：%v", err)
	}

	// 解析YAML数据到一个 map[string]interface{}
	var yamlMap map[string]interface{}
	err = yaml.Unmarshal(yamlData, &yamlMap)
	if err != nil {
		log.Fatalf("无法解析YAML：%v", err)
	}

	// 修改特定的配置项
	yamlMap["ptuningv2"] = true

	// 将修改后的YAML数据重新编码为YAML
	modifiedYAML, err := yaml.Marshal(&yamlMap)
	if err != nil {
		log.Fatalf("无法编码为YAML：%v", err)
	}

	// 将修改后的YAML写回文件
	err = os.WriteFile(path, modifiedYAML, 0644)
	if err != nil {
		log.Fatalf("无法写回文件：%v", err)
	}

	fmt.Println("已修改配置项 'p' 为 'true'")
}
