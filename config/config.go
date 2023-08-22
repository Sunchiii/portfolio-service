package config

import (
	"fmt"
	"os"
	"strconv"

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
  pgUrl := fmt.Sprintf(`postgresql://%s:%s@%s:%s/%s?sslmode=disable`,os.Getenv("PG_USER"),
    os.Getenv("PG_PASSWORD"),
    os.Getenv("PG_HOST"),os.Getenv("PG_PORT"),
    os.Getenv("PG_DATABASE"))


  return &ConfigResult{PGUrl: pgUrl,Port: os.Getenv("PORT")},nil
}

func GetPasetoConfig()(*PasetoConfig,error){
  expHour,err := strconv.Atoi(os.Getenv("TOKEN_EXP"))
  if err != nil {
    return nil,err
  }
  return &PasetoConfig{
    SignatureKey:os.Getenv("SIGNATURE_KEY"),
    ExpHour:expHour,
  },nil
}
