package repository

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type Repository struct {
	DB *sql.DB
}

var repo Repository

func Init() { /* Attempts to create a database using sqlite3, failing if not */
	if db, err := sql.Open("sqlite3", "test.db"); err == nil { 
		repo = Repository{DB: db}
	} else {
		log.Fatal("Database initialisation")
	}
}

func Create() int { /* Creates a database */
	const sql = "CREATE TABLE IF NOT EXISTS Cells" +
		    "(Id TEXT PRIMARY KEY, Audio TEXT)" /* Runs this command, returning 0 if the table is created, and -1 if not */
	if _, err := repo.DB.Exec(sql); err == nil {
		return 0
	} else {
		return -1
	}
}


func Clear() int { /* Deletes the database*/
	const sql = "DELETE FROM Cells" /* Runs this command */
	if _, err := repo.DB.Exec(sql); err == nil { /* Returning 0 if successful, and -1 if not */
		return 0
	} else {
		return -1
	}
}

func Read(id string) (Cell, int64) { /* Reads a cell with a given id */
	const sql = "SELECT * FROM Cells WHERE Id = ?" /* Create a stmt asking based on ID */
	if stmt, err := repo.DB.Prepare(sql); err == nil { 
		defer stmt.Close()
		var c Cell /* Otherwise */
		row := stmt.QueryRow(id) /* Query the statement */
		if err := row.Scan(&c.Id, &c.Audio); err == nil { /* If there are no errors */
			return c, 1 /* Return the cell and 1 if the track is found */
		} else {
			return Cell{}, 0 /* Return an empty cell and 0 if the track is not found */
		}
	}
	return Cell{}, -1 /* Return -1 and an empty cell if there was en error */
}

func Update(c Cell) int64 {
	const sql = "UPDATE Cells SET Audio = ?" + /* Update track data with the given value */
		  "WHERE id = ?"
	if stmt, err := repo.DB.Prepare(sql); err == nil { /* Prepare the statement */
		defer stmt.Close() /* Close the statement once it is done */
		if res, err := stmt.Exec(c.Audio, c.Id); err == nil { /* Execute the statement*/
			if n, err := res.RowsAffected(); err == nil { /* If there are no errors from execution and checking the affected rows */
				return n /* Return the number of affected rows */
			}
		}
	}
	return -1 /* Return -1 if there were errors */
}

func Insert(c Cell) int64 {
	const sql = "INSERT INTO Cells(Id, Audio)" + /* Insert track data with the given value and id */
		  "VALUES (?,?)"
	if stmt, err := repo.DB.Prepare(sql); err == nil { /* Prepare the statement */
		defer stmt.Close() /* Close the statement once it is done */
		if res, err := stmt.Exec(c.Id, c.Audio); err == nil { /* Execute the statement*/
			if n, err := res.RowsAffected(); err == nil { /* If there are no errors from execution and checking the affected rows */
				return n /* Return the number of affected rows */
			}
		}
	}
	return -1 /* Return -1 if there were errors */
}

func GetIDs() ([]string, int64) {
	const sql = "SELECT * FROM Cells" /* Select every cell with this statement */
	if stmt, err := repo.DB.Prepare(sql); err == nil { /* Prepare the statement whilst checking for errors */
		defer stmt.Close() /* Close the statement at the end */
		rows, _ := stmt.Query() /* Run the statement, storing the retrieved cells in rows */
		defer rows.Close() /* Close the row enumerator at the end */
		ids := make([]string, 0) /* Make an empty slice */
		for rows.Next() { /* Loop through the results */
			var id string
			var audio string
			rows.Scan(&id, &audio) /* Retrieve the id and audio from each row */
			ids = append(ids, id) /* Append the id to the list */
		}
		return ids, 1 /* Return the list once finished */
	}
	return make([]string, 0), -1 /* Return an empty list and -1 when there is a failure */
}

func Delete(c Cell) int64 {
	const sql = "DELETE FROM Cells WHERE id = ?" /* Update track data with the given value */
	if stmt, err := repo.DB.Prepare(sql); err == nil { /* Prepare the statement */
		defer stmt.Close() /* Close the statement once it is done */
		if res, err := stmt.Exec(c.Id); err == nil { /* Execute the statement*/
			if n, err := res.RowsAffected(); err == nil { /* If there are no errors from execution and checking the affected rows */
				return n /* Return the number of affected rows */
			}
		}
	}
	return -1 /* Return -1 if there were errors */
}