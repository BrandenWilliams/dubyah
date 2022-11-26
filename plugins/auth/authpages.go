package plugin

import "github.com/hoisie/mustache"

type AuthPages struct {
	Login  mustache.Template
	SignUp mustache.Template
}
