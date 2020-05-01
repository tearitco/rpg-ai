package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	googleAuthIDTokenVerifier "github.com/futurenda/google-auth-id-token-verifier"
)

type GoogleLoginRequest struct {
	TokenID string
}

type FacebookLoginRequest struct {
	AccessToken string
}

type LoginResponse struct {
	User *User
}

type APIService struct {
	db *Database
}

func (s *APIService) GoogleLogin(r *http.Request, args *GoogleLoginRequest, reply *LoginResponse) error {
	v := googleAuthIDTokenVerifier.Verifier{}
	if err := v.VerifyIDToken(args.TokenID, []string{os.Getenv("GOOGLE_CLIENT_ID")}); err != nil {
		return err
	}
	claimSet, err := googleAuthIDTokenVerifier.Decode(args.TokenID)
	if err != nil {
		return err
	}
	user, err := s.db.GetUserByEmail(claimSet.Email)
	if err != nil {
		return err
	}
	reply.User = user
	authenticatedUser := r.Context().Value(ContextAuthenticatedUserKey).(*AuthenticatedUser)
	authenticatedUser.InternalUser = user
	authenticatedUser.GoogleUser = claimSet
	return nil
}

func (s *APIService) FacebookLogin(r *http.Request, args *FacebookLoginRequest, reply *LoginResponse) error {
	url := "https://graph.facebook.com/v3.2/me?fields=email&access_token=" + args.AccessToken
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fbResp := struct {
		ID    string `json:"id"`
		Email string `json:"email"`
	}{}
	if err := json.Unmarshal(body, &fbResp); err != nil {
		return err
	}
	user, err := s.db.GetUserByEmail(fbResp.Email)
	if err != nil {
		return err
	}
	reply.User = user
	authenticatedUser := r.Context().Value(ContextAuthenticatedUserKey).(*AuthenticatedUser)
	authenticatedUser.InternalUser = user
	authenticatedUser.FacebookUser = fbResp
	return nil
}
