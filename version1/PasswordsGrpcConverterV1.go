package version1

import (
	"encoding/json"

	"github.com/pip-services-users2/client-passwords-go/protos"
	"github.com/pip-services3-gox/pip-services3-commons-gox/convert"
	"github.com/pip-services3-gox/pip-services3-commons-gox/errors"
)

func fromError(err error) *protos.ErrorDescription {
	if err == nil {
		return nil
	}

	desc := errors.ErrorDescriptionFactory.Create(err)
	obj := &protos.ErrorDescription{
		Type:          desc.Type,
		Category:      desc.Category,
		Code:          desc.Code,
		CorrelationId: desc.CorrelationId,
		Status:        convert.StringConverter.ToString(desc.Status),
		Message:       desc.Message,
		Cause:         desc.Cause,
		StackTrace:    desc.StackTrace,
		Details:       fromMap(desc.Details),
	}

	return obj
}

func toError(obj *protos.ErrorDescription) error {
	if obj == nil || (obj.Category == "" && obj.Message == "") {
		return nil
	}

	description := &errors.ErrorDescription{
		Type:          obj.Type,
		Category:      obj.Category,
		Code:          obj.Code,
		CorrelationId: obj.CorrelationId,
		Status:        convert.IntegerConverter.ToInteger(obj.Status),
		Message:       obj.Message,
		Cause:         obj.Cause,
		StackTrace:    obj.StackTrace,
		Details:       toMap(obj.Details),
	}

	return errors.ApplicationErrorFactory.Create(description)
}

func fromMap(val map[string]interface{}) map[string]string {
	r := map[string]string{}

	for k, v := range val {
		r[k] = convert.StringConverter.ToString(v)
	}

	return r
}

func toMap(val map[string]string) map[string]interface{} {
	r := make(map[string]interface{}, 0)

	for k, v := range val {
		r[k] = v
	}

	return r
}

func toJson(value interface{}) string {
	if value == nil {
		return ""
	}

	b, err := json.Marshal(value)
	if err != nil {
		return ""
	}
	return string(b[:])
}

func fromJson(value string) interface{} {
	if value == "" {
		return nil
	}

	var m interface{}
	json.Unmarshal([]byte(value), &m)
	return m
}

func fromPasswordInfo(info *UserPasswordInfoV1) *protos.PasswordInfo {
	if info == nil {
		return nil
	}

	obj := &protos.PasswordInfo{
		Id:         info.Id,
		ChangeTime: convert.StringConverter.ToString(info.ChangeTime),
		Locked:     info.Locked,
		LockTime:   convert.StringConverter.ToString(info.LockTime),
	}

	return obj
}

func toPasswordInfo(obj *protos.PasswordInfo) *UserPasswordInfoV1 {
	if obj == nil {
		return nil
	}

	info := &UserPasswordInfoV1{
		Id:         obj.Id,
		ChangeTime: convert.DateTimeConverter.ToDateTime(obj.ChangeTime),
		Locked:     obj.Locked,
		LockTime:   convert.DateTimeConverter.ToDateTime(obj.LockTime),
	}

	return info
}
