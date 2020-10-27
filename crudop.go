package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// Student type
type Student struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func main() {

	// inserting a rows
	insert(Student{101, "RIT", "CSE"})
	insert(Student{102, "SAM", "CIVIL"})

	// updating the customer by id
	updateById(Student{101, "RITSAM", "CSE"})

	// select all customers
	results := selectAll()

	// iterating a results
	for results.Next() {
		var Student Student
		results.Scan(&student.Id, &student.Name)
		fmt.Println(student.Id, student.Name)
	}

	// select customer by id
	result := selectById(101)
	var student Student
	result.Scan(&student.Id, &student.Name)
	fmt.Println(student.Id, student.Name)

	// delete a customer by id
	delete(102)
}

// function to get a database connection
func connect() *sql.DB {
	db, err := sql.Open("mysql", "root:root123@/crud")
	if err != nil {
		fmt.Println("Error! Getting connection...")
	}
	return db
}

// function to insert a row in customer table
func insert(student Student) {
	db := connect()
	insert, err := db.Query("INSERT INTO crudop VALUES (?, ?)", student.Id, student.Name)
	if err != nil {
		fmt.Println("Error! Inserting records...")
	}
	defer insert.Close()
	defer db.Close()
}

// function to select all records from customer table
func selectAll() *sql.Rows {
	db := connect()
	results, err := db.Query("SELECT * FROM crudop")
	if err != nil {
		fmt.Println("Error! Getting records...")
	}
	defer db.Close()
	return results
}

// function to select a customer record from table by customer id
func selectById(id int) *sql.Row {
	db := connect()
	result := db.QueryRow("SELECT * FROM crudop WHERE id=?", id)
	defer db.Close()
	return result
}

// function to update a customer record by customer id
func updateById(student Student) {
	db := connect()
	db.QueryRow("UPDATE crudop SET name=? WHERE id=?", student.Name, student.Id)
}

// function to delete a customer by customer id
func delete(id int) {
	db := connect()
	db.QueryRow("DELETE FROM crudop WHERE id=?", id)
}
