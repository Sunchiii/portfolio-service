package config

import (
	"io/ioutil"
	"gopkg.in/yaml.v3"
)

type Config struct{
  PGUrl string `yaml:"pg_url"`
  Port string `yaml:"port"`
}

type PasetoConfig struct{
  SignatureKey string `yaml:"signature_key"`
  ExpHour int  `yaml:"token_exp"`
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

  return &Config{PGUrl: config.PGUrl,Port: config.Port},nil
}

func GetPasetoConfig()(*PasetoConfig,error){
  data, err := ioutil.ReadFile("env.yaml")
  if err != nil{
    return nil,err
  }

  var pasetoConfig PasetoConfig
  err = yaml.Unmarshal(data,&pasetoConfig)
  if err != nil{
    return nil,err
  }

  return &pasetoConfig,nil
}
