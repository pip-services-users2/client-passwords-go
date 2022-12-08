package version1

import "context"

type IPasswordsClientV1 interface {
	GetPasswordInfo(ctx context.Context, correlationId string,
		userId string) (result *UserPasswordInfoV1, err error)

	SetTempPassword(ctx context.Context, correlationId string,
		userId string) (password string, err error)

	SetPassword(ctx context.Context, correlationId string, userId string, password string) error

	DeletePassword(ctx context.Context, correlationId string, userId string) error

	Authenticate(ctx context.Context, correlationId string, userId string,
		password string) (authenticated bool, err error)

	ChangePassword(ctx context.Context, correlationId string, userId string,
		oldPassword string, newPassword string) error

	ValidateCode(ctx context.Context, correlationId string, userId string,
		code string) (valid bool, err error)

	ResetPassword(ctx context.Context, correlationId string, userId string,
		code string, password string) error

	RecoverPassword(ctx context.Context, correlationId string, userId string) error
}
