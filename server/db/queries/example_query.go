package queries

import (
	"fmt"

	database "github.com/aasumitro/govue/server/db"
	"github.com/aasumitro/govue/server/entities"
)

func All() []entities.Example {
	db, err := database.Connect()
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	defer db.Close()

	rows, err := db.Query("select * from examples")
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	defer rows.Close()

	var result []entities.Example

	for rows.Next() {
		var each = entities.Example{}
		var err = rows.Scan(&each.Id, &each.Title, &each.Description)

		if err != nil {
			fmt.Println(err.Error())
			return nil
		}

		result = append(result, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return result
}

func Find(id int) entities.Example {
	db, err := database.Connect()
	if err != nil {
		fmt.Println(err.Error())
		return entities.Example{}
	}
	defer db.Close()

	var result = entities.Example{}
	err = db.QueryRow("select * from examples where id = ?", id).
		Scan(&result.Id, &result.Title, &result.Description)
	if err != nil {
		fmt.Println(err.Error())
		return entities.Example{}
	}

	return result
}

func Destroy(id int) (int64, error) {
	db, err := database.Connect()
	if err != nil {
		fmt.Println(err.Error())
		return 0, nil
	}
	defer db.Close()

	result, err := db.Exec("delete from examples where id = ?", id)
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}

	return result.RowsAffected()
}

func Update(id int, title string, description string) (int64, error) {
	db, err := database.Connect()
	if err != nil {
		fmt.Println(err.Error())
		return 0, nil
	}
	defer db.Close()

	result, err := db.Exec("update examples set title = ?, description = ? where id = ?", title, description, id)
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}

	return result.RowsAffected()
}

func Store(title string, description string) bool {
	db, err := database.Connect()
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	defer db.Close()

	if title == "" && description == "" {
		return false
	}

	_, err = db.Exec("insert into examples values (?, ?, ?)", nil, title, description)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	return true
}
