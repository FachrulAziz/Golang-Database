package belajargolangdatabase

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"
)

// Insert ke Database
func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "INSERT INTO customer1 (id, name) VALUES ('6', 'Fachrul Aziz')"
	_, err := db.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}
	fmt.Println("Succees insert customer")
}

func TestQuerySQL(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()

	script := "select id, name from customer1"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	fmt.Println("Berhasil")

	time.Sleep(3 * time.Second)

	for rows.Next() {
		var id, name string
		err = rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println("ID = ", id)
		fmt.Println("Nama =", name)
	}
}

func TestDataColoumn(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "insert into customer2(id, name, email, balance, rating, birth_date, married) values ('1', 'Fachrul Aziz', 'FachrulAziz@gmail.com', '10000', '5.0', '1995-12-28', true),('2', 'Desyana Putri Nurma Intani', 'DesyanaIntani@gmail.com', '9000', '7.0', '1995-12-03', false)"
	_, err := db.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}
	fmt.Println("Insert Succesfull")
}

func TestEditSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "update customer2 set email = 'desyana@gmail.com', birth_date = '1995-12-07' where id = '2'"
	_, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	fmt.Println("Edit Berhasil")
	time.Sleep(3 * time.Second)

	script1 := "SELECT id, name, email, balance, rating, birth_date, created_at, married FROM customer2"
	rows, err := db.QueryContext(ctx, script1)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	time.Sleep(3 * time.Second)

	for rows.Next() {
		var id, name, email string
		var balance int
		var rating float64
		var birth_date, created_at time.Time
		var married bool
		if rows != nil {

		}
		err = rows.Scan(&id, &name, &email, &balance, &rating, &birth_date, &created_at, &married)
		if err != nil {
			panic(err)
		}
		fmt.Println("========================")
		fmt.Println("ID = ", id)
		fmt.Println("Name = ", name)
		fmt.Println("Email = ", email)
		fmt.Println("Balance = ", balance)
		fmt.Println("Rating = ", rating)
		fmt.Println("Birth Date = ", birth_date)
		fmt.Println("Married = ", married)
		fmt.Println("Create At = ", created_at)
	}
}
func TestQuerySqlComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT id, name, email, balance, rating, birth_date, created_at, married FROM customer2"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	// fmt.Println("Mohon tunggu sebentar")
	// time.Sleep(3 * time.Second)

	for rows.Next() {
		var id, name, email string
		var balance int
		var rating float64
		var birth_date, created_at time.Time
		var married bool
		if rows != nil {
			//	fmt.Println("Mohon Tunggu Sebentar")
		}
		time.Sleep(3 * time.Second)
		err = rows.Scan(&id, &name, &email, &balance, &rating, &birth_date, &created_at, &married)
		if err != nil {
			panic(err)
		}
		fmt.Println("========================")
		fmt.Println("ID = ", id)
		fmt.Println("Name = ", name)
		fmt.Println("Email = ", email)
		fmt.Println("Balance = ", balance)
		fmt.Println("Rating = ", rating)
		fmt.Println("Birth Date = ", birth_date)
		fmt.Println("Married = ", married)
		fmt.Println("Create At = ", created_at)
	}
}

func TestSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin"
	password := "admin"

	script := "select username from user where username ='" + username + "' and password='" + password + "' limit 1"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {

	}
	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Login Berhasil", username)
	} else {
		fmt.Println("Username atau password salah")
	}
}

// login dengan parameter
func TestSqlInjectionSafe(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin'; #"
	password := "admin"

	scriptSql := "select username from user where username = ? and password = ? limit 1"
	rows, err := db.QueryContext(ctx, scriptSql, username, password)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Login berhasil", username)
	} else {
		fmt.Println("Username atau Password Salah")
	}
}

// insert dengan parameter
func TestExecSqlParameter(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "daffara"
	password := "ayot"

	script := "insert into user (username, password) values (?, ?)"
	_, err := db.ExecContext(ctx, script, username, password)
	if err != nil {
		panic(err)
	}
	fmt.Println("Berhasil")
}

// insert dengan Auto Increment
func TestAutoIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	email := "fachrulaziz@gmail.com"
	comment := "Test Comentar kedua"

	query := "insert into comments (email, comment) values (?,?)"
	result, err := db.ExecContext(ctx, query, email, comment)
	if err != nil {
		panic(err)
	}
	insertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	fmt.Println("Insert berhasil dengan id =", insertId)
}

func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "insert into comments(email, comment) values (?, ?)"
	stmt, err := db.PrepareContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	for i := 0; i < 10; i++ {
		email := "NauraAlmeera" + strconv.Itoa(i) + "@gmail.com"
		comment := "Komentar ke = " + strconv.Itoa(i)

		result, err := stmt.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}
		fmt.Println("Comment id ", id)
	}
}

func TestTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	script := "insert into comments (email, comment) values (?, ?)"
	for i := 1; i <= 10; i++ {
		email := "Naura " + strconv.Itoa(i) + "@gmail.com"
		comment := "komentar ke =" + strconv.Itoa(i)

		result, err := tx.ExecContext(ctx, script, email, comment)
		if err != nil {
			panic(err)
		}
		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}
		fmt.Println("Comment id ke =", id)
	}
	//err = tx.Rollback() // <= Untuk insert ke database tetapi hanya meninggalkan urutan id di database dan tidak benar-benar masuk ke database
	err = tx.Commit() // <= untuk Commit atau memasukan sepenuhnya ke database
	if err != nil {
		panic(err)
	}
}
