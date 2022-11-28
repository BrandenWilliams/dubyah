package plugin

import "github.com/hoisie/mustache"

type Pages struct {
	Homepage    mustache.Template
	TechSupport mustache.Template
	Websites    mustache.Template
	Resume      mustache.Template
	NotFound    mustache.Template

	// Onboarding pages
	SignUp mustache.Template
	Login  mustache.Template

	// Taskpages
	TaskListsManagement mustache.Template
	Tasklist            mustache.Template
}
