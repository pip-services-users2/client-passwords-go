package version1

import (
	"context"
	"crypto/sha256"
	"time"

	cconv "github.com/pip-services3-gox/pip-services3-commons-gox/convert"
	cdata "github.com/pip-services3-gox/pip-services3-commons-gox/data"
	cerr "github.com/pip-services3-gox/pip-services3-commons-gox/errors"
	crnd "github.com/pip-services3-gox/pip-services3-commons-gox/random"
)

type PasswordsMockClientV1 struct {
	passwords        []UserPasswordV1
	lockTimeout      int64
	attemptTimeout   int64
	attemptCount     int32
	recExpireTimeout int64
	lockEnabled      bool
	magicCode        string
	code_length      int
}

func NewPasswordsMockClientV1() *PasswordsMockClientV1 {
	c := PasswordsMockClientV1{
		passwords:        make([]UserPasswordV1, 0),
		lockTimeout:      1800000,      // 30 mins
		attemptTimeout:   60000,        // 1 min
		attemptCount:     4,            // 4 times
		recExpireTimeout: 24 * 3600000, // 24 hours
		lockEnabled:      false,
		magicCode:        "",
		code_length:      9, // Generated code length
	}
	return &c
}

func (c *PasswordsMockClientV1) GetPasswordInfo(ctx context.Context, correlationId string, userId string) (result *UserPasswordInfoV1, err error) {

	var index = -1
	for k, v := range c.passwords {
		if v.Id == userId {
			index = k
			break
		}
	}

	if index < 0 {
		err := cerr.NewNotFoundError(
			correlationId,
			"USER_NOT_FOUND",
			"User "+userId+" was not found",
		).WithDetails("user_id", userId)

		return nil, err
	}
	data := c.passwords[index]

	info := UserPasswordInfoV1{
		Id:         data.Id,
		ChangeTime: data.ChangeTime,
		Locked:     data.Locked,
		LockTime:   data.LockTime,
	}

	return &info, nil
}

func (c *PasswordsMockClientV1) SetTempPassword(ctx context.Context, correlationId string, userId string) (password string, err error) {

	password = "pass" + cconv.StringConverter.ToString(crnd.Integer.Next(0, 1000))

	password = c.hashPassword(password)
	userPassword := *NewUserPasswordV1(userId, password)
	userPassword.ChangeTime = time.Now().UTC()
	c.passwords = append(c.passwords, userPassword)
	return password, nil
}

func (c *PasswordsMockClientV1) SetPassword(ctx context.Context, correlationId string, userId string, password string) error {
	password = c.hashPassword(password)
	userPassword := *NewUserPasswordV1(userId, password)
	c.passwords = append(c.passwords, userPassword)
	return nil
}

func (c *PasswordsMockClientV1) DeletePassword(ctx context.Context, correlationId string, userId string) error {
	var index = -1
	for i, v := range c.passwords {
		if v.Id == userId {
			index = i
			break
		}
	}

	if index < 0 {
		return nil
	}

	if index == len(c.passwords) {
		c.passwords = c.passwords[:index-1]
	} else {
		c.passwords = append(c.passwords[:index], c.passwords[index+1:]...)
	}
	return nil
}

