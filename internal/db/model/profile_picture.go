package model

// ProfilePicture holds data related to profile picture of a Superhero.
type ProfilePicture struct {
	ID                int64  `db:"id"`
	SuperheroID       string `db:"superhero_id"`
	ProfilePictureURL string `db:"profile_pic_url"`
	Position          int    `db:"position"`
}
