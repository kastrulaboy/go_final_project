package db

import ("fmt")

type Task struct {
	ID      string `json:"id"`
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat"`
}

func AddTask(task *Task) (int64, error) {
	query := `
		INSERT INTO scheduler(date, title, comment, repeat)
		VALUES (?, ?, ?, ?)
	`

	res, err := db.Exec(
		query,
		task.Date,
		task.Title,
		task.Comment,
		task.Repeat,
	)

	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func Tasks(limit int) ([]*Task, error) {

	rows, err := db.Query(`
		SELECT id, date, title, comment, repeat
		FROM scheduler
		ORDER BY date
		LIMIT ?
	`, limit)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*Task

	for rows.Next() {
		var t Task

		err := rows.Scan(
			&t.ID,
			&t.Date,
			&t.Title,
			&t.Comment,
			&t.Repeat,
		)

		if err != nil {
			return nil, err
		}

		tasks = append(tasks, &t)
	}

	if tasks == nil {
		tasks = []*Task{}
	}

	return tasks, nil
}

func GetTask(id string) (*Task, error) {
	row := db.QueryRow(`
		SELECT id, date, title, comment, repeat
		FROM scheduler
		WHERE id = ?
		`, id)

	var t Task

	err := row.Scan(
		&t.ID,
		&t.Date,
		&t.Title,
		&t.Comment,
		&t.Repeat,
	)

	if err != nil {
		return nil, err
	}

	return &t, nil
}

func DeleteTask(id string) error {
    res, err := db.Exec(
        `
        DELETE FROM scheduler
        WHERE id = ?
        `,
        id,
    )
    if err != nil {
        return err
    }

    count, err := res.RowsAffected()
    if err != nil {
        return err
    }

    if count == 0 {
        return fmt.Errorf("ошибка")
    }

    return nil
}

func UpdateDate(next string, id string) error {
    res, err := db.Exec(
        `
        UPDATE scheduler
        SET date = ?
        WHERE id = ?
        `,
        next,
        id,
    )
    if err != nil {
        return err
    }

    count, err := res.RowsAffected()
    if err != nil {
        return err
    }

    if count == 0 {
        return fmt.Errorf("ошибка")
    }

    return nil
}

func UpdateTask(task *Task) error {
	res, err := db.Exec(`
		UPDATE scheduler
		SET date = ?, title = ?, comment = ?, repeat = ?
		WHERE id = ?
	`,
		task.Date,
		task.Title,
		task.Comment,
		task.Repeat,
		task.ID,
	)

	if err != nil {
		return err
	}

	n, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if n == 0 {
		return fmt.Errorf("ошибка")
	}

	return nil
}