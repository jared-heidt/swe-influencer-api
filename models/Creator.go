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

func GetCreator(id uint32) (Creator, error) {
	var creator Creator
	query := `SELECT name, profession, focus, youtube FROM creator WHERE id = $1`
	row, err := db.Query(query, id)
	if err != nil {
		return creator, err
	}
	defer row.Close()
	if row.Next() {
		var name, profession, focus, youtube string
		err := row.Scan(&name, &profession, &focus, &youtube)
		if err != nil {
			return creator, err
		}
		creator = Creator{
			ID:         id,
			Name:       name,
			Profession: profession,
			Focus:      focus,
			Youtube:    youtube,
		}
	}
	return creator, nil
}

func CreateCreator(creator Creator) error {
	query := `INSERT INTO creator(name, profession, focus, youtube) VALUES ($1, $2, $3, $4)`
	_, err := db.Exec(query, creator.Name, creator.Profession, creator.Focus, creator.Youtube)
	if err != nil {
		return err
	}
	return nil
}

func UpdateCreator(creator Creator) error {
	query := `UPDATE creator SET name=$1, profession=$2, focus=$3, youtube=$4 WHERE id=$5;`
	_, err := db.Exec(query, creator.Name, creator.Profession, creator.Focus, creator.Youtube, creator.ID)
	if err != nil {
		return err
	}
	return nil
}

func DeleteCreator(id uint32) error {
	query := `DELETE FROM posts WHERE id=$1;`
	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
