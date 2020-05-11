package conf

import (
	"errors"
	"io/ioutil"
	"strings"

	"gopkg.in/yaml.v2"
)

// 配置
type Conf struct {
	Db struct {
		Driver   string `yaml:"driver"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Database string `yaml:"database"`
		Charset  string `yaml:"charset"`
	}
	Server struct {
		Port string `yaml:"port"`
	}
}

// 将ymal文件中的内容进行加载
func (c *Conf) Load(path string) error {
	ext := c.guessFileType(path)
	if ext == "" {
		return errors.New("cant not load" + path + " config")
	}
	return c.loadFromYaml(path)
}

// 判断配置文件名是否为yaml格式
func (c *Conf) guessFileType(path string) string {
	s := strings.Split(path, ".")
	ext := s[len(s)-1]
	switch ext {
	case "yaml", "yml":
		return "yaml"
	}
	return ""
}

// 将配置yaml文件中的进行加载
func (c *Conf) loadFromYaml(path string) error {
	yamlS, readErr := ioutil.ReadFile(path)
	if readErr != nil {
		return readErr
	}
	// yaml解析的时候c.data如果没有被初始化，会自动为你做初始化
	err := yaml.Unmarshal(yamlS, &c)
	if err != nil {
		return errors.New("can not parse " + path + " config")
	}
	return nil
}
