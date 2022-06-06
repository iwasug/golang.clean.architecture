package middleware

type Result struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type DataResult struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func Success(message string) Result {
	return Result{
		Success: true,
		Message: message,
	}
}

func Error(message string) Result {
	return Result{
		Success: false,
		Message: message,
	}
}

func SuccessData(message string, data any) DataResult {
	return DataResult{
		Success: true,
		Message: message,
		Data:    data,
	}
}

func ErrorData(message string, data any) DataResult {
	return DataResult{
		Success: false,
		Message: message,
		Data:    data,
	}
}
