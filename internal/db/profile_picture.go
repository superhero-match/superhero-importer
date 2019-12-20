package db

import (
	"github.com/superhero-importer/internal/db/model"
)

// GetProfilePictures fetches profile pictures for the Superheros ids.
func (db *DB) GetProfilePictures(ids string) (profilePictures map[string][]model.ProfilePicture, err error) {
	profilePics := make([]model.ProfilePicture, 0)

	err = db.stmtGetProfilePictures.Select(&profilePics, ids)
	if err != nil {
		return nil, err
	}

	profilePictures = make(map[string][]model.ProfilePicture)

	for _, profilePicture := range profilePics {
		profilePictures[profilePicture.SuperheroID] = append(profilePictures[profilePicture.SuperheroID], profilePicture)
	}

	return profilePictures, nil
}
