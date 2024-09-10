package router

import (
	"net/http"
	"url-shortner/constant"
	"url-shortner/controller"

)

var urlShortner = Routes{
	Route{"Url Shortner Service", http.MethodPost, constant.UrlShortnerPath, controller.ShortTheUrl},
	Route{"Redirect URL", http.MethodGet, constant.RedirectUrlPath, controller.RedirectURL},
}	