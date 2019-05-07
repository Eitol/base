package cerror

// Hector Oliveros - 2019
// hector.oliveros.leon@gmail.com

import (
	"base/api/gen/go/api/protos"
	"github.com/golang/protobuf/proto"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strconv"
	"strings"
)

const PrimitiveErrorTag = "PRIMITIVE_ERROR"
const maxDep = 1000 // avoid infinite loops

var Nil = Error{InternalCode: ""}

const AppendDebugCodes = false

type Error struct {
	Type         codes.Code        `json:"type"`
	InternalCode string            `json:"code"`
	Text         string            `json:"text"`
	Description  string            `json:"description"`
	Reason       string            `json:"reason"`
	IsSensible   bool              `json:"isSensible" bson:"isSensible"`
	Severity     Severity          `json:"severity" bson:"severity"`
	NeedReport   bool              `json:"needReport" bson:"needReport"`
	Meta         map[string]string `json:"meta,omitempty"`
	ComesFrom    *Error            `json:"comesFrom,omitempty"`
}

type Error2 struct {
	genpb.Error
}

func (err Error) GetParents() []*Error {
	list := make([]*Error, 0)
	cont := 0
	currentErr := &err
	for {
		cont++
		if currentErr.ComesFrom == nil || cont > maxDep {
			break
		}
		list = append(list, currentErr.ComesFrom)
		currentErr = currentErr.ComesFrom
	}
	return list
}

func (err Error) IsError() bool {
	return err.InternalCode != Nil.InternalCode || err.ComesFrom != nil
}

func (err Error) Equals(err2 Error) bool {
	return err.InternalCode == err2.InternalCode
}

func (err Error) SetParam(num int, val string) Error {
	errR := err
	param := "{" + strconv.Itoa(num) + "}"
	errR.Reason = strings.Replace(errR.Reason, param, val, 1)
	return errR
}

func From(err2 error) Error {
	err := Error{}
	return err.From(err2)
}

func (err Error) From(err2 error) Error {
	errNew := err
	var ptr Error
	switch err2.(type) {
	case Error:
		ptr = err2.(Error)
	default:
		ptr = Error{
			InternalCode: PrimitiveErrorTag,
			Text:         err2.Error(),
		}
	}
	errNew.ComesFrom = &ptr
	return errNew
}

func (err Error) SetReason(reason string) Error {
	errNew := err
	errNew.Reason = reason
	return errNew
}

func (err Error) AddMeta(key string, val string) Error {
	errNew := err
	if errNew.Meta == nil {
		errNew.Meta = make(map[string]string)
	}
	errNew.Meta[key] = val
	return errNew
}

func (err Error) Error() string {
	st := status.New(err.Type, err.Text)
	parentErrors := err.GetParents()

	errDetList := make([]proto.Message, 1)
	errCodeList := make([]string, 1)
	if len(parentErrors) > 0 {
		for _, err := range parentErrors {
			if err == nil {
				continue
			}
			errCodeList = append(errCodeList, err.InternalCode)
		}
	}

	if AppendDebugCodes {
		debugInfo := &errdetails.DebugInfo{
			StackEntries: errCodeList,
			Detail:       "",
		}
		errDetList = append(errDetList, debugInfo)
	}

	switch err.Type {
	case codes.Unknown:
		return status.New(codes.NotFound, err.Text).Err().Error()

	case codes.InvalidArgument:
		for fieldName, desc := range err.Meta {
			v := &errdetails.BadRequest_FieldViolation{
				Field:       fieldName,
				Description: desc,
			}
			br := &errdetails.BadRequest{}
			br.FieldViolations = append(br.FieldViolations, v)
		}
		sb, _ := st.WithDetails(errDetList...)
		return sb.Err().Error()
	}
	return status.New(err.Type, err.Text+" | "+err.InternalCode).Err().Error()
}
