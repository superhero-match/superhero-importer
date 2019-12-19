package db

import (
	"github.com/superhero-importer/internal/db/model"
)

// GetProfilePictures fetches profile pictures for the Superheros ids.
func (db *DB) GetProfilePictures(ids string) (profilePictures map[string]model.ProfilePicture, err error) {
	rows, err := db.stmtGetProfilePictures.Query(ids)
	if err != nil {
		return nil, err
	}

	profilePictures = make(map[string]model.ProfilePicture)

	for rows.Next() {
		var profilePicture model.ProfilePicture

		err = rows.Scan(&profilePicture)
		if err != nil {
			return nil, err
		}

		profilePictures[profilePicture.SuperheroID] = profilePicture
	}

	return profilePictures, nil
}
