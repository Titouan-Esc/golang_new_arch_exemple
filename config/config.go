package config

import (
	"exemple.com/swagTest/domain/model"
	"github.com/spf13/viper"
	"log"
)

func LoadConfig() (model.Env, error) {
	var env model.Env

	vp := viper.New()

	vp.SetConfigName("config")
	vp.SetConfigType("json")
	vp.AddConfigPath("./config")
	vp.AddConfigPath(".")

	if err := vp.ReadInConfig(); err != nil {
		return env, err
	}

	if err := vp.Unmarshal(&env); err != nil {
		return env, err
	}

	log.Println(env)

	return env, nil
}
