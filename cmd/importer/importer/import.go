package importer

import (
	"fmt"
	"github.com/superhero-importer/internal/es/model"
	"gopkg.in/olivere/elastic.v7"
)

// Import import Superheros data from DB to Elasticsearch.
func (i *Importer) Import() error {
	offset := int64(0)

	for {
		// 1. Fetch the first batch of Superheros from DB.
		superherosDB, err := i.DB.GetSuperheros(offset)
		if err != nil {
			return err
		}

		if len(superherosDB) == 0 {
			break
		}

		// 2. Set offset for the next batch.
		offset += int64(len(superherosDB))

		var ids string

		superheros := make(map[string]model.Superhero)

		// 3. Loop through the DB batch of Superheros.
		for superheroID, superheroDB := range superherosDB {
			// 4. Save all the Superheros ids into a comma-separated-list (one string no spaces).
			// These ids are used in FIND_IN_SET in the stored procedure.
			ids = fmt.Sprintf("%s,%s", ids, superheroID)

			// 5. Map DB Superhero to ES Superhero.
			// Store ES Superheros into a map so it could be found quick and easy
			// when adding profile pictures for the ES Superhero.
			superheros[superheroID] = model.Superhero{
				ID:                    superheroDB.ID,
				Email:                 superheroDB.Email,
				Name:                  superheroDB.Name,
				SuperheroName:         superheroDB.SuperheroName,
				MainProfilePicURL:     superheroDB.MainProfilePicURL,
				ProfilePictures:       make([]model.ProfilePicture, 0),
				Gender:                superheroDB.Gender,
				LookingForGender:      superheroDB.LookingForGender,
				Age:                   superheroDB.Age,
				LookingForAgeMin:      superheroDB.LookingForAgeMin,
				LookingForAgeMax:      superheroDB.LookingForAgeMax,
				LookingForDistanceMax: superheroDB.LookingForDistanceMax,
				DistanceUnit:          superheroDB.DistanceUnit,
				Location: elastic.GeoPoint{
					Lat: superheroDB.Lat,
					Lon: superheroDB.Lon,
				},
				Birthday:    superheroDB.Birthday,
				Country:     superheroDB.Country,
				City:        superheroDB.City,
				SuperPower:  superheroDB.SuperPower,
				AccountType: superheroDB.AccountType,
				CreatedAt:   superheroDB.CreatedAt,
			}
		}

		// 6. Fetch all the Superhero profile pictures.
		profilePicturesDB, err := i.DB.GetProfilePictures(ids)
		if err != nil {
			return err
		}

		superherosES := make([]model.Superhero, 0)

		// 7. Loop through all the DB profile pictures and map them to ES profile picture.
		for superheroID, profilePictures := range profilePicturesDB {
			superhero, ok := superheros[superheroID]
			if !ok {
				return fmt.Errorf("did not find superhero with id: %s in the superheros map", superheroID)
			}

			profilePicturesES := make([]model.ProfilePicture, 0)

			for _, profilePicture := range profilePictures {
				profilePicturesES = append(profilePicturesES, model.ProfilePicture{
					ID:                profilePicture.ID,
					SuperheroID:       profilePicture.SuperheroID,
					ProfilePictureURL: profilePicture.ProfilePictureURL,
					Position:          profilePicture.Position,
				})
			}

			superhero.ProfilePictures = append(superhero.ProfilePictures, profilePicturesES...)

			superherosES = append(superherosES, superhero)
		}

		err = i.ES.StoreSuperheros(superherosES)
		if err != nil {
			return err
		}

		// 8. Check if the len of batch returned is less than the batch size,
		// if so, it means it was the last iteration, after the Superheros
		// have been saved to ES the loop can be broken.
		if len(superherosDB) < i.DB.Limit {
			break
		}
	}

	return nil
}