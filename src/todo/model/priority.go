package model

type Priority struct {
	PriorityId int
	Value      string
}

// Priorities get all Priorities from the database
func Priorities() (priorities []Priority, err error) {
	rows, err := Database.Query("select priority_id, value from priorities")
	if err != nil {
		return priorities, err
	}

	for rows.Next() {
		priority := Priority{}
		err := rows.Scan(
			&priority.PriorityId,
			&priority.Value,
		)
		if err != nil {
			return priorities, err
		}

		priorities = append(priorities, priority)
	}
	rows.Close()

	return priorities, err
}

// PriorityById return a single priority that matches the provided id
func PriorityById(id int) (priority Priority, err error) {
	priority = Priority{}
	err = Database.QueryRow(
		"select priority_id, value from priorities where priority_id = $1 order by value", id,
	).Scan(
		&priority.PriorityId,
		&priority.Value,
	)

	return priority, err
}
