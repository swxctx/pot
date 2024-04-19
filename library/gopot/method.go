package gopot

import (
	"fmt"
	"time"
)

// sendAndReceive
func (c *PotClient) sendAndReceive(actionCommand command) (*response, error) {
	// 发送命令
	if err := c.encoder.Encode(actionCommand); err != nil {
		return nil, fmt.Errorf("pot cmd Encode err-> %v, key-> %s", err, actionCommand.Key)
	}

	// 接收响应
	var (
		actionResponse *response
	)
	if err := c.decoder.Decode(&actionResponse); err != nil {
		return nil, fmt.Errorf("pot cmd response Decode err-> %v, key-> %s", err, actionCommand.Key)
	}

	return actionResponse, nil
}

// actionStatusCmd
func (c *PotClient) actionStatusCmd(actionCommand command) *StatusCmd {
	// return data
	statusCmd := newStatusCmd(actionCommand.Key)

	// send cmd
	actionResponse, err := c.sendAndReceive(actionCommand)
	if err != nil {
		statusCmd.setErr(err)
		return statusCmd
	}

	if actionResponse == nil || actionResponse.Code != code_cmd_action_normal {
		// err
		statusCmd = handleErrorResponse(actionResponse, statusCmd)
		statusCmd.setSuccess(false)
		return statusCmd
	}

	statusCmd.setSuccess(true)
	if actionCommand.Cmd == cmd_ttl {
		statusCmd.setResult(actionResponse.TTL)
	}
	if actionCommand.Cmd == cmd_exists {
		statusCmd.setResult(actionResponse.R)
	}
	return statusCmd
}

// actionStringCmd
func (c *PotClient) actionStringCmd(actionCommand command) *StringCmd {
	// return data
	stringCmd := newStringCmd(actionCommand.Key)

	// send cmd
	actionResponse, err := c.sendAndReceive(actionCommand)
	if err != nil {
		stringCmd.setErr(err)
		return stringCmd
	}

	if actionResponse == nil || actionResponse.Code != code_cmd_action_normal {
		stringCmd = handleErrorResponseStringCmd(actionResponse, stringCmd)
		return stringCmd
	}

	stringCmd.setVal(actionResponse.Value)
	return stringCmd
}

// Set
func (c *PotClient) Set(key string, value interface{}, expiration ...time.Duration) *StatusCmd {
	actionCommand := command{
		Cmd:   cmd_set,
		Key:   c.GetKey(key),
		Value: valConvertString(value),
	}
	if len(expiration) > 0 {
		actionCommand.Expire = int64(expiration[0].Seconds())
	}

	return c.actionStatusCmd(actionCommand)
}

// Get
func (c *PotClient) Get(key string) *StringCmd {
	actionCommand := command{
		Cmd: cmd_get,
		Key: c.GetKey(key),
	}

	return c.actionStringCmd(actionCommand)
}

// Del
func (c *PotClient) Del(key string) *StatusCmd {
	actionCommand := command{
		Cmd: cmd_del,
		Key: c.GetKey(key),
	}

	return c.actionStatusCmd(actionCommand)
}

// Exists
func (c *PotClient) Exists(key string) *StatusCmd {
	actionCommand := command{
		Cmd: cmd_exists,
		Key: c.GetKey(key),
	}

	return c.actionStatusCmd(actionCommand)
}

// TTL
func (c *PotClient) TTL(key string) *StatusCmd {
	actionCommand := command{
		Cmd: cmd_ttl,
		Key: c.GetKey(key),
	}

	return c.actionStatusCmd(actionCommand)
}

// Expire
func (c *PotClient) Expire(key string, expiration time.Duration) *StatusCmd {
	actionCommand := command{
		Cmd:    cmd_set,
		Key:    c.GetKey(key),
		Expire: int64(expiration.Seconds()),
	}

	return c.actionStatusCmd(actionCommand)
}
