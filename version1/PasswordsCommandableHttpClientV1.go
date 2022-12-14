package version1

import (
	"context"

	cdata "github.com/pip-services3-gox/pip-services3-commons-gox/data"
	cclients "github.com/pip-services3-gox/pip-services3-rpc-gox/clients"
)

type PasswordsCommandableHttpClientV1 struct {
	*cclients.CommandableHttpClient
}

func NewPasswordsCommandableHttpClientV1() *PasswordsCommandableHttpClientV1 {
	c := &PasswordsCommandableHttpClientV1{
		CommandableHttpClient: cclients.NewCommandableHttpClient("v1/passwords"),
	}
	return c
}

func (c *PasswordsCommandableHttpClientV1) GetPasswordInfo(ctx context.Context, correlationId string,
	userId string) (result *UserPasswordInfoV1, err error) {
	params := cdata.NewAnyValueMapFromTuples(
		"user_id", userId,
	)

	res, err := c.CallCommand(ctx, "get_password_info", correlationId, params)
	if err != nil {
		return nil, err
	}

	return cclients.HandleHttpResponse[*UserPasswordInfoV1](res, correlationId)
}

func (c *PasswordsCommandableHttpClientV1) SetTempPassword(ctx context.Context, correlationId string,
	userId string) (password string, err error) {
	params := cdata.NewAnyValueMapFromTuples(
		"user_id", userId,
	)

	res, err := c.CallCommand(ctx, "set_temp_password", correlationId, params)
	if err != nil {
		return "", err
	}

	return cclients.HandleHttpResponse[string](res, correlationId)
}

func (c *PasswordsCommandableHttpClientV1) SetPassword(ctx context.Context, correlationId string, userId string, password string) error {
	params := cdata.NewAnyValueMapFromTuples(
		"user_id", userId,
		"password", password,
	)

	_, err := c.CallCommand(ctx, "set_password", correlationId, params)
	return err
}

func (c *PasswordsCommandableHttpClientV1) DeletePassword(ctx context.Context, correlationId string, userId string) error {
	params := cdata.NewAnyValueMapFromTuples(
		"user_id", userId,
	)

	_, err := c.CallCommand(ctx, "delete_password", correlationId, params)
	return err
}

func (c *PasswordsCommandableHttpClientV1) Authenticate(ctx context.Context, correlationId string, userId string,
	password string) (authenticated bool, err error) {
	params := cdata.NewAnyValueMapFromTuples(
		"user_id", userId,
		"password", password,
	)

	res, err := c.CallCommand(ctx, "authenticate", correlationId, params)
	if err != nil {
		return false, err
	}

	result, err := cclients.HandleHttpResponse[map[string]bool](res, correlationId)

	if err != nil {
		return false, err
	}

	val, valOk := result["authenticated"]
	if !valOk {
		return false, nil
	}
	return val, nil
}

func (c *PasswordsCommandableHttpClientV1) ChangePassword(ctx context.Context, correlationId string, userId string,
	oldPassword string, newPassword string) error {
	params := cdata.NewAnyValueMapFromTuples(
		"user_id", userId,
		"old_password", oldPassword,
		"new_password", newPassword,
	)

	_, err := c.CallCommand(ctx, "change_password", correlationId, params)
	return err
}

func (c *PasswordsCommandableHttpClientV1) ValidateCode(ctx context.Context, correlationId string, userId string,
	code string) (valid bool, err error) {
	params := cdata.NewAnyValueMapFromTuples(
		"user_id", userId,
		"code", code,
	)

	res, err := c.CallCommand(ctx, "validate_code", correlationId, params)
	if err != nil {
		return false, err
	}

	result, err := cclients.HandleHttpResponse[map[string]bool](res, correlationId)

	if err != nil {
		return false, err
	}

	val, valOk := result["valid"]
	if !valOk {
		return false, nil
	}
	return val, nil
}

func (c *PasswordsCommandableHttpClientV1) ResetPassword(ctx context.Context, correlationId string, userId string,
	code string, password string) error {
	params := cdata.NewAnyValueMapFromTuples(
		"user_id", userId,
		"code", code,
		"password", password,
	)
	_, err := c.CallCommand(ctx, "reset_password", correlationId, params)
	return err
}

func (c *PasswordsCommandableHttpClientV1) RecoverPassword(ctx context.Context, correlationId string, userId string) error {
	params := cdata.NewAnyValueMapFromTuples(
		"user_id", userId,
	)
	_, err := c.CallCommand(ctx, "recover_password", correlationId, params)
	return err

}
