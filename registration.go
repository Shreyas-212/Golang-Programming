package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var student Students
	if err := context.ShouldBindJSON(&student); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid student data"})
		return
	}

	err := student.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create student. Try again later."})
		return
	}

	token, err := generateToken(student.Email, int64(student.Id))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to generate token"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Student created!", "student": student, "token": token})
}

func studentExists(studentId int64) (bool, error) {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM students WHERE id = ?)"
	err := db.QueryRow(query, studentId).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("failed to check if student exists: %v", err)
	}
	return exists, nil
}

func login(context *gin.Context) {
	var student Students
	err := context.ShouldBindJSON(&student)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid student data"})
		return
	}
	if err := student.authenticateLogin(); err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
		return
	}

	token, err := generateToken(student.Email, int64(student.Id))
	context.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})
}

func (s *Students) authenticateLogin() error {
    query := "SELECT id, password FROM students WHERE email = ?"
    var dbPassword string

    log.Printf("Authenticating email: %s", s.Email)

    row := db.QueryRow(query, s.Email)
    err := row.Scan(&s.Id, &dbPassword)
    if err != nil {
        if err == sql.ErrNoRows {
            log.Printf("No student found with email: %s", s.Email)
            return fmt.Errorf("Invalid credentials")
        }
        log.Printf("Database error: %v", err)
        return fmt.Errorf("Internal server error")
    }

    if !compareHashAndPassword(dbPassword, s.Password) {
        log.Printf("Password mismatch for email: %s", s.Email)
        return fmt.Errorf("Invalid credentials")
    }

    log.Printf("Authentication successful for email: %s", s.Email)
    return nil
}

func registerForCourse(context *gin.Context) {
	var registration struct {
		CourseID  int64
		StudentID int64
	}

	if err := context.ShouldBindJSON(&registration); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input data"})
		return
	}

	if registration.CourseID == 0 || registration.StudentID == 0 {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Both course_id and student_id are required"})
		return
	}

	exists, err := studentExists(registration.StudentID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error verifying student existence"})
		return
	}
	if !exists {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Student does not exist"})
		return
	}
	studentName, err := getStudentNameByID(registration.StudentID)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Student does not exist"})
		return
	}
	courseName, err := getCourseNameByID(registration.CourseID)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Course does not exist"})
		return
	}

	err = Register(registration.CourseID, registration.StudentID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("Failed to register student for course: %v", err)})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Course registered successfully",
		"student": studentName,
		"course":  courseName})
}

func Register(courseId int64, studentId int64) error {
	var courseName string
	err := db.QueryRow("SELECT name FROM courses WHERE id = ?", courseId).Scan(&courseName)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("course with ID %d not found", courseId)
		}
		return fmt.Errorf("error fetching course details: %w", err)
	}

	var studentName string
	err = db.QueryRow("SELECT name FROM students WHERE id = ?", studentId).Scan(&studentName)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("student with ID %d not found", studentId)
		}
		return fmt.Errorf("error fetching student details: %w", err)
	}

	query := `INSERT INTO registrations (course_id, course_name, student_id, student_name) VALUES (?, ?, ?, ?)`
	_, err = db.Exec(query, courseId, courseName, studentId, studentName)
	if err != nil {
		return fmt.Errorf("failed to register student for course: %w", err)
	}

	fmt.Printf("Student %s registered for course %s (ID: %d)\n", studentName, courseName, courseId)
	return nil
}

func cancelRegisterForCourse(context *gin.Context) {
	var registration struct {
		CourseID  int64
		StudentID int64
	}

	err := context.ShouldBindJSON(&registration)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input data"})
		return
	}
	if registration.CourseID == 0 || registration.StudentID == 0 {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Both course_id and student_id are required"})
		return
	}
	err = CancelRegistration(registration.CourseID, registration.StudentID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("Failed to cancel registration: %v", err)})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Registration cancelled successfully"})
}

func CancelRegistration(courseid, studentid int64) error {
	query := "DELETE FROM registrations WHERE course_id=? AND student_id=?"
	_, err := db.Exec(query, courseid, studentid)
	if err != nil {
		fmt.Printf("Error deleting registration: %v\n", err)
	}
	return nil
}