package main

import (
	"encoding/json"
	"fmt"
)

// command pot server command
type command struct {
	// cache cmd
	Cmd string `json:"c,omitempty"`
	// cache key
	Key string `json:"k,omitempty"`
	// cache value
	Value string `json:"v,omitempty"`
	// cache key expire
	Expire int64 `json:"ep,omitempty"`
}

// parseCmd 命令解析
func parseCmd(data []byte) (*command, error) {
	var (
		cmd *command
	)
	if err := json.Unmarshal(data, &cmd); err != nil {
		return nil, fmt.Errorf("parseCmd: Unmarshal err-> %v", err)
	}

	return cmd, nil
}

// response
type response struct {
	// cache err
	Code int `json:"c,omitempty"`
	// cache value
	Value string `json:"v,omitempty"`
	// cache ttl
	TTL int64 `json:"tl,omitempty"`
	// result exists
	R int64 `json:"r,omitempty"`
}

// marshalResponse
func marshalResponse(response *response) (string, error) {
	data, err := json.Marshal(&response)
	if err != nil {
		return "", fmt.Errorf("marshalResponse: err-> %v", err)
	}

	return string(data), err
}

// getCodeResponse
func getCodeResponse(code int) string {
	return fmt.Sprintf("{\"c\":%d}", code)
}
