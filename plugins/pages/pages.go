package plugin

import "github.com/hoisie/mustache"

type Pages struct {
	Homepage      mustache.Template
	TechSupport   mustache.Template
	StackShowcase mustache.Template
	Websites      mustache.Template
	Resume        mustache.Template
}
