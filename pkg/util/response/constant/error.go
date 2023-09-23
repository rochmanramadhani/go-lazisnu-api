package constant

import "github.com/jackc/pgx/v5"

const (
	//basic
	E_UNAUTHORIZED        = "E-4001"
	E_DUPLICATE           = "E-4002"
	E_BAD_REQUEST         = "E-4003"
	E_VALIDATION          = "E-4004"
	E_ROUTE_NOTFOUND      = "E-4005"
	E_DATA_NOTFOUND       = "E-4006"
	E_UNPROCESSABL_ENTITY = "E-4007"
	E_INVALID_CREDENTIALS = "E-4008"
	E_RECORD_NOTFOUND     = "E-4009"
	E_EMAIL_USED          = "E-4010"
	E_EMAIL_NOTFOUND      = "E-4011"

	E_SERVER_ERROR       = "E-5001"
	E_TX_CLOSED          = "E-5002"
	E_TX_COMMIT_ROLLBACK = "E-5003"
)

var (
	CODE_TO_MESSAGE = map[string]string{
		E_UNAUTHORIZED:        "unauthorized, please login",
		E_DUPLICATE:           "created value already exists",
		E_BAD_REQUEST:         "bad request, please check payload",
		E_VALIDATION:          "invalid request parameters",
		E_ROUTE_NOTFOUND:      "route not found",
		E_DATA_NOTFOUND:       "data not found",
		E_UNPROCESSABL_ENTITY: "unprocessable entity",
		E_INVALID_CREDENTIALS: "invalid credentials",
		E_RECORD_NOTFOUND:     pgx.ErrNoRows.Error(),
		E_EMAIL_USED:          "email already used",
		E_EMAIL_NOTFOUND:      "email not found",
		E_SERVER_ERROR:        "internal server error",
		E_TX_CLOSED:           pgx.ErrTxClosed.Error(),
		E_TX_COMMIT_ROLLBACK:  pgx.ErrTxCommitRollback.Error(),
	}

	MESSAGE_TO_CODE = func(message string) string {
		for c, v := range CODE_TO_MESSAGE {
			if v == message {
				return c
			}
		}
		return E_SERVER_ERROR
	}
)
