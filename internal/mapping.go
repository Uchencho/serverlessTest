package internal

import "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int64  `json:"expires_in"`
	TokenType    string `json:"token_type"`
	IDToken      string `json:"id_token"`
}

func ToTokenResult(r *cognitoidentityprovider.AuthenticationResultType) Token {
	return Token{
		AccessToken:  *r.AccessToken,
		RefreshToken: *r.RefreshToken,
		ExpiresIn:    *r.ExpiresIn,
		TokenType:    *r.TokenType,
		IDToken:      *r.IdToken,
	}
}
