package constant

// db
const DB_DEFAULT_SYSTEM = "system"

// code
const (
	CODE_LENGTH = 5

	CODE_ROLE_PREFIX = "R-"
)

// trx
const (
	TRX_TYPE_IN      = "in"
	TRX_TYPE_OUT     = "out"
	TRX_TYPE_CONVERT = "convert"
)

// status
const (
	STATUS_SUCCESS = "SUCCESS"
	STATUS_FAILED  = "FAILED"
)

// convert
const (
	CONVERT_TYPE_ORIGIN      = "origin"
	CONVERT_TYPE_DESTINATION = "destination"
)

// action
const (
	ACTION_INSERT = "insert"
	ACTION_UPDATE = "update"
	ACTION_DELETE = "delete"
)

// firebase
const (
	FIRESTORE_MAX_DATA                   = 50
	FIRESTORE_COLLECTION_DASHBOARD_ORDER = "dashboard-order"
	FIRESTORE_COLLECTION_TOTAL_ORDER     = "total-order"
)
