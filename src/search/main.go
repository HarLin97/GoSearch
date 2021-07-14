package main

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"search/src/config"
	"search/src/util"
	"time"
)

var f = flag.String("f", "config.yml", "配置文件路径")

func main() {
	fmt.Println("全库检索开始")
	data, _ := ioutil.ReadFile(*f)
	c := config.Config{}
	err := yaml.Unmarshal(data, &c)
	if err != nil {
		return
	}
	flag.Parse()
	// 遍历所有目录
	for _, dir := range c.RootDirs {
		// 遍历所有类型
		for _, s := range c.Suffix {
			files, _ := util.WalkDir(dir, s)
			if c.Search != "all" {
				files, _ = util.ListDir(dir, s)
			}
			// 遍历所有文件
			for _, file := range files {
				// 遍历所有正则规则
				for _, regular := range c.Regulars {
					// 匹配正则表达式
					err := util.ReadLine(file, regular)
					if err != nil {
						return
					}
				}
			}
		}
	}
	fmt.Println("日志文件保存在：" + "./log_" + time.Now().Format("2006-01-02") + ".log")
}
