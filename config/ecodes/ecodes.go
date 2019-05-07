package ecodes

import (
	"base/pkg/cerror"
	"google.golang.org/grpc/codes"
)

var Nil = cerror.Nil

var (
	ErrNotFound = cerror.Error{
		Type:         codes.NotFound,
		InternalCode: "DOC_NOT_FOUND",
		Text:         "doc not found",
		IsSensible:   false,
		Severity:     cerror.ERROR,
		NeedReport:   false,
	}

	ErrUserEmailNotFound = cerror.Error{
		Type:         codes.NotFound,
		InternalCode: "USER_EMAIL_NOT_FOUND",
		Text:         "User email not found",
		Description:  "The user is searched in the db by the email and not founded",
		IsSensible:   false,
		Severity:     cerror.ERROR,
		NeedReport:   false,
	}

	ErrDuplicatedData = cerror.Error{
		Type:         codes.AlreadyExists,
		InternalCode: "DUPLICATED_DATA",
		Text:         "field value is already in the database",
		IsSensible:   false,
		Severity:     cerror.ERROR,
		NeedReport:   false,
	}

	ErrInvalidPassword = cerror.Error{
		Type:         codes.PermissionDenied,
		InternalCode: "INV_PASSWORD",
		Text:         "Invalid password",
		IsSensible:   true,
		Severity:     cerror.ERROR,
		NeedReport:   false,
	}

	ErrTokenExpired = cerror.Error{
		Type:         codes.Unauthenticated,
		InternalCode: "ERR_TOKEN_EXPIRED",
		Text:         "Token expired",
		IsSensible:   false,
		Severity:     cerror.EXPECTED_ERROR,
		NeedReport:   false,
	}

	ErrTokenMalformed = cerror.Error{
		Type:         codes.Unauthenticated,
		InternalCode: "ERR_TOKEN_MALFORMED",
		Text:         "Token malformed",
		IsSensible:   false,
		Severity:     cerror.SUSPECT,
		NeedReport:   true,
	}

	ErrInvTokenSignature = cerror.Error{
		Type:         codes.Unauthenticated,
		InternalCode: "ERR_TOKEN_BAD_SIG",
		Text:         "The signature of the token is invalid",
		IsSensible:   false,
		Severity:     cerror.SUSPECT,
		NeedReport:   true,
	}
)
