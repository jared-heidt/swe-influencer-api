package models

type Creator struct {
	ID         uint32 `json:"id"`
	Name       string `json:"name"`
	Profession string `json:"profession"`
	Focus      string `json:"focus"`
	Youtube    string `json:"youtube"`
}

func GetAllCreators() ([]Creator, error) {
	var creators []Creator

	query := `SELECT * FROM creator`

	rows, err := db.Query(query)

	if err != nil {
		return creators, err
	}

	defer rows.Close()

	for rows.Next() {
		var id uint32
		var name, profession, focus, youtube string

		err := rows.Scan(&id, &name, &profession, &focus, &youtube)

		if err != nil {
			return creators, err
		}

		creator := Creator{
			ID:         id,
			Name:       name,
			Profession: profession,
			Focus:      focus,
			Youtube:    youtube,
		}

		creators = append(creators, creator)

	}

	return creators, nil
}
