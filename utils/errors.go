package utils

import (
	"encoding/json"
	"fmt"
)

type ErrorType string

const (
	// ErrorTypeNetwork Network error
	ErrorTypeNetwork ErrorType = "network_error"
	// ErrorTypeAPI API error
	ErrorTypeAPI ErrorType = "api_error"
	// ErrorTypeAuth Auth error
	ErrorTypeAuth ErrorType = "auth_error"
	// ErrorTypeParam param error
	ErrorTypeParam ErrorType = "param_error"
	// ErrorTypeInternal Internal error
	ErrorTypeInternal ErrorType = "internal_error"
)

type MoreLoginError struct {
	Type    ErrorType `json:"type"`
	Message string    `json:"message"`
	Code    int       `json:"code,omitempty"`
	Details string    `json:"details,omitempty"`
}

func (e *MoreLoginError) Error() string {
	if e.Details != "" {
		return fmt.Sprintf("[%s] %s (code: %d, details: %s)", e.Type, e.Message, e.Code, e.Details)
	}
	return fmt.Sprintf("[%s] %s (code: %d)", e.Type, e.Message, e.Code)
}

func NewNetworkError(err error) *MoreLoginError {
	return &MoreLoginError{
		Type:    ErrorTypeNetwork,
		Message: "Network communication error",
		Details: err.Error(),
	}
}

func NewAPIError(statusCode int, body []byte) *MoreLoginError {
	var details string
	var apiErr struct {
		Message string `json:"message"`
		Error   string `json:"error"`
	}

	if err := json.Unmarshal(body, &apiErr); err == nil {
		if apiErr.Message != "" {
			details = apiErr.Message
		} else if apiErr.Error != "" {
			details = apiErr.Error
		}
	}

	if details == "" {
		details = string(body)
	}

	return &MoreLoginError{
		Type:    ErrorTypeAPI,
		Message: fmt.Sprintf("API returned error status: %d", statusCode),
		Code:    statusCode,
		Details: details,
	}
}

func NewAuthError() *MoreLoginError {
	return &MoreLoginError{
		Type:    ErrorTypeAuth,
		Message: "Authentication failed, please check your access token",
	}
}

func NewParamError(param string, details string) *MoreLoginError {
	return &MoreLoginError{
		Type:    ErrorTypeParam,
		Message: fmt.Sprintf("Invalid parameter: %s", param),
		Details: details,
	}
}

func NewInternalError(err error) *MoreLoginError {
	return &MoreLoginError{
		Type:    ErrorTypeInternal,
		Message: "Internal server error",
		Details: err.Error(),
	}
}

func IsAuthError(err error) bool {
	if moreLoginErr, ok := err.(*MoreLoginError); ok {
		return moreLoginErr.Type == ErrorTypeAuth
	}
	return false
}

func IsAPIError(err error) bool {
	if moreLoginErr, ok := err.(*MoreLoginError); ok {
		return moreLoginErr.Type == ErrorTypeAPI
	}
	return false
}

func IsNetworkError(err error) bool {
	if moreLoginErr, ok := err.(*MoreLoginError); ok {
		return moreLoginErr.Type == ErrorTypeNetwork
	}
	return false
}

func IsParamError(err error) bool {
	if moreLoginErr, ok := err.(*MoreLoginError); ok {
		return moreLoginErr.Type == ErrorTypeParam
	}
	return false
}
