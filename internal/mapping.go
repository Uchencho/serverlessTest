package internal

import (
	"errors"

	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int64  `json:"expires_in"`
	TokenType    string `json:"token_type"`
	IDToken      string `json:"id_token"`
}

type RegisterUserRequest struct {
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required"`
	ConfirmPassword string `json:"confirm_password" validate:"eqfield=Password"`
}

func ToTokenResult(r *cognitoidentityprovider.AuthenticationResultType) (Token, error) {

	if r == nil {
		return Token{}, errors.New("Nil result type passed")
	}

	return Token{
		AccessToken:  *r.AccessToken,
		RefreshToken: *r.RefreshToken,
		ExpiresIn:    *r.ExpiresIn,
		TokenType:    *r.TokenType,
		IDToken:      *r.IdToken,
	}, nil
}
