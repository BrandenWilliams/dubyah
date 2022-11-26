package meta

type Meta struct {
	ViewingUserID string
	CurrentURL    string
	Query         string

	IsLoggedIn bool
	IsAdmin    bool
}
