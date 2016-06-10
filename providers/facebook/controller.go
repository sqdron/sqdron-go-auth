package facebook

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
	//app "../../app"
	p "../../providers"
)

var conf = &oauth2.Config{
	ClientID:     "",
	ClientSecret: "",
	RedirectURL:  "http://localhost:8080/",
	Scopes:       []string{"public_profile"},
	Endpoint:     facebook.Endpoint,
}

//func (a *app.Application) UseFacebookAuth(){
//
//}

type FacebookProvider struct {
	p.Provider
}

type facebookController interface {
	//handleMain()
	Login()
	//handleFacebookCallback()
}


