package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPconfig struct {
	Adress string `yaml:"address"`
}

type AppFlags struct {
	ConfigPath string
}

func MustLoad(cfgPath string, cfg any) {
	if cfgPath == "" {
		log.Fatal("Config path is not set")
	}
	if _, err := os.Stat(cfgPath); os.IsNotExist(err) {
		log.Fatalf("Config file does not exist by this path: %s", cfgPath)
	}
	if err := cleanenv.ReadConfig(cfgPath, cfg); err != nil {
		log.Fatalf("Error reading config %s", err)
	}
}

func ParseFlags() AppFlags {
	configPath := flag.String("config", "", "Path to config")
	return AppFlags{
		ConfigPath: *configPath,
	}
}
