package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" // "_"代码不直接使用包，底层连接要使用！
	"github.com/jinzhu/gorm"
)

// 创建全局的连接池的句柄
var GlobalConn *gorm.DB

func main() {
	// 连接数据库 ——格式： 用户名：密码@协议(ip:port)/数据库名
	conn, err := gorm.Open("mysql", "root:12345678@tcp(127.0.0.1:3306)/ihome?parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("gorm.Open err:", err)
		return
	}
	//defer conn.Close()

	GlobalConn = conn

	// 设置空闲时的连接个数10个
	GlobalConn.DB().SetMaxIdleConns(10)
	GlobalConn.DB().SetMaxOpenConns(100)

	//// 借助gorm创建数据库表
	if GlobalConn.AutoMigrate(new(Student)).Error != nil {
		fmt.Println("GlobalConn.AutoMigrate(new(Student)) Error")
		return
	}

	// 插入数据
	InsertData()
	// 查询数据
	//SearchData()
	UpdateData()
}

// 创建全局结构体
type Student struct {
	gorm.Model // go语言中的匿名结构体成员  ———— 继承
	// string 默认大小255 可以修改size大小
	Name  string `gorm:"size:100;default:'xiaoming'"`
	Age   int
	Class int `gorm:"not null"`
}

func InsertData() {
	// 创建数据
	student := Student{
		Name: "yc",
		Age:  18,
	}
	// 插入（创建）数据
	GlobalConn.Create(&student)

}

func SearchData() {
	var stu []Student

	//GlobalConn.First(&stu)
	//GlobalConn.Select("name, age").First(&stu)
	//GlobalConn.Select("name,age").Find(&stu)

	// where 的使用
	// 查询姓名为yc的name和age
	//GlobalConn.Select("name, age").Where("name = ?", "yc").Find(&stu)
	//GlobalConn.Select("name, age").Where("name = ?", "yc").
	//	Where("age = ?", "18").Find(&stu)
	//GlobalConn.Select("name, age").Where("name = ? and age = ?", "yc", 18).Find(&stu)
	fmt.Println(stu)

}

func UpdateData() {
	var stu Student

	stu.Name = "zhangsan"
	stu.Age = 100
	//fmt.Println(GlobalConn.Model(new(Student)).Where("name = ?", "yc").Update("name","zhangsan").Error)
	fmt.Println(GlobalConn.Model(new(Student)).Where("name = ?", "yc").
		Updates(map[string]interface{}{"name": "tyc", "age": 77}).Error)
}

func DeleteData() {

}
