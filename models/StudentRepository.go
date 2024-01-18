// models/StudentRepository.go
package models

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// สร้าง struct ชื่อ StudentRepository ที่มีฟิลด์ชื่อ Db เป็น pointer ของ gorm.DB
type StudentRepository struct {
	Db *gorm.DB
}

// ทำหน้าที่สร้าง Instance ของ StudentRepository และส่งคืนกลับไป ให้เรียกใช้งานได้
func NewStudentRepository(db *gorm.DB) *StudentRepository {
	return &StudentRepository{Db: db}
}

// ทำหน้าที่ดึงข้อมูล Student ทั้งหมดจากฐานข้อมูล และส่งกลับไปให้ผู้ใช้งาน
// (r *StudentRepository) คือการระบุว่าเป็น method ของ StudentRepository
// r คือตัวแปรที่เป็น pointer ของ StudentRepository
// ฟังก์ชัน GetStudents รับพารามิเตอร์เป็น c *gin.Context และมีการส่งกลับค่าเป็น JSON กลับไปให้ผู้ใช้งานผ่าน c.JSON(200, students) ซึ่ง c คือตัวแปรที่เป็น pointer ของ gin.Context
// ในกรณีนี้เราไม่ต้องใช้ return เพราะเราใช้ c.JSON แทน
// ฟังก์ชันนี้ parameter ที่รับมาจะเป็น pointer ของ gin.Context เพราะเราจะใช้ c.JSON ส่งค่ากลับไปให้ผู้ใช้งาน
func (r *StudentRepository) GetStudents(c *gin.Context) {
	var students []Student
	// ดึงข้อมูล Student ทั้งหมดจากฐานข้อมูล และเก็บลงในตัวแปร students
	r.Db.Find(&students)  // SELECT * FROM students
	c.JSON(200, students) // ส่งข้อมูลกลับไปให้ผู้ใช้งาน
}

// ทำหน้าที่เพิ่มข้อมูล Student ลงในฐานข้อมูล และส่งกลับไปให้ผู้ใช้งานผ่าน c.JSON(200, newStudent)
func (r *StudentRepository) PostStudent(c *gin.Context) {
	var newStudent Student
	c.BindJSON(&newStudent)  // รับค่า JSON จากผู้ใช้งาน และแปลงเป็น struct ของ Student
	r.Db.Create(&newStudent) // INSERT INTO students (name, price) VALUES (newStudent.Name, newStudent.Price)
	c.JSON(200, newStudent)  // ส่งข้อมูลกลับไปให้ผู้ใช้งาน
}

// ฟังก์ชันค้นหา Student จากฐานข้อมูล โดยใช้ id ที่รับเข้ามาเป็นเงื่อนไขในการค้นหา
// ฟังก์ชันนี้ parameter ที่รับมาจะเป็น pointer ของ gin.Context เพราะเราจะใช้ c.JSON ส่งค่ากลับไปให้ผู้ใช้งาน
func (r *StudentRepository) GetStudent(c *gin.Context) {
	id := c.Param("id")      // รับค่า id จากผู้ใช้งาน
	var student Student      // สร้างตัวแปร student เพื่อเก็บข้อมูลที่ค้นหาได้
	r.Db.First(&student, id) // SELECT * FROM students WHERE id = id
	c.JSON(200, student)     // ส่งข้อมูลกลับไปให้ผู้ใช้งาน
}

// ฟังก์ชันอัพเดทข้อมูล Student ลงในฐานข้อมูล และส่งกลับไปให้ผู้ใช้งานผ่าน c.JSON(200, student)
func (r *StudentRepository) UpdateStudent(c *gin.Context) {
	id := c.Param("id")
	var student Student
	r.Db.First(&student, id) // SELECT * FROM students WHERE id = id
	c.BindJSON(&student)
	r.Db.Save(&student) // UPDATE students SET name = student.Name, price = student.Price WHERE id = id
	c.JSON(200, student)
}

// ฟังก์ชันลบข้อมูล Student ออกจากฐานข้อมูล และส่งกลับไปให้ผู้ใช้งานผ่าน c.JSON(200, gin.H{"id" + id: "is deleted"})
// gin.H ทำหน้าที่สร้าง map ของ key และ value และส่งกลับไปให้ผู้ใช้งาน
// H คือตัวย่อของ Header ใน HTTP
func (r *StudentRepository) DeleteStudent(c *gin.Context) {
	id := c.Param("id")
	var student Student
	r.Db.Delete(&student, id) // DELETE FROM students WHERE id = id
	c.JSON(200, gin.H{"id" + id: "is deleted"})
}
