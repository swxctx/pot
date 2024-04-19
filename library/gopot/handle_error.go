package gopot

import (
	"fmt"
)

// handleErrorResponse
func handleErrorResponse(actionResponse *response, statusCmd *StatusCmd) *StatusCmd {
	if actionResponse != nil {
		statusCmd.setErr(fmt.Errorf("pot cmd failed, code-> %d", actionResponse.Code))
	} else {
		statusCmd.setErr(fmt.Errorf("pot server response is nil"))
	}
	return statusCmd
}

// handleErrorResponseExists
func handleErrorResponseExists(actionResponse *response, statusCmd *StatusCmd) *StatusCmd {
	if actionResponse != nil {
		// not exists
		if actionResponse.Code == code_cmd_action_key_not_exists {
			statusCmd.setResult(EXPIRATION_IS_EXPIRED)
			statusCmd.setSuccess(true)
			return statusCmd
		}
		statusCmd.setErr(fmt.Errorf("pot exists failed, code-> %d", actionResponse.Code))
	} else {
		statusCmd.setErr(fmt.Errorf("pot exists server response is nil"))
	}
	return statusCmd
}

// handleErrorResponseStringCmd
func handleErrorResponseStringCmd(actionResponse *response, stringCmd *StringCmd) *StringCmd {
	if actionResponse != nil {
		stringCmd.setErr(fmt.Errorf("pot cmd failed, code-> %d", actionResponse.Code))
	} else {
		stringCmd.setErr(fmt.Errorf("pot server response is nil"))
	}
	return stringCmd
}
