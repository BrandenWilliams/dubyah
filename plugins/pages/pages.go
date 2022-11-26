package plugin

import "github.com/hoisie/mustache"

type Pages struct {
	Homepage       mustache.Template
	TechSupport    mustache.Template
	TaskManagement mustache.Template
	Websites       mustache.Template
	Resume         mustache.Template
	NotFound       mustache.Template
}
