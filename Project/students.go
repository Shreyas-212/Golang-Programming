package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Students struct {
	Id int 
	Name string 
	Email string 
	Password string
}

func (s *Students) Save() error {
	query := `
        INSERT INTO students (name, email, password)
        VALUES (?, ?, ?)`
    
    hashedPassword, err := hashPassword(s.Password)
    if err != nil {
        fmt.Printf("Error hashing password")
        return fmt.Errorf("failed to hash password")
    }
    result, err := db.Exec(query, s.Name, s.Email, hashedPassword)
    if err!= nil {
        fmt.Printf("Error executing query: %v\n", err)
        return fmt.Errorf("failed to insert student: %v", err)
    }

    id, err := result.LastInsertId()
    if err!= nil {
        fmt.Printf("Error getting last insert ID: %v\n", err)
        return fmt.Errorf("failed to get last insert ID: %w", err)
    }

    fmt.Printf("Student inserted successfully with id: %d\n", id)
    return nil
}

func getStudentNameByID(studentID int64) (string, error) {
	var studentName string
	query := "SELECT name FROM students WHERE id = ?"
	err := db.QueryRow(query, studentID).Scan(&studentName)
	if err != nil {
		return "", err
	}
	return studentName, nil
}

func getStudents(context *gin.Context) {
    students, err := getAllStudents()
    if err!= nil {
        context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch students"})
        return
    }
    context.JSON(200, gin.H{"students": students})
}

func getAllStudents() ([]Students, error) {
    query := "SELECT id, name, email, password FROM students"
    rows, err := db.Query(query)
    if err!= nil {
        return nil, err
    }
    defer rows.Close()

    var students []Students

    for rows.Next() {
        var student Students
        err := rows.Scan(&student.Id, &student.Name, &student.Email, &student.Password)
        if err!= nil {
            return nil, err
        }
        students = append(students, student)
    }

    return students, nil
}

func updateStudent(context *gin.Context) {
    studentid, err := strconv.ParseInt(context.Param("id"), 10, 64)
    if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid student ID"})
        return
    }

    var updatedStudent Students
    err = context.ShouldBindJSON(&updatedStudent)
    if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
        return
    }

    err = updateS(studentid, updatedStudent)
    if err!= nil {
        context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update student"})
        return
    }
    context.JSON(http.StatusOK, gin.H{"message": "Student updated successfully"})
}

func updateS(studentId int64, updatedStudent Students) error {
    query := "UPDATE students SET name = ?, email = ? WHERE id = ?"
    _, err := db.Exec(query, updatedStudent.Name, updatedStudent.Email, studentId)
    if err != nil {
        fmt.Printf("Error executing query: %v\n", err)
        return fmt.Errorf("failed to update student: %w", err)
    }
    return nil
}

func deleteStudent(context *gin.Context) {
    studentid, err := strconv.ParseInt(context.Param("id"), 10, 64)
    if err!= nil {
        context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid student ID"})
        return
    }

    err = deleteS(studentid)
    if err!= nil {
        context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete student"})
        return
    }
    context.JSON(http.StatusOK, gin.H{"message": "Student deleted successfully"})
}

func deleteS(studentId int64) error {
    _, err := db.Exec("DELETE FROM registrations WHERE student_id = ?", studentId)
    if err != nil {
        fmt.Printf("Error deleting student registrations: %v\n", err)
        return fmt.Errorf("failed to delete student registrations: %v", err)
    }

    query := "DELETE FROM students WHERE id = ?"
    _, err = db.Exec(query, studentId)
    if err != nil {
        fmt.Printf("Error executing query to delete student: %v\n", err)
        return fmt.Errorf("failed to delete student: %v", err)
    }

    return nil
}

func getStudent(context *gin.Context) {
    studentid, err := strconv.ParseInt(context.Param("id"), 10, 64)
    if err!= nil {
        context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid student ID"})
        return
    }
    student, err := getStudentByID(studentid)
    if err!= nil {
        context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch student"})
        return
    }
    context.JSON(http.StatusOK, gin.H{"student": student})
}

func getStudentByID(studentID int64) (*Students, error) {
    query := "SELECT id, name, email FROM students WHERE id = ?"
    var student Students
    row := db.QueryRow(query, studentID)

    err := row.Scan(&student.Id, &student.Name, &student.Email)
    if err != nil {
        return nil, err
    }

    return &student, nil
}
