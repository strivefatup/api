package config

import (
	"os"
	"log"
	"time"
	"strings"
	"path/filepath"
	"github.com/spf13/viper"
)

const (
	LevelInfo    = "[INFO]"
	LevelDebug   = "[DEBUG]"
	LevelError   = "[ERROR]"
	FileType     = "yaml"
)

type Logger struct {

}


func IsExist(path string) bool {
    _, err := os.Stat(path)
    return err == nil || os.IsExist(err)
}

func Config(key string) interface{} {
	//判断是否是正确的获取配置文件的规范
	if (strings.Index(key, ".") == -1) ||  (key == "") {
		return nil
	}

	//将字符串解析成文件名和key
	resolve := strings.SplitN(key, ".", 2)

	//获取app的绝对路径
	appPath, _ := filepath.Abs(filepath.Dir(""))
	//读取yaml文件
	v := viper.New()
	//设置读取的配置文件
	v.SetConfigName(resolve[0])
	//添加读取的配置文件路径
	v.AddConfigPath(appPath + "/config/")
	//windows环境下为%GOPATH，linux环境下为$GOPATH
	v.AddConfigPath("$GOPATH/src/")
	//设置配置文件类型
	v.SetConfigType(FileType)
 
	if err := v.ReadInConfig();err != nil {
		return nil
	}
	return v.Get(resolve[1])
}

func GetLogFile() string {
	appPath, _ := filepath.Abs(filepath.Dir(""))
	logPath := appPath + "/resources/log/" + time.Now().Format("2006-01-02") + ".log"
	logIsExist := IsExist(logPath)
	if !logIsExist {
		file, _ := os.Create(logPath)
		defer file.Close()
	}
	return logPath
}

func (l *Logger) Info(info ...interface{}) {
	file, _ := os.OpenFile(GetLogFile(), os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	log.New(file, LevelInfo, log.LstdFlags)
	log.SetOutput(file)
	log.Println(info)
}

func (l *Logger) Debug(debug ...interface{}) {
	file, _ := os.OpenFile(GetLogFile(), os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	log.New(file, LevelDebug, log.LstdFlags)
	log.SetOutput(file)
	log.Panic(debug)
}

func (l *Logger) Error(err ...interface{}) {
	file, _ := os.OpenFile(GetLogFile(), os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	log.New(file, LevelError, log.LstdFlags)
	log.SetOutput(file)
	log.Fatal(err)
}



