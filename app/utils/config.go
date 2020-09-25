package utils

import (
	"fmt"
	"io/ioutil"
	"path"

	"github.com/go-yaml/yaml"
)

// 配置内容
var _config = map[string]interface{}{}

// Init 读取配置目录
func LoadConfig(dir string) error {
	fileInfoArr, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, file := range fileInfoArr {
		if file.IsDir() {
			continue
		}

		content, err := ioutil.ReadFile(path.Join(dir, file.Name()))
		if err != nil {
			return err
		}

		baseName := path.Split(file.Name())
		fmt.Println("Basename", baseName)
		c := make(map[string]interface{})

		err = yaml.Unmarshal(content, &c)
		if err != nil {
			return err
		}

		_config[file.Name()] = c

	}

	return nil

}

// Config 获取 {name} 配置内容. name 支持 . 操作 比如 app.name
func Config(name string) (value interface{}) {
	fmt.Println(_config)
	return nil
}
