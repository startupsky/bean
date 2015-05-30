package config

import (
	"flag"
	goyaml "github.com/nporsche/goyaml"
	"io/ioutil"
	"log"
)

type BeanGame struct {
	Listen     int
	Processors int
	Debug      bool
}

var This BeanGame

func Reload(path string) {
	var content []byte
	var err error
	if content, err = ioutil.ReadFile(path); err != nil {
		log.Fatalln("ReloadConfig failure from path:" + path)
		return
	}
	if err := goyaml.Unmarshal(content, &This); err != nil {
		log.Fatalln("ReloadConfig unmarshal failure from path:" + path)
		return
	}
	log.Println("ReloadConfig OK!")
	return
}

func Init() {
	config_path := flag.String("C", "./conf/bean_game.yaml", "The path of bean_game file")
	flag.Parse()
	Reload(*config_path)
}
