// models/SubjectRepository.go
package models

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// สร้าง struct ชื่อ SubjectRepository ที่มีฟิลด์ชื่อ Db เป็น pointer ของ gorm.DB
type SubjectRepository struct {
	Db *gorm.DB
}

// ทำหน้าที่สร้าง Instance ของ SubjectRepository และส่งคืนกลับไป ให้เรียกใช้งานได้
func NewSubjectRepository(db *gorm.DB) *SubjectRepository {
	return &SubjectRepository{Db: db}
}

// ทำหน้าที่ดึงข้อมูล Subject ทั้งหมดจากฐานข้อมูล และส่งกลับไปให้ผู้ใช้งาน
// (r *SubjectRepository) คือการระบุว่าเป็น method ของ SubjectRepository
// r คือตัวแปรที่เป็น pointer ของ SubjectRepository
// ฟังก์ชัน GetSubjects รับพารามิเตอร์เป็น c *gin.Context และมีการส่งกลับค่าเป็น JSON กลับไปให้ผู้ใช้งานผ่าน c.JSON(200, subjects) ซึ่ง c คือตัวแปรที่เป็น pointer ของ gin.Context
// ในกรณีนี้เราไม่ต้องใช้ return เพราะเราใช้ c.JSON แทน
// ฟังก์ชันนี้ parameter ที่รับมาจะเป็น pointer ของ gin.Context เพราะเราจะใช้ c.JSON ส่งค่ากลับไปให้ผู้ใช้งาน
func (r *SubjectRepository) GetSubjects(c *gin.Context) {
	var subjects []Subject
	// ดึงข้อมูล Subject ทั้งหมดจากฐานข้อมูล และเก็บลงในตัวแปร subjects
	r.Db.Find(&subjects)  // SELECT * FROM subjects
	c.JSON(200, subjects) // ส่งข้อมูลกลับไปให้ผู้ใช้งาน
}

// ทำหน้าที่เพิ่มข้อมูล Subject ลงในฐานข้อมูล และส่งกลับไปให้ผู้ใช้งานผ่าน c.JSON(200, newSubject)
func (r *SubjectRepository) PostSubject(c *gin.Context) {
	var newSubject Subject
	c.BindJSON(&newSubject)  // รับค่า JSON จากผู้ใช้งาน และแปลงเป็น struct ของ Item
	r.Db.Create(&newSubject) // INSERT INTO items (name, price) VALUES (newItem.Name, newItem.Price)
	c.JSON(200, newSubject)  //ส่งข้อมูลกลับไปให้ผู้ใช้งาน
}

// ฟังก์ชันค้นหา Subject จากฐานข้อมูล โดยใช้ id ที่รับเข้ามาเป็นเงื่อนไขในการค้นหา
// ฟังก์ชันนี้ parameter ที่รับมาจะเป็น pointer ของ gin.Context เพราะเราจะใช้ c.JSON ส่งค่ากลับไปให้ผู้ใช้งาน
func (r *SubjectRepository) GetSubject(c *gin.Context) {
	id := c.Param("id")      // รับค่า id จากผู้ใช้งาน
	var subject Subject      // สร้างตัวแปร subject เพื่อเก็บข้อมูลที่ค้นหาได้
	r.Db.First(&subject, id) // SELECT * FROM subjects WHERE id = id
	c.JSON(200, subject)     // ส่งข้อมูลกลับไปให้ผู้ใช้งาน
}

// ฟังก์ชันอัพเดทข้อมูล Subject ลงในฐานข้อมูล และส่งกลับไปให้ผู้ใช้งานผ่าน c.JSON(200, subject)
func (r *SubjectRepository) UpdateSubject(c *gin.Context) {
	id := c.Param("id")
	var subject Subject
	r.Db.First(&subject, id) // SELECT * FROM subjects WHERE id = id
	c.BindJSON(&subject)
	r.Db.Save(&subject) // UPDATE subjects SET name = subject.Name, price = subject.Price WHERE id = id
	c.JSON(200, subject)
}

// ฟังก์ชันลบข้อมูล Subject ออกจากฐานข้อมูล และส่งกลับไปให้ผู้ใช้งานผ่าน c.JSON(200, gin.H{"id" + id: "is deleted"})
// gin.H ทำหน้าที่สร้าง map ของ key และ value และส่งกลับไปให้ผู้ใช้งาน
// H คือตัวย่อของ Header ใน HTTP
func (r *SubjectRepository) DeleteSubject(c *gin.Context) {
	id := c.Param("id")
	var subject Subject
	r.Db.Delete(&subject, id) // DELETE FROM subjects WHERE id = id
	c.JSON(200, gin.H{"id" + id: "is deleted"})
}
