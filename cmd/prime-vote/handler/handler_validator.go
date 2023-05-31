package handler

import "fmt"

type ErrorMsg struct {
	Field string
	Msg   string
}

func (e ErrorMsg) Error() string {
	return fmt.Sprintf("field: %s msg %s", e.Field, e.Msg)
}

type ApiError struct {
	ErrorMsg []ErrorMsg
}

func (a ApiError) Error() string {
	msgAll := ""
	for _, v := range a.ErrorMsg {
		msgAll += fmt.Sprintf("error: %s", v.Error())
	}

	return msgAll
}

func (a ApiError) ToJSONError() []map[string]interface{} {
	errMsgs := []map[string]interface{}{}
	for _, v := range a.ErrorMsg {
		errMsgs = append(errMsgs, map[string]interface{}{
			"field":   v.Field,
			"message": v.Msg,
		})
	}

	return errMsgs
}

func msgForTag(tag string) string {
	switch tag {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email"
	case "oneof":
		return "Some value is not match"
	case "gt":
		return "Should be not empty"
	}

	return ""
}
