package model

type Superhero struct {
	ID                    string  `db:"id"`
	Email                 string  `db:"email"`
	Name                  string  `db:"name"`
	SuperheroName         string  `db:"superheroName"`
	MainProfilePicURL     string  `db:"mainProfilePicUrl"`
	Gender                int     `db:"gender"`
	LookingForGender      int     `db:"lookingForGender"`
	Age                   int     `db:"age"`
	LookingForAgeMin      int     `db:"lookingForAgeMin"`
	LookingForAgeMax      int     `db:"lookingForAgeMax"`
	LookingForDistanceMax int     `db:"lookingForDistanceMax"`
	DistanceUnit          string  `db:"distanceUnit"`
	Lat                   float64 `db:"lat"`
	Lon                   float64 `db:"lon"`
	Birthday              string  `db:"birthday"`
	Country               string  `db:"country"`
	City                  string  `db:"city"`
	SuperPower            string  `db:"superpower"`
	AccountType           string  `db:"accountType"`
	IsDeleted             bool    `db:"isDeleted"`
	IsBlocked             bool    `db:"isBlocked"`
	UpdatedAt             string  `db:"updatedAt"`
	CreatedAt             string  `db:"createdAt"`
	BlockedAt             string  `db:"blockedAt"`
	DeletedAt             string  `db:"deletedAt"`
}
