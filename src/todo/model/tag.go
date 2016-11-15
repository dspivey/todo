package model

type Tag struct {
	TagId int
	Value string
}

// Tags get all tag values from the database
func Tags() (tags []Tag, err error) {
	rows, err := Database.Query("select tag_id, value from tags")
	if err != nil {
		return tags, err
	}

	for rows.Next() {
		tag := Tag{}
		err := rows.Scan(
			&tag.TagId,
			&tag.Value,
		)
		if err != nil {
			return tags, err
		}

		tags = append(tags, tag)
	}
	rows.Close()

	return tags, err
}

// TagById return a single tag that matches the provided id
func TagById(id int) (tag Tag, err error) {
	tag = Tag{}
	err = Database.QueryRow(
		"select tag_id, value from tags where tag_id = $1", id,
	).Scan(
		&tag.TagId,
		&tag.Value,
	)

	return tag, err
}
