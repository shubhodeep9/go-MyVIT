package status

type StatusStruct struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func SessionError() StatusStruct {
	return StatusStruct{
		Message: "Session Timed Out",
		Code:    11,
	}
}

func Success() StatusStruct {
	return StatusStruct{
		Message: "Success",
		Code:    0,
	}
}

func CredentialError() StatusStruct {
	return StatusStruct{
		Message: "Invalid Credentials",
		Code:    12,
	}
}
