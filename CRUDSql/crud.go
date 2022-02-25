package CRUDSql

import (
	"database/sql"

	"errors"

	_ "github.com/go-sql-driver/mysql"
)

type Empl struct {
	id    int
	Name  string
	Email string
	role  string
}

// func DbConn(db_name string) (db *sql.DB) {
// 	dbDriver := "mysql"
// 	dbUser := "root"
// 	dbPass := "password"
// 	dbName := db_name
// 	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	return db
// }

// // func CreateTable()
// func CreateTable(db_name string, table_name string) {
// 	db := DbConn("Employee_Db")
// 	defer db.Close()
// 	query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %v(id int PRIMARY KEY AUTO_INCREMENT, Name varchar(30) NOT NULL, Email varchar(30), role varchar(30));", table_name)
// 	fmt.Println(query)
// 	res, err := db.Exec(query)
// 	if err != nil {
// 		fmt.Println("here")
// 		log.Fatal(err)
// 	}
// 	fmt.Println(res.RowsAffected())
// }

// // Create
// func InsertData(table_name string, db *sql.DB, name []string, email []string, role []string) error {
// 	query := "INSERT INTO Employee_Details (Name,Email,role) values (?,?,?);"

// 	res, err := db.Prepare(query)
// 	if err != nil {
// 		return err
// 	}

// 	for i := 0; i < len(name); i++ {
// 		_, err2 := res.Exec(name[i], email[i], role[i])

// 		if err2 != nil {
// 			return err2
// 		}
// 	}
// 	return nil
// }

func InsertData(table_name string, db *sql.DB, name string, email string, role string) error {
	query := "INSERT INTO Employee_Details (Name,Email,role) values (?,?,?);"

	res, err := db.Prepare(query)
	if err != nil {
		return errors.New("Prepare Error")
	}

	_, err = res.Exec(name, email, role)

	return err

}

// // Read
// func GetDetailsById(db *sql.DB, id int) (*Empl, error) {

// 	res, err := db.Prepare("select * from Employee_Details where id=?")

// 	if err != nil {
// 		return nil, err
// 	}

// 	defer res.Close()

// 	var e Empl
// 	err2 := res.QueryRow(id).Scan(&e.id, &e.Name, &e.Email, &e.role)

// 	if err2 != nil {
// 		return nil, err2
// 	}

// 	return &e, nil
// }

func GetById(db *sql.DB, id int) (*Empl, error) {
	query := "select * from Employee_Details where id=?"
	res, err := db.Prepare(query)
	var e Empl

	if err != nil {
		return nil, errors.New("")
	}

	err2 := res.QueryRow(id).Scan(&e.id, &e.Name, &e.Email, &e.role)

	if err2 != nil {
		return nil, err2
	}

	return &e, nil

}

// // Update
func UpdateById(db *sql.DB, id int, Name string, Email string, role string) error {

	res, err := db.Prepare("update Employee_Details set Name=?, Email=?, role=? where id=?")
	if err != nil {
		return errors.New("Prepare Error")
	}

	defer res.Close()

	_, err = res.Exec(Name, Email, role, id)

	return err

}

// // Delete
func DeleteById(db *sql.DB, id int) error {

	res, err := db.Prepare("delete from Employee_Details where id=?")

	if err != nil {
		return errors.New("Prepare Error")
	}

	defer res.Close()
	_, err = res.Exec(id)

	return err

}

// func GetAll(db *sql.DB, table_name string) {

// 	query := "Select * from " + table_name
// 	res, err := db.Prepare(query)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	resp, err := res.Query()
// 	var (
// 		id    int
// 		name  string
// 		email string
// 		role  string
// 	)
// 	for resp.Next() {
// 		err := resp.Scan(&id, &name, &email, &role)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		fmt.Printf("Id : %v, Name: %v, Email: %v,Role: %v\n", id, name, email, role)
// 	}
// }
