package util

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"search/src/config"
	"strconv"
	"time"
)

func ReadLine(filePth string, regStr config.Regular) error {

	file := "./log_" + time.Now().Format("2006-01-02") + ".log"
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if nil != err {
		panic(err)
	}
	debugLog := log.New(logFile, "[Regular："+regStr.Description+"]", log.Ltime)
	f, err := os.Open(filePth)
	if err != nil {
		return err
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)

	bfRd := bufio.NewReader(f)
	for i := 0; i <= bfRd.Size(); i++ {
		line, err := bfRd.ReadBytes('\n')
		//解析正则表达式，如果成功返回解释器
		reg1 := regexp.MustCompile(regStr.Expression)
		if reg1 == nil { //解释失败，返回nil
			fmt.Println("regexp err")
		}
		//根据规则提取关键信息
		result1 := reg1.FindAllStringSubmatch(string(line), -1)
		if result1 != nil {
			debugLog.Println(" filePath:"+filePth+" num: "+strconv.Itoa(i+1)+" matches:", result1)
		}
		if err != nil { //遇到任何错误立即返回，并忽略 EOF 错误信息
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
	return nil
}
