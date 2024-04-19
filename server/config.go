package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var (
	cfg *config
)

// config
type config struct {
	// 监听端口
	Port string `json:"port"`
	// 无效key清理间隔，单位为毫秒
	CleanerInterval int64 `json:"cleaner_interval"`
	// 显示详细运行日志
	ShowTrace bool `json:"show_trace"`
}

// reloadConfig
func reloadConfig() error {
	confPath := "./config/config.json"

	// config file not exists
	if _, err := os.Stat(confPath); os.IsNotExist(err) {
		if err := os.MkdirAll("./config", 0755); err != nil {
			return fmt.Errorf("reloadConfig: mkdir config err-> %v", err)
		}

		// pot default config
		cfg = &config{
			Port:            "9577",
			CleanerInterval: 1000,
			ShowTrace:       true,
		}

		data, err := json.MarshalIndent(cfg, "", "    ")
		if err != nil {
			return fmt.Errorf("reloadConfig: MarshalIndent err-> %v", err)
		}

		// write config.json
		if err := ioutil.WriteFile(confPath, data, 0644); err != nil {
			return fmt.Errorf("reloadConfig: first write config file err-> %v", err)
		}

		return nil
	} else if err != nil {
		return err
	}

	// read config file
	bs, err := ioutil.ReadFile(confPath)
	if err != nil {
		return fmt.Errorf("reloadConfig: read config file err-> %v", err)
	}

	// decode config
	decode := json.NewDecoder(bytes.NewReader(bs))
	decode.UseNumber()
	return decode.Decode(&cfg)
}
