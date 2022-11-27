package plugin

import (
	"github.com/BrandenWilliams/dubyah/plugins/meta"
)

type LoginData struct {
	LoginErr    string
	RedirectURL string
	Meta        *meta.Meta
}
