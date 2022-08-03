package response

const (
	STATUS_GENERAL_SUCCESS           = "0000"
	STATUS_CREATE_SUCCESS            = "0001"
	STATUS_READ_SUCCESS              = "0002"
	STATUS_UPDATE_SUCCESS            = "0003"
	STATUS_DELETE_SUCCESS            = "0004"
	STATUS_GENERAL_ERROR             = "5000"
	STATUS_CREATE_ERROR              = "5001"
	STATUS_READ_ERROR                = "5002"
	STATUS_UPDATE_ERROR              = "5003"
	STATUS_DELETE_ERROR              = "5004"
	STATUS_REQUEST_VALIDATION_FAILED = "5005"
	STATUS_DATA_NOT_FOUND            = "5007"
	STATUS_DATA_IS_EMPTY             = "5008"
	STATUS_SQL_ERROR                 = "5050"
)

const (
	MESSAGE_SUCCESS = "Success"
	MESSAGE_ERROR   = "Something went wrong"
)
