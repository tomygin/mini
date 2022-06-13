package info

import (
	"errors"
	"io/ioutil"

	yaml "gopkg.in/yaml.v3"
)

//获取配置文件的信息
func InfoMap() (map[string]interface{}, error) {
	infofile, err := ioutil.ReadFile("./info.yaml")
	if err != nil {
		return nil, errors.New("打开配置文件失败")
	}
	result := make(map[string]interface{})
	err = yaml.Unmarshal(infofile, &result)
	if err != nil {
		return nil, errors.New("解析配置文件失败")
	}
	return result, nil
}

//info yaml 对应的结构
type Info struct {
	Sql struct {
		Name     string `yaml:"name"`
		Password string `yaml:"password"`
		DbName   string `yaml:"dbname"`
		Port     string `yaml:"port"`
		Ip       string `yaml:"ip"`
	}
	Host struct {
		Ip   string `yaml:"ip"`
		Port string `yaml:"port"`
	}
	Email struct {
		Account string `yaml:"account"`
		Kay     string `yaml:"kay"`
	}
	Jwt struct {
		Salt string `yaml:"salt"`
	}
}

//获取配置信息
func InfoStruct() (Info, error) {
	info := Info{}
	infofile, err := ioutil.ReadFile("./info.yaml")
	if err != nil {
		return info, errors.New("打开配置文件失败")
	}
	err = yaml.Unmarshal(infofile, &info)
	if err != nil {
		return info, errors.New("解析配置文件失败")
	}
	return info, nil
}

var Allinfo Info

// func init() {
// 	all, err := InfoStruct()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	// fmt.Println(All)
// 	All = all
// }

//过滤的词汇
type Filter struct {
	Words []string `yaml:"words"`
}

//读取过滤的信息
func FilterWords() ([]string, error) {
	filterwords := Filter{}
	file, err := ioutil.ReadFile("./filter.yaml")
	if err != nil {
		return filterwords.Words, errors.New("打开过滤词汇文件失败")
	}
	err = yaml.Unmarshal(file, &filterwords)
	if err != nil {
		return filterwords.Words, errors.New("解析过滤文件失败")
	}
	return filterwords.Words, nil
}
