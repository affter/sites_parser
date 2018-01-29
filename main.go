package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/golang-collections/collections/stack"
	"github.com/jmoiron/sqlx"
)

var (
	configFile = flag.String("config", "conf.json", "Path to config file")
)
var Db struct {
	Connection *sqlx.DB
}
var Config struct {
	MysqlLogin    string `json:"mysql_login"`
	MysqlPassword string `json:"mysql_password"`
	MysqlHost     string `json:"mysql_host"`
	MysqlDb       string `json:"mysql_db"`
	HttpdDir      string `json:"httpd_dir"`
	PatternPath   string `json:"pattern_path"`
}

func LoadConfig() error {
	jsonData, err := ioutil.ReadFile(*configFile)
	if err != nil {
		return err
	}
	return json.Unmarshal(jsonData, &Config)
}

func main() {
	log.Println("Starting application...")
	flag.Parse()
	if err := LoadConfig(); err != nil {
		log.Fatal(err)
	}
	bytes, err := ioutil.ReadFile(Config.PatternPath);
	if err != nil {
		panic(err)
	}
	parsedPattern := strings.Split(string(bytes), "\n")
	files, err := ioutil.ReadDir(Config.HttpdDir)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, file := range files {
		bytes, err := ioutil.ReadFile(Config.HttpdDir + file.Name());
		if err != nil {
			panic(err)
		}
		parsedFile := strings.Split(string(bytes), "\n")
		i := 0
		tags := stack.New()
		for _, line := range parsedFile {
			if strings.Index(line, "</") != -1 {
				tags.Pop()
			} else if
			strings.Index(line, "<") != -1 {
				tags.Push(line)
			}
			if strings.Index(line, parsedPattern[i]) == 0 {
				i++
			}
		}
	}
}
