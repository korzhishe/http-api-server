package main

import "github.com/gopherschool/http-rest-api/internal/app/apiserver"

import "log"

import "flag"

import "github.com/BurntSushi/toml"

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-paht", "configs/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	s := apiserver.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
