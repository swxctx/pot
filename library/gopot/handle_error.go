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

// handleErrorResponseStringCmd
func handleErrorResponseStringCmd(actionResponse *response, stringCmd *StringCmd) *StringCmd {
	if actionResponse != nil {
		stringCmd.setErr(fmt.Errorf("pot cmd failed, code-> %d", actionResponse.Code))
	} else {
		stringCmd.setErr(fmt.Errorf("pot server response is nil"))
	}
	return stringCmd
}
