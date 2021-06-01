package errors

import "errors"

var (
	RecordNotFound = errors.New("record not found")
	DataExisted = errors.New("data existed")
	ParentDoesNotExist = errors.New("parent does not exist")
)
