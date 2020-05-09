package config

import (
	"errors"
	"github.com/go-yaml/yaml"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func Load(dir string) {
	path, err := realPath(dir)

	if err != nil {
		log.Fatal(err)
	}

	bytes, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(bytes, &Conf)

	if err != nil {
		log.Fatal(err)
	}
}

func realPath(path string) (string, error) {
	state, err := os.Stat(path)

	if err != nil {
		return "", err
	}

	if state.IsDir() {
		return "", errors.New("config file can not be directory")
	}

	abs, err := filepath.Abs(path)

	if err != nil {
		return "", err
	}

	return abs, err
}
