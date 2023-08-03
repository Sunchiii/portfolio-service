package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type Config struct{
  PGUrl string `yaml:"pg_url"`
  Port string `yaml:"port"`
}


func NewConfig() (*Config,error){
  // read the env.yaml file
  data ,err := ioutil.ReadFile("env.yaml")
  if err != nil{
    return nil,err
  }

  //Parse the YAML data into  a config  struct
  config := Config{}
  err = yaml.Unmarshal(data,&config)
  if err != nil{
    return nil,err
  }

  return &Config{PGUrl: config.PGUrl},nil
}
