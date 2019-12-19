package model

// ProfilePicture holds data related to profile picture of a Superhero.
type ProfilePicture struct {
	SuperheroID       string `db:"superhero_id,"`
	ProfilePictureURL string `db:"profile_pic_url"`
	Position          int8   `db:"position"`
	UpdatedAt         string `db:"updated_at"`
	CreatedAt         string `db:"created_at"`
	DeletedAt         string `db:"deleted_at"`
}
