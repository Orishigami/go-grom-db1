// main.go
package main

import (
	"log"      // ใช้สำหรับแสดงข้อความ error ออกทางหน้าจอ
	"net/http" // ใช้สำหรับสร้าง web server
	"os"       // ใช้สำหรับอ่านค่า environment variable
	"time"

	"github.com/gin-contrib/cors"

	"github.com/Orishigami/go-grom-db1/db"     // นำเข้า db
	"github.com/Orishigami/go-grom-db1/models" // นำเข้า models
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv" // ใช้สำหรับอ่านค่าจากไฟล์ .env
)

func main() {
	// อ่านค่า environment variable จากไฟล์ .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// อ่านค่า environment variable ที่ต้องการใช้งาน
	dbType := os.Getenv("DB_TYPE")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	// เชื่อมต่อฐานข้อมูล
	database, err := db.ConnectDatabase(dbType, dbUser, dbPassword, dbHost, dbPort, dbName)
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// สร้างตารางในฐานข้อมูล
	err = database.AutoMigrate(&models.Item{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
	err = database.AutoMigrate(&models.Student{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
	err = database.AutoMigrate(&models.Subject{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	// สร้างตัวแปร itemRepo เพื่อเรียกใช้งาน ItemRepository
	itemRepo := models.NewItemRepository(database)
	studentRepo := models.NewStudentRepository(database)
	subjectRepo := models.NewSubjectRepository(database)

	r := gin.Default()

	// กำหนด cors (Cross-Origin Resource Sharing)
	r.Use(cors.New(cors.Config{
		// 3000 คือ port ที่ใช้งานใน frontend react
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// api /items จะเป็นการเรียกใช้งานฟังก์ชัน GetItems ใน ItemRepository
	r.GET("/items", itemRepo.GetItems)
	r.GET("/students", studentRepo.GetStudents)
	r.GET("/subjects", subjectRepo.GetSubjects)

	// api /items/:id จะเป็นการเรียกใช้งานฟังก์ชัน GetItem ใน ItemRepository
	r.POST("/items", itemRepo.PostItem)
	r.POST("/students", studentRepo.PostStudent)
	r.POST("/subjects", subjectRepo.PostSubject)

	// api /items/:id จะเป็นการเรียกใช้งานฟังก์ชัน GetItem ใน ItemRepository
	// /items/1 จะเป็นการส่งค่า id ที่เป็นตัวเลข 1 ไปยังฟังก์ชัน GetItem ใน ItemRepository
	r.GET("/items/:id", itemRepo.GetItem)
	r.GET("/students/:id", studentRepo.GetStudent)
	r.GET("/subjects/:id", subjectRepo.GetSubject)

	// api /items/:id จะเป็นการเรียกใช้งานฟังก์ชัน UpdateItem ใน ItemRepository
	r.PUT("/items/:id", itemRepo.UpdateItem)
	r.PUT("/students/:id", studentRepo.UpdateStudent)
	r.PUT("/subjects/:id", subjectRepo.UpdateSubject)

	// api /items/:id จะเป็นการเรียกใช้งานฟังก์ชัน DeleteItem ใน ItemRepository
	r.DELETE("/items/:id", itemRepo.DeleteItem)
	r.DELETE("/students/:id", studentRepo.DeleteStudent)
	r.DELETE("/subjects/:id", subjectRepo.DeleteSubject)

	// ถ้าไม่มี api ที่ตรงกับที่กำหนด จะแสดงข้อความ Not found
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not found"})
	})

	// Run the server
	if err := r.Run(":5000"); err != nil {
		log.Fatalf("Server is not running: %v", err)
	}
}
