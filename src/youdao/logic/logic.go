package logic

import (
	"alfred/workflow/config"
	"alfred/workflow/output"
	"os"
)

type Config struct {
	AppKey string `json:"appKey"`
	Secret string `json:"secret"`
}

func Excute() (result *output.Result) {
	conf := &Config{}

	if err := config.Load(AppName, conf); err != nil {
		return output.NewNotice("Load settings failed", err)
	}

	if len(os.Args) <= 1 || os.Args[1] == "" {
		return output.NewNotice("Invalid parameters", nil)
	}

	keys := SetKeyRegex.FindStringSubmatch(os.Args[1])

	if len(keys) == 3 {
		conf.AppKey, conf.Secret = keys[1], keys[2]
		return config.SaveResult(AppName, conf)
	} else {
		if conf.AppKey == "" || conf.Secret == "" {
			return output.NewNotice("Invalid appkey and secret", nil)
		}

		items, err := queryWord(conf.AppKey, conf.Secret, os.Args[1])

		if err != nil {
			return output.NewNotice("Query data failed", err)
		}

		if len(items) == 0 {
			return output.NewNotice("No results", nil)
		}

		return &output.Result{Items: items}
	}
}
