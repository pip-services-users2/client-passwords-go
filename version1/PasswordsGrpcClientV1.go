package version1

import (
	"context"

	"github.com/pip-services-users2/client-passwords-go/protos"
	"github.com/pip-services3-gox/pip-services3-grpc-gox/clients"
)

type PasswordGrpcClientV1 struct {
	*clients.GrpcClient
}

func NewPasswordGrpcClientV1() *PasswordGrpcClientV1 {
	return &PasswordGrpcClientV1{
		GrpcClient: clients.NewGrpcClient("passwords_v1.Passwords"),
	}
}

func (c *PasswordGrpcClientV1) GetPasswordInfo(ctx context.Context, correlationId string,
	userId string) (result *UserPasswordInfoV1, err error) {
	timing := c.Instrument(ctx, correlationId, "passwords_v1.get_password_info")
	defer timing.EndTiming(ctx, err)

	req := &protos.PasswordIdRequest{
		CorrelationId: correlationId,
		UserId:        userId,
	}

	reply := new(protos.PasswordInfoReply)
	err = c.CallWithContext(ctx, "get_password_info", correlationId, req, reply)
	if err != nil {
		return nil, err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return nil, err
	}

	result = toPasswordInfo(reply.Info)

	return result, nil
}

func (c *PasswordGrpcClientV1) SetTempPassword(ctx context.Context, correlationId string,
	userId string) (password string, err error) {
	timing := c.Instrument(ctx, correlationId, "passwords_v1.set_temp_password")
	defer timing.EndTiming(ctx, err)

	req := &protos.PasswordIdRequest{
		CorrelationId: correlationId,
		UserId:        userId,
	}

	reply := new(protos.PasswordValueReply)
	err = c.CallWithContext(ctx, "set_temp_password", correlationId, req, reply)
	if err != nil {
		return "", err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return "", err
	}

	password = reply.Password

	return password, nil
}

func (c *PasswordGrpcClientV1) SetPassword(ctx context.Context, correlationId string,
	userId string, password string) (err error) {
	timing := c.Instrument(ctx, correlationId, "passwords_v1.set_password")
	defer timing.EndTiming(ctx, err)

	req := &protos.PasswordIdAndValueRequest{
		CorrelationId: correlationId,
		UserId:        userId,
		Password:      password,
	}

	reply := new(protos.PasswordEmptyReply)
	err = c.CallWithContext(ctx, "set_password", correlationId, req, reply)
	if err != nil {
		return err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return err
	}

	return nil
}

func (c *PasswordGrpcClientV1) DeletePassword(ctx context.Context, correlationId string,
	userId string) (err error) {
	timing := c.Instrument(ctx, correlationId, "passwords_v1.delete_password")
	defer timing.EndTiming(ctx, err)

	req := &protos.PasswordIdRequest{
		CorrelationId: correlationId,
		UserId:        userId,
	}

	reply := new(protos.PasswordEmptyReply)
	err = c.CallWithContext(ctx, "delete_password", correlationId, req, reply)
	if err != nil {
		return err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return err
	}

	return nil
}

func (c *PasswordGrpcClientV1) Authenticate(ctx context.Context, correlationId string,
	userId string, password string) (authenticated bool, err error) {
	timing := c.Instrument(ctx, correlationId, "passwords_v1.authenticate")
	defer timing.EndTiming(ctx, err)

	req := &protos.PasswordIdAndValueRequest{
		CorrelationId: correlationId,
		UserId:        userId,
		Password:      password,
	}

	reply := new(protos.PasswordAuthenticateReply)
	err = c.CallWithContext(ctx, "authenticate", correlationId, req, reply)
	if err != nil {
		return false, err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return false, err
	}

	authenticated = reply.Authenticated

	return authenticated, nil
}

func (c *PasswordGrpcClientV1) ChangePassword(ctx context.Context, correlationId string,
	userId string, oldPassword string, newPassword string) (err error) {
	timing := c.Instrument(ctx, correlationId, "passwords_v1.change_password")
	defer timing.EndTiming(ctx, err)

	req := &protos.PasswordIdAndValuesRequest{
		CorrelationId: correlationId,
		UserId:        userId,
		OldPassword:   oldPassword,
		NewPassword:   newPassword,
	}

	reply := new(protos.PasswordEmptyReply)
	err = c.CallWithContext(ctx, "change_password", correlationId, req, reply)
	if err != nil {
		return err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return err
	}

	return nil
}

func (c *PasswordGrpcClientV1) ValidateCode(ctx context.Context, correlationId string,
	userId string, code string) (valid bool, err error) {
	timing := c.Instrument(ctx, correlationId, "passwords_v1.validate_code")
	defer timing.EndTiming(ctx, err)

	req := &protos.PasswordIdAndCodeRequest{
		CorrelationId: correlationId,
		UserId:        userId,
		Code:          code,
	}

	reply := new(protos.PasswordValidReply)
	err = c.CallWithContext(ctx, "validate_code", correlationId, req, reply)
	if err != nil {
		return false, err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return false, err
	}

	valid = reply.Valid

	return valid, nil
}

func (c *PasswordGrpcClientV1) ResetPassword(ctx context.Context, correlationId string,
	userId string, code string, password string) (err error) {
	timing := c.Instrument(ctx, correlationId, "passwords_v1.reset_password")
	defer timing.EndTiming(ctx, err)

	req := &protos.PasswordIdAndCodeAndValueRequest{
		CorrelationId: correlationId,
		UserId:        userId,
		Code:          code,
		Password:      password,
	}

	reply := new(protos.PasswordEmptyReply)
	err = c.CallWithContext(ctx, "reset_password", correlationId, req, reply)
	if err != nil {
		return err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return err
	}

	return nil
}

func (c *PasswordGrpcClientV1) RecoverPassword(ctx context.Context, correlationId string,
	userId string) (err error) {
	timing := c.Instrument(ctx, correlationId, "passwords_v1.recover_password")
	defer timing.EndTiming(ctx, err)

	req := &protos.PasswordIdRequest{
		CorrelationId: correlationId,
		UserId:        userId,
	}

	reply := new(protos.PasswordEmptyReply)
	err = c.CallWithContext(ctx, "recover_password", correlationId, req, reply)
	if err != nil {
		return err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return err
	}

	return nil
}
