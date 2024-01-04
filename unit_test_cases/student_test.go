package repository

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestInsertOrUpdateStudentConfig(t *testing.T) {
	// Create a mocked database connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock database: %v", err)
	}
	defer db.Close()

	// Initialize your StudentConfig for testing
	student := StudentConfig{
		StudentID:   1,
		FirstName:   "John",
		LastName:    "Doe",
		DateOfBirth: "1999-01-01",
		Gender:      "Male",
		Email:       "john@example.com",
		PhoneNumber: "1234567890",
	}

	// Create expected database query and mock behavior for InsertOrUpdateStudentConfig
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(insertOrUpdateStudentConfig)).WithArgs(student.StudentID, student.FirstName, student.LastName, student.DateOfBirth, student.Gender, student.Email, student.PhoneNumber).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()
	mock.ExpectQuery(regexp.QuoteMeta(getStudentConfig)).WithArgs(student.StudentID).
		WillReturnRows(sqlmock.NewRows([]string{"student_id", "first_name", "last_name", "date_of_birth", "gender", "email", "phone_number"}).
			AddRow(student.StudentID, student.FirstName, student.LastName, student.DateOfBirth, student.Gender, student.Email, student.PhoneNumber))

	// Test InsertOrUpdateStudentConfig function
	insertedConfig, err := InsertOrUpdateStudentConfig(db, student)
	if err != nil {
		t.Fatalf("InsertOrUpdateStudentConfig failed: %v", err)
	}

	// Add assertions to check the correctness of insertedConfig
	if insertedConfig.StudentID != student.StudentID {
		t.Errorf("Expected StudentID: %d, Got: %d", student.StudentID, insertedConfig.StudentID)
	}
	if insertedConfig.FirstName != student.FirstName {
		t.Errorf("Expected FirstName: %s, Got: %s", student.FirstName, insertedConfig.FirstName)
	}
	// Similarly, add assertions for other fields if needed

	// Ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("mock expectations were not met: %v", err)
	}
}

func TestGetStudentConfig(t *testing.T) {
	// Create a mocked database connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock database: %v", err)
	}
	defer db.Close()

	// Initialize a StudentConfig for testing
	studentID := 1
	expectedStudent := StudentConfig{
		StudentID:   studentID,
		FirstName:   "John",
		LastName:    "Doe",
		DateOfBirth: "1999-01-01",
		Gender:      "Male",
		Email:       "john@example.com",
		PhoneNumber: "1234567890",
	}

	// Define the expected SELECT query and mock behavior for GetStudentConfig
	mock.ExpectQuery(regexp.QuoteMeta(getStudentConfig)).WithArgs(studentID).
		WillReturnRows(sqlmock.NewRows([]string{"student_id", "first_name", "last_name", "date_of_birth", "gender", "email", "phone_number"}).
			AddRow(expectedStudent.StudentID, expectedStudent.FirstName, expectedStudent.LastName, expectedStudent.DateOfBirth, expectedStudent.Gender, expectedStudent.Email, expectedStudent.PhoneNumber))

	// Test GetStudentConfig function
	retrievedStudent, err := GetStudentConfig(db, studentID)
	if err != nil {
		t.Fatalf("GetStudentConfig failed: %v", err)
	}

	// Add assertions to check the correctness of retrievedStudent
	if retrievedStudent.StudentID != expectedStudent.StudentID {
		t.Errorf("Expected StudentID: %d, Got: %d", expectedStudent.StudentID, retrievedStudent.StudentID)
	}
	if retrievedStudent.FirstName != expectedStudent.FirstName {
		t.Errorf("Expected FirstName: %s, Got: %s", expectedStudent.FirstName, retrievedStudent.FirstName)
	}
	// Similarly, add assertions for other fields if needed

	// Ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("mock expectations were not met: %v", err)
	}
}