func (c *PasswordsMockClientV1) Authenticate(ctx context.Context, correlationId string, userId string, password string) (authenticated bool, err error) {
	hashedPassword := c.hashPassword(password)
	currentTime := time.Now().UTC()

	// Retrieve user password
	userPassword, err := c.readUserPassword(ctx, correlationId, userId)

	// Check password and process failed attempts

	passwordMatch := userPassword.Password == hashedPassword
	var lastFailureTimeout time.Duration
	if !userPassword.FailTime.IsZero() {
		lastFailureTimeout = currentTime.Sub(userPassword.FailTime)
	}

	//verify user account is still locked from last authorization failure or just tell user that it"s user is locked
	if !c.lockEnabled && passwordMatch {
		userPassword.Locked = false
	} else {
		if passwordMatch && userPassword.Locked && lastFailureTimeout > time.Duration(c.lockTimeout*time.Millisecond.Milliseconds()) {
			userPassword.Locked = false //unlock user
		} else if userPassword.Locked {

			err = cerr.NewBadRequestError(
				correlationId,
				"ACCOUNT_LOCKED",
				"Account for user "+userId+" is locked",
			).WithDetails("user_id", userId)
			return false, err
		}

		if !passwordMatch {
			if lastFailureTimeout < time.Duration(c.attemptTimeout*time.Millisecond.Milliseconds()) {
				userPassword.FailCount = userPassword.FailCount + 1
			}

			userPassword.FailTime = currentTime

			if userPassword.FailCount >= c.attemptCount {
				userPassword.Locked = true

				err = cerr.NewBadRequestError(
					correlationId,
					"ACCOUNT_LOCKED",
					"Number of attempts exceeded. Account for user "+userId+" was locked",
				).WithDetails("user_id", userId)

				return false, err

			} else {

				err = cerr.NewBadRequestError(
					correlationId,
					"WRONG_PASSWORD",
					"Invalid password",
				).WithDetails("user_id", userId)
				return false, err

			}

			var index = -1
			for i, v := range c.passwords {
				if v.Id == userPassword.Id {
					index = i
					break
				}
			}

			if index < 0 {
				return false, nil
			}

			c.passwords[index] = *userPassword

		}
	}

	// Perform authentication and save user
	// Update user last signin date
	userPassword.FailCount = 0
	userPassword.FailTime = time.Time{}

	var index = -1
	for i, v := range c.passwords {
		if v.Id == userPassword.Id {
			index = i
			break
		}
	}

	if index < 0 {
		return false, nil
	}
	c.passwords[index] = *userPassword

	return userPassword != nil, nil
}

func (c *PasswordsMockClientV1) ChangePassword(ctx context.Context, correlationId string, userId string,
	oldPassword string, newPassword string) error {

	var userPassword *UserPasswordV1

	verify, err := c.verifyPassword(ctx, correlationId, newPassword)
	if !verify || err != nil {
		return nil
	}

	oldPassword = c.hashPassword(oldPassword)
	newPassword = c.hashPassword(newPassword)

	// Retrieve user

	userPassword, err = c.readUserPassword(ctx, correlationId, userId)

	// Verify and reset password

	// Password must be different then the previous one
	if userPassword.Password != oldPassword {

		err = cerr.NewBadRequestError(
			correlationId,
			"WRONG_PASSWORD",
			"Invalid password",
		).WithDetails("user_id", userId)

		return err
	}

	if oldPassword == newPassword {

		err = cerr.NewBadRequestError(
			correlationId,
			"PASSWORD_NOT_CHANGED",
			"Old and new passwords are identical",
		).WithDetails("user_id", userId)

		return err
	}

	// Reset password
	userPassword.Password = newPassword
	userPassword.RecCode = ""
	userPassword.RecExpireTime = time.Time{}
	userPassword.Locked = false
	// Todo: Add change password policy
	userPassword.ChangeTime = time.Time{}

	// Save the new password
	var index = -1
	for i, v := range c.passwords {
		if v.Id == userPassword.Id {
			index = i
			break
		}
	}

	if index < 0 {
		return nil
	}
	c.passwords[index] = *userPassword

	return nil
}

func (c *PasswordsMockClientV1) ValidateCode(ctx context.Context, correlationId string, userId string,
	code string) (valid bool, err error) {

	data, err := c.readUserPassword(ctx, correlationId, userId)

	if err == nil && data != nil {
		valid := data.RecCode == code && data.RecExpireTime.UnixNano() > time.Now().UTC().UnixNano()
		return valid, nil
	} else {
		return false, nil
	}

}

