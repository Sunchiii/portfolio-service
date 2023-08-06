package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type Config struct{
  Port string `yaml:"port"`

  PGUser string `yaml:"pg_user"`
  PGPassword string `yaml:"pg_password"`
  PGPort string `yaml:"pg_port"`
  PGDatabase string `yaml:"pg_database"`
  PGHost string `yaml:"pg_host"`
}

type ConfigResult struct{
  PGUrl string
  Port string
}

type PasetoConfig struct{
  SignatureKey string `yaml:"signature_key"`
  ExpHour int  `yaml:"token_exp"`
}


func NewConfig() (*ConfigResult,error){
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

  pgUrl := fmt.Sprintf(`postgresql://%s:%s@%s:%s/%s?sslmode=disable`,config.PGUser,config.PGPassword,config.PGHost,config.PGPort,config.PGDatabase)

  return &ConfigResult{PGUrl: pgUrl,Port: config.Port},nil
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
