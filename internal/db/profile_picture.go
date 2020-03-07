/*
  Copyright (C) 2019 - 2020 MWSOFT
  This program is free software: you can redistribute it and/or modify
  it under the terms of the GNU General Public License as published by
  the Free Software Foundation, either version 3 of the License, or
  (at your option) any later version.
  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU General Public License for more details.
  You should have received a copy of the GNU General Public License
  along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/
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