func (c *PasswordsMockClientV1) ResetPassword(ctx context.Context, correlationId string, userId string, code string, password string) error {

	var userPassword *UserPasswordV1

	verify, err := c.verifyPassword(ctx, correlationId, password)
	if !verify || err != nil {
		return nil
	}

	password = c.hashPassword(password)
	// Retrieve user
	userPassword, err = c.readUserPassword(ctx, correlationId, userId)
	// Validate reset code and reset the password
	if userPassword.RecCode != code && code != c.magicCode {

		err = cerr.NewBadRequestError(
			correlationId,
			"WRONG_CODE",
			"Invalid password recovery code "+code,
		).WithDetails("user_id", userId)

		return err
	}

	// Check if code already expired
	if !(userPassword.RecExpireTime.Unix() > time.Now().UTC().Unix()) && code != c.magicCode {

		err = cerr.NewBadRequestError(
			correlationId,
			"CODE_EXPIRED",
			"Password recovery code "+code+" expired",
		).WithDetails("user_id", userId)

		return err
	}

	// Reset the password
	userPassword.Password = password
	userPassword.RecCode = ""
	userPassword.RecExpireTime = time.Time{}
	userPassword.Locked = false
	// Todo: Add change password policy
	userPassword.ChangeTime = time.Time{}

	// Save the new password
	var index = -1
	for i, v := range c.passwords {
		if v.Id == userPassword.Id {
			index = i
			break
		}
	}

	if index < 0 {
		return nil
	}
	c.passwords[index] = *userPassword
	return nil
}

func (c *PasswordsMockClientV1) RecoverPassword(ctx context.Context, correlationId string, userId string) error {

	// Retrieve user
	userPassword, err := c.readUserPassword(ctx, correlationId, userId)

	if err != nil {
		return err
	}

	// Update and save recovery code
	currentTicks := time.Now().UnixNano() / time.Millisecond.Milliseconds()
	expireTicks := (currentTicks + c.recExpireTimeout) * time.Millisecond.Milliseconds()
	expireTime := time.Time{}
	expireTime.Add(time.Duration(expireTicks))

	userPassword.RecCode = c.generateVerificationCode()
	userPassword.RecExpireTime = expireTime

	var index = -1
	for i, v := range c.passwords {
		if v.Id == userPassword.Id {
			index = i
			break
		}
	}

	if index < 0 {
		return nil
	}
	c.passwords[index] = *userPassword
	return nil

}

func (c *PasswordsMockClientV1) readUserPassword(ctx context.Context, correlationId string, userId string) (result *UserPasswordV1, err error) {

	var index = -1
	for k, v := range c.passwords {
		if v.Id == userId {
			index = k
			break
		}
	}

	if index < 0 {
		err := cerr.NewNotFoundError(
			correlationId,
			"USER_NOT_FOUND",
			"User "+userId+" was not found",
		).WithDetails("user_id", userId)

		return nil, err
	}

	item := c.passwords[index]
	return &item, nil
}

func (c *PasswordsMockClientV1) generateVerificationCode() string {
	return cdata.IdGenerator.NextShort()[0:6]
}

func (c *PasswordsMockClientV1) hashPassword(password string) string {
	if password == "" {
		return ""
	}
	hash := sha256.New()
	sum := hash.Sum(([]byte)(password))
	return (string)(sum)
}

func (c *PasswordsMockClientV1) verifyPassword(ctx context.Context, correlationId string, password string) (res bool, err error) {

	if password == "" {
		err = cerr.NewBadRequestError(
			correlationId,
			"NO_PASSWORD",
			"Missing user password",
		)
		return false, err
	}

	if len(password) < 6 || len(password) > 20 {

		err = cerr.NewBadRequestError(
			correlationId,
			"BAD_PASSWORD",
			"User password should be 5 to 20 symbols long",
		)
		return false, err
	}
	return true, nil
}
