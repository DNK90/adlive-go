package helpers

import "errors"

const (
	Success                   = 0
	ScreenAlreadyHasOwnerCode = 100001
	InvalidMACAddress         = 100002
	InvalidUserId             = 100003
	InternalError             = 100004
	RecordNotFound            = 100005
	InvalidLocation           = 100006
	InvalidArea               = 100007
	UpdateScreenFailed        = 100008
	NotOwner                  = 100009
	UpdateAreaFailed          = 100010
	InvalidScreenId           = 100011
	ScreenIsNotAvailable      = 100012
	InvalidVideoId            = 100013
)

var (
	InvalidUserIdErr = errors.New(MessageFromCode[InvalidUserId])
	InvalidScreenIdErr = errors.New(MessageFromCode[InvalidScreenId])
	ScreenIsNotAvailableErr = errors.New(MessageFromCode[ScreenIsNotAvailable])
	InvalidVideoIdErr = errors.New(MessageFromCode[InvalidVideoId])
)

var MessageFromCode = map[int]string{
	Success: "success",
	ScreenAlreadyHasOwnerCode: "screen already has owner",
	InvalidMACAddress: "invalid mac address",
	InvalidUserId: "invalid user id",
	InternalError: "internal error",
	RecordNotFound: "record not found",
	InvalidLocation: "invalid location",
	InvalidArea: "invalid area",
	UpdateScreenFailed: "update screen failed",
	NotOwner: "user is not owner",
	UpdateAreaFailed: "update area failed",
	InvalidScreenId: "invalid screen id",
	ScreenIsNotAvailable: "screen is not available",
	InvalidVideoId: "invalid video id",
}
