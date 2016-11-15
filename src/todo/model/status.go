package model

type Status struct {
	StatusId int
	Value    string
}

// Status get all status values from the database
func Statuses() (status []Status, err error) {
	rows, err := Database.Query("select status_id, value from status")
	if err != nil {
		return status, err
	}

	for rows.Next() {
		stat := Status{}
		err := rows.Scan(
			&stat.StatusId,
			&stat.Value,
		)
		if err != nil {
			return status, err
		}

		status = append(status, stat)
	}
	rows.Close()

	return status, err
}

// StatusById return a single status that matches the provided id
func StatusById(id int) (status Status, err error) {
	status = Status{}
	err = Database.QueryRow(
		"select status_id, value from status where status_id = $1", id,
	).Scan(
		&status.StatusId,
		&status.Value,
	)

	return status, err
}
