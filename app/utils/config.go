package utils

import (
	"errors"
	"io/ioutil"
	"path"
	"strings"

	"github.com/go-yaml/yaml"
)

// ErrConfigMissed - 配置丢失
var ErrConfigMissed = errors.New("config is missed")

// 配置内容
var _config = map[interface{}]interface{}{}

// LoadConfig 读取配置目录
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

		name := file.Name()
		name = name[0 : len(name)-len(path.Ext(name))]
		c := make(map[interface{}]interface{})

		err = yaml.Unmarshal(content, &c)
		if err != nil {
			return err
		}

		_config[name] = c
	}

	Info("Config loaded success")

	return nil

}

// Config 获取 {name} 配置内容. name 支持 . 操作 比如 app.name
func Config(name string) (value interface{}, err error) {

	parts := strings.Split(name, ".")

	var v interface{} = _config
	for _, part := range parts {
		v = v.(map[interface{}]interface{})[part]
		if v == nil {
			err = ErrConfigMissed
			break
		}
	}

	value = v

	return
}

// ConfigInt 获取  {name} 的数字值
func ConfigInt(name string) (num int, err error) {
	v, err := Config(name)

	num = v.(int)
	return
}

// ConfigString 获取 {name} 的字符串值
func ConfigString(name string) (str string, err error) {
	v, err := Config(name)
	str = v.(string)
	return
}

// ConfigMap 获取 {name} 的字典配置
func ConfigMap(name string) (m map[interface{}]interface{}, err error) {
	v, err := Config(name)
	m = v.(map[interface{}]interface{})
	return
}
