package repository

import (
	"database/sql"
	"fmt"
)

type StudentConfig struct {
	StudentID   int
	FirstName   string
	LastName    string
	DateOfBirth string
	Gender      string
	Email       string
	PhoneNumber string
}

const (
	insertOrUpdateStudentConfig = `
	INSERT INTO student (student_id, first_name, last_name, date_of_birth, gender, email, phone_number)
	VALUES ($1, $2, $3, $4, $5, $6, $7)
	ON CONFLICT (student_id) DO UPDATE
	SET
		first_name = $2,
		last_name = $3,
		date_of_birth = $4,
		gender = $5,
		email = $6,
		phone_number = $7;
`

	getStudentConfig = `
SELECT student_id, first_name, last_name, date_of_birth, gender, email, phone_number
FROM student
WHERE student_id = $1
`
)

func InsertOrUpdateStudentConfig(db *sql.DB, student StudentConfig) (StudentConfig, error) {
	tx, err := db.Begin()
	if err != nil {
		return StudentConfig{}, fmt.Errorf("could not begin transaction: %v", err)
	}

	_, err = tx.Exec(insertOrUpdateStudentConfig,
		student.StudentID,
		student.FirstName,
		student.LastName,
		student.DateOfBirth,
		student.Gender,
		student.Email,
		student.PhoneNumber,
	)

	if err != nil {
		tx.Rollback()
		return StudentConfig{}, fmt.Errorf("error executing query: %v", err)
	}

	err = tx.Commit()
	if err != nil {
		return StudentConfig{}, fmt.Errorf("error committing transaction: %v", err)
	}

	insertedConfig, err := GetStudentConfig(db, student.StudentID)
	if err != nil {
		return StudentConfig{}, fmt.Errorf("error fetching inserted student configuration: %v", err)
	}

	return insertedConfig, nil
}

func GetStudentConfig(db *sql.DB, studentID int) (StudentConfig, error) {

	var student StudentConfig

	row := db.QueryRow(getStudentConfig, studentID)
	err := row.Scan(
		&student.StudentID,
		&student.FirstName,
		&student.LastName,
		&student.DateOfBirth,
		&student.Gender,
		&student.Email,
		&student.PhoneNumber,
	)
	if err != nil {
		return StudentConfig{}, fmt.Errorf("error fetching student configuration: %v", err)
	}

	return student, nil
}
