package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (c Course) Save() error {
	query := `
        INSERT INTO courses (name, course_no, duration, student_id)
        VALUES (?, ?, ?, ?)`

	result, err := db.Exec(query, c.Name, c.Course_no, c.Duration, c.StudentId)
	if err != nil {
		fmt.Printf("Error executing query: %v\n", err)
		return fmt.Errorf("failed to insert course: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		fmt.Printf("Error getting last insert ID: %v\n", err)
		return fmt.Errorf("failed to get last insert ID: %w", err)
	}

	fmt.Printf("Course inserted successfully with id: %d\n", id)
	return nil
}

func getCourses(c *gin.Context) {
	courses, err := getAllCourses()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch courses"})
		return
	}
	c.JSON(200, gin.H{"courses": courses})
}

func getAllCourses() ([]Course, error) {
	query := "SELECT * FROM courses"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []Course

	for rows.Next() {
		var course Course
		err := rows.Scan(&course.Id, &course.Name, &course.Course_no, &course.Duration, &course.StudentId)

		if err != nil {
			return nil, err
		}

		courses = append(courses, course)
	}

	return courses, nil
}

func getCourse(context *gin.Context) {
	courseid := context.Param("id")
	parsedCourseId, err := strconv.ParseInt(courseid, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid course ID format."})
		return
	}

	course, err := getCourseById(parsedCourseId)
	if err != nil {
		if err.Error() == "course not found" {
			context.JSON(http.StatusNotFound, gin.H{"message": "Course not found"})
			return
		}
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch course."})
		return
	}

	context.JSON(http.StatusOK, course)
}

func getCourseById(courseId int64) (*Course, error) {
	query := "SELECT * FROM courses WHERE id = ?"
	row := db.QueryRow(query, courseId)

	var course Course
	err := row.Scan(&course.Id, &course.Name, &course.Course_no, &course.Duration, &course.StudentId)
	if err != nil {
		return nil, err
	}
	return nil, fmt.Errorf("error scanning row: %v", err)
}

func createCourse(context *gin.Context) {
	var course Course
	err := context.ShouldBindJSON(&course)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	err = course.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create course. Try again later."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Course created!", "course": course})
}

func updateCourse(context *gin.Context) {
	courseId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid course ID format"})
		return
	}

	var updatedCourse Course
	if err := context.ShouldBindJSON(&updatedCourse); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid course data"})
		return
	}

	err = updateC(courseId, updatedCourse)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Course updated successfully"})
}

func updateC(courseId int64, updatedCourse Course) error {
	query := "UPDATE courses SET name=?, course_no=?, duration=?, student_id=? WHERE id=?"
	_, err := db.Exec(query, updatedCourse.Name, updatedCourse.Course_no, updatedCourse.Duration, updatedCourse.StudentId, courseId)

	if err != nil {
		return fmt.Errorf("error updating course: %v", err)
	}
	fmt.Printf("Course Updated successfully with id: %d", updatedCourse.Id)
	return nil
}

func deleteCourse(context *gin.Context) {
	courseId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid course ID format"})
		return
	}

	err = delete(courseId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Course deleted successfully"})
}

func delete(courseid int64) error {
	query := "DELETE FROM courses WHERE id=?"
	_, err := db.Exec(query, courseid)
	if err != nil {
		fmt.Printf("Error deleting course: %v\n", err)
	}
	return nil
}

func getCourseNameByID(courseID int64) (string, error) {
	var courseName string
	query := "SELECT name FROM courses WHERE id = ?"
	err := db.QueryRow(query, courseID).Scan(&courseName)
	if err != nil {
		return "", err
	}
	return courseName, nil
}
