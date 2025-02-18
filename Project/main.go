package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Course struct {
	Id        int
	Name      string
	Course_no int
	Duration  int
	StudentId int64
}

var db *sql.DB

func main() {
	_, err := InitDB()

	if err != nil {
		log.Fatal("err:", err)
	}
	router := gin.Default()
	userGroup := router.Group("/user")

    userGroup.GET("/course/list", getCourses)
    userGroup.GET("/course/:id", getCourse)

	authenticated := router.Group("/user")
	authenticated.Use(Authentication)
	authenticated.POST("/course", createCourse)
	authenticated.PUT("/course/:id", updateCourse)
	authenticated.DELETE("/course/:id", deleteCourse)
	authenticated.GET("/student/list", getStudents)
	authenticated.GET("/student/:id", getStudent)
	authenticated.PUT("/student/:id", updateStudent)
	authenticated.DELETE("/student/:id", deleteStudent)
	authenticated.POST("/course/:id/register", registerForCourse)
	authenticated.DELETE("/course/:id/register", cancelRegisterForCourse)

    userGroup.POST("student/signup", signup)
	userGroup.POST("/login", login)
	
	router.Run(":8080") // localhost:8080
}

func InitDB() (*sql.DB, error) {
	connstr := "root:shre202124@@tcp(localhost:3306)/Project"

	var err error

	db, err = sql.Open("mysql", connstr)
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Cannot reach the database:", err)
		return nil, err
	}

	fmt.Println("Connected to MySQL successfully!")

	createTables()
	return db, nil
}

func createTables() {
	var createStudentsTable = `
	CREATE TABLE IF NOT EXISTS students (
        id INT UNSIGNED NOT NULL AUTO_INCREMENT,
        name TEXT NOT NULL,
        email TEXT NOT NULL,
        password TEXT NOT NULL,
		PRIMARY KEY (id)
	);`
	_, err := db.Exec(createStudentsTable)
	if err != nil {
		panic("Error creating students table: " + err.Error())
	}

	fmt.Println("Students table created successfully!")

	var createCoursesTable = `
	CREATE TABLE IF NOT EXISTS courses (
        id INT UNSIGNED NOT NULL AUTO_INCREMENT,
        name VARCHAR(255) NOT NULL,
        course_no INT NOT NULL,
        duration INT NOT NULL,
        student_id INT UNSIGNED NOT NULL,
		PRIMARY KEY (id), 
		FOREIGN KEY (student_id) REFERENCES students(id) 
		ON DELETE CASCADE
		ON UPDATE CASCADE
    );`
	_, err = db.Exec(createCoursesTable)
	if err != nil {
		log.Fatal("Error creating courses table: " + err.Error())
	}

	fmt.Println("Courses table created successfully.")

	var createRegistrationsTable = `
	CREATE TABLE IF NOT EXISTS registrations (
    id INT UNSIGNED NOT NULL AUTO_INCREMENT,
    course_id INT UNSIGNED NOT NULL,
    student_id INT UNSIGNED NOT NULL,
    course_name VARCHAR(255) NOT NULL,
    student_name VARCHAR(255) NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (course_id) REFERENCES courses(id)
        ON DELETE CASCADE
        ON UPDATE CASCADE,
    FOREIGN KEY (student_id) REFERENCES students(id)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);`
	_, err = db.Exec(createRegistrationsTable)
	if err != nil {
		log.Fatal("Error creating registrations table: " + err.Error())
	}

	fmt.Println("Registrations table created successfully.")
}
