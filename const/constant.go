package constant

import (
	"errors"
)

const (
	ErrInvalidRequestBody    = "Invalid request body"
	ErrFailedToParseRequest  = "Failed to parse request body"
	ErrFailedToCreateSubject = "Failed to create subject"
	ErrInvalidLimitParam     = "Invalid limit parameter"
	ErrInvalidPageParam      = "Invalid page parameter"
	ErrFailedGetSubjectList  = "Failed to get subject list in server"
	ErrFailedGetSubjectById  = "Failed to get subject by ID"
)

// Errors
var (
	ErrInvalidRequestBodyErr    = errors.New(ErrInvalidRequestBody)
	ErrFailedToParseRequestErr  = errors.New(ErrFailedToParseRequest)
	ErrFailedToCreateSubjectErr = errors.New(ErrFailedToCreateSubject)
	ErrInvalidLimitParamErr     = errors.New(ErrInvalidLimitParam)
	ErrInvalidPageParamErr      = errors.New(ErrInvalidPageParam)
	ErrFailedGetSubjectListErr  = errors.New(ErrFailedGetSubjectList)
	ErrFailedGetSubjectByIdErr  = errors.New(ErrFailedGetSubjectById)
)
