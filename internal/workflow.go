package internal

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

type RegisterFunc func(email, password string) error

type LoginFunc func(email, password string) (Token, error)

// Register adds a new user to cognito user pools
func Register(c Config) RegisterFunc {

	return func(email, password string) error {
		sess, err := session.NewSession(&aws.Config{
			Region: aws.String("region"),
		})
		if err != nil {
			log.Println(err)
			return err
		}

		cognitoClient := cognitoidentityprovider.New(sess)
		newUserData := &cognitoidentityprovider.AdminCreateUserInput{
			DesiredDeliveryMediums: []*string{
				aws.String(email),
			},
			UserAttributes: []*cognitoidentityprovider.AttributeType{
				{
					Name:  aws.String("email"),
					Value: aws.String(email),
				},
			},
			UserPoolId:        &c.UserPoolID,
			Username:          &email,
			TemporaryPassword: &password,
		}

		_, err = cognitoClient.AdminCreateUser(newUserData)
		return err
	}
}

// Login provides access and refresh token details from cognito
func Login(c Config) LoginFunc {

	return func(email, password string) (Token, error) {
		sess, err := session.NewSession(&aws.Config{
			Region: aws.String("region"),
		})
		if err != nil {
			log.Println(err)
			return Token{}, err
		}

		params := &cognitoidentityprovider.InitiateAuthInput{
			AuthFlow: aws.String("USER_PASSWORD_AUTH"),
			AuthParameters: map[string]*string{
				"USERNAME": aws.String(email),
				"PASSWORD": aws.String(password),
			},
			ClientId: &c.AppClientID,
		}
		cognitoClient := cognitoidentityprovider.New(sess)
		resp, err := cognitoClient.InitiateAuth(params)

		if resp.ChallengeName == aws.String("NEW_PASSWORD_REQUIRED") {

			newParams := &cognitoidentityprovider.AdminRespondToAuthChallengeInput{
				ChallengeName: aws.String("NEW_PASSWORD_REQUIRED"),
				ClientId:      &c.AppClientID,
				UserPoolId:    &c.UserPoolID,
				ChallengeResponses: map[string]*string{
					"USERNAME":     aws.String(email),
					"NEW_PASSWORD": aws.String(password),
				},
				Session: resp.Session,
			}

			challengeResp, err := cognitoClient.AdminRespondToAuthChallenge(newParams)
			if err != nil {
				log.Println(err)
				return Token{}, err
			}
			return ToTokenResult(challengeResp.AuthenticationResult), nil

		}

		return ToTokenResult(resp.AuthenticationResult), nil
	}
}
