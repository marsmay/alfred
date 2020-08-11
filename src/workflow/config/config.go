package config

import (
	"encoding/json"
	"github.com/marsmay/alfred/workflow/output"
	"io/ioutil"
	"os"
	"path"
)

const DataPath = "Library/Application Support/Alfred/Workflow Data"
const FileName = "config.json"

func getConfigPath(appName string) string {
	return path.Join(os.Getenv("HOME"), DataPath, appName)
}

func getConfigFile(appName string) string {
	return path.Join(getConfigPath(appName), FileName)
}

func Load(appName string, config interface{}) (err error) {
	configPath := getConfigPath(appName)
	_, err = os.Stat(configPath)

	if err != nil && os.IsNotExist(err) {
		err = os.MkdirAll(configPath, 0766)
	}

	if err != nil {
		return
	}

	configFile := getConfigFile(appName)
	fd, err := os.OpenFile(configFile, os.O_RDWR|os.O_CREATE, 0766)

	if err != nil {
		return
	}

	defer func() { _ = fd.Close() }()

	bs, err := ioutil.ReadAll(fd)

	if err != nil {
		return
	}

	_ = json.Unmarshal(bs, config)
	return
}

func Save(appName string, config interface{}) (err error) {
	content, err := json.Marshal(config)

	if err != nil {
		return
	}

	configFile := getConfigFile(appName)
	return ioutil.WriteFile(configFile, content, 0666)
}

func SaveResult(appName string, config interface{}) (res *output.Result) {
	if err := Save(appName, config); err != nil {
		return output.NewNotice("Save Settings failed", err)
	}

	return output.NewNotice("Settings saved", nil)
}
