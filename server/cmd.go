package main

import (
	"bufio"
	"fmt"
	"net"
	"time"

	"github.com/swxctx/pot"

	"github.com/swxctx/pot/plog"
)

// cmdHandleConnection logic
func cmdHandleConnection(conn net.Conn) {
	defer conn.Close()

	remoteAddr := conn.RemoteAddr().String()
	plog.Infof("Handling new connection, addr-> %s", remoteAddr)

	// read
	scanner := bufio.NewScanner(conn)

	// read for
	for scanner.Scan() {
		// read cmd
		cmdData := scanner.Bytes()
		plog.Tracef("%s: %s", remoteAddr, string(cmdData))

		// run cmd
		cmdRunLogic(conn, cmdData)
	}
}

// cmdRunLogic
func cmdRunLogic(conn net.Conn, cmdData []byte) {
	// parse data
	cmd, err := parseCmd(cmdData)
	if err != nil {
		responseClient(conn, getCodeResponse(code_err_for_parse_cmd))
		return
	}

	if cmd == nil {
		responseInterval(conn)
		return
	}

	// cache action logic
	switch cmd.Cmd {
	case cmd_set:
		setCmd := potClient.Set(cmd.Key, cmd.Value, time.Duration(cmd.Expire)*time.Second)
		if setCmd.Err() != nil {
			plog.Errorf("cmdRunLogic: cache set, key-> %s, err-> %v", cmd.Key, err)
			responseInterval(conn)
			return
		}
		if !setCmd.Success() {
			plog.Errorf("cmdRunLogic: cache set, failed, key-> %s", cmd.Key)
			responseActionCmdFailed(conn)
			return
		}
		plog.Tracef("cmdRunLogic: cache set, success, key-> %s, expire-> %d", cmd.Key, cmd.Expire)
		responseSuccessForOnlyAction(conn)
		break
	case cmd_get:
		// cache get
		getCmd := potClient.Get(cmd.Key)
		if getCmd.Err() != nil {
			plog.Errorf("cmdRunLogic: cache get, key-> %s, err-> %v", cmd.Key, err)
			responseInterval(conn)
			return
		}
		// response marshal
		val, err := marshalResponse(&response{
			Code:  code_cmd_action_normal,
			Value: getCmd.String(),
		})
		if err != nil {
			plog.Errorf("cmdRunLogic: cache get, key-> %s, err-> %v", cmd.Key, err)
			responseInterval(conn)
		}
		plog.Tracef("cmdRunLogic: cache get, key-> %s, val-> %s", cmd.Key, val)
		responseClient(conn, val)
		break
	case cmd_del:
		delCmd := potClient.Del(cmd.Key)
		if delCmd.Err() != nil {
			plog.Errorf("cmdRunLogic: cache del, key-> %s, err-> %v", cmd.Key, err)
			responseInterval(conn)
			return
		}
		if !delCmd.Success() {
			plog.Errorf("cmdRunLogic: cache del, failed, key-> %s", cmd.Key)
			responseActionCmdFailed(conn)
			return
		}
		plog.Tracef("cmdRunLogic: cache del, success, key-> %s", cmd.Key)
		responseSuccessForOnlyAction(conn)
		break
	case cmd_exists:
		existsCmd := potClient.Exists(cmd.Key)
		if existsCmd.Err() != nil {
			plog.Errorf("cmdRunLogic: cache exists, key-> %s, err-> %v", cmd.Key, err)
			responseInterval(conn)
			return
		}
		if !existsCmd.Success() {
			plog.Errorf("cmdRunLogic: cache exists, failed, key-> %s", cmd.Key)
			responseActionCmdFailed(conn)
			return
		}

		code := code_cmd_action_normal
		if existsCmd.Result() != pot.POT_ACTION_RESULT_EXISTS {
			code = code_cmd_action_key_not_exists
		}
		plog.Tracef("cmdRunLogic: cache exists, success, key-> %s, exists-> %d", cmd.Key, code)
		responseOnlyCode(conn, code)
		break
	case cmd_ttl:
		// cache ttl
		ttlCmd := potClient.TTL(cmd.Key)
		if ttlCmd.Err() != nil {
			plog.Errorf("cmdRunLogic: cache ttl, key-> %s, err-> %v", cmd.Key, err)
			responseInterval(conn)
			return
		}
		// response marshal
		val, err := marshalResponse(&response{
			Code: code_cmd_action_normal,
			TTL:  ttlCmd.Result(),
		})
		if err != nil {
			plog.Errorf("cmdRunLogic: cache ttl, key-> %s, err-> %v", cmd.Key, err)
			responseInterval(conn)
		}
		plog.Tracef("cmdRunLogic: cache ttl, key-> %s, val-> %s", cmd.Key, val)
		responseClient(conn, val)
		break
	case cmd_expire:
		expireCmd := potClient.Expire(cmd.Key, time.Duration(cmd.Expire)*time.Second)
		if expireCmd.Err() != nil {
			plog.Errorf("cmdRunLogic: cache expire, key-> %s, err-> %v", cmd.Key, err)
			responseInterval(conn)
			return
		}
		if !expireCmd.Success() {
			plog.Errorf("cmdRunLogic: cache expire, failed, key-> %s", cmd.Key)
			responseActionCmdFailed(conn)
			return
		}
		plog.Tracef("cmdRunLogic: cache expire, success, key-> %s, expire-> %d", cmd.Key, cmd.Expire)
		responseSuccessForOnlyAction(conn)
		break
	default:
		responseClient(conn, getCodeResponse(code_err_for_cmd_valid))
	}
}

// responseClient
func responseClient(conn net.Conn, data string) {
	fmt.Fprintf(conn, data)
}

// responseInterval
func responseInterval(conn net.Conn) {
	responseClient(conn, getCodeResponse(code_err_for_parse_internal))
}

// responseInterval
func responseActionCmdFailed(conn net.Conn) {
	responseClient(conn, getCodeResponse(code_err_for_cmd_action_failed))
}

// responseSuccessForOnlyAction
func responseSuccessForOnlyAction(conn net.Conn) {
	responseClient(conn, getCodeResponse(code_cmd_action_normal))
}

// responseOnlyCode
func responseOnlyCode(conn net.Conn, code int) {
	responseClient(conn, getCodeResponse(code))
}
