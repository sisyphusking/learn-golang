package main

import (
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//参考文章：http://mindbowser.com/golang-go-with-gorm-2/
//官方文档：http://doc.gorm.io/crud.html#query

// type User struct {
// 	ID       int
// 	Username string
// }

type User struct {
	gorm.Model
	Name string
}

type UserModel struct {
	Id      int    `gorm:"primary_key";"AUTO_INCREMENT"` // 主键、自增
	Name    string `gorm:"size:255"`
	Address string `gorm:"type:varchar(100)"`
}

type Model struct {
	ID        int `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time // 这里为什么要引用指针？它不会删除记录只是将DeletedAt字段的值设置为当前时间，并且在查询时您将找不到记录，即我们称之为软删除
}

func main() {

	db, err := gorm.Open("mysql", "root:mysql1012!@tcp(127.0.0.1:3306)/db_apiserver?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	//在新建表名的时候名称不要用复数
	db.SingularTable(true)

	if err != nil {
		log.Println("connection failed to open")
	}
	log.Println("connnection established")

	//创建用户
	// db.DropTableIfExists(&User{})
	// db.CreateTable(&User{})
	// db.HasTable(&User{})

	// migrate
	db.Debug().DropTableIfExists(&UserModel{})
	db.Debug().AutoMigrate(&UserModel{})

	//创建记录
	user := &UserModel{Name: "xavieryin", Address: "shezhen"}
	//它的作用类似于：INSERT INTO `user_models` (`name`,`address`) VALUES ('xavieryin','shenzhen')
	db.Create(user)
	var users []UserModel = []UserModel{
		UserModel{Name: "xingwei", Address: "xinyang"},
		UserModel{Name: "xingwei", Address: "xinxian"},
	}
	for _, user := range users {
		//这里一定要注意，range中是值拷贝，create一定要用指针传值
		db.Create(&user)
	}
	//更新一条记录的信息
	db.Debug().Find(&user).Update("address", "shenzhen nanshan")

	//使用列名更新
	db.Debug().Model(&user).Update("Name", "yxw")

	//查询用户信息，然后更新
	db.Debug().Find(&user)
	user.Address = "tencent binhai"
	db.Debug().Save(&user)

	//原生sql，打印出dubug日志
	db.Debug().Exec("UPDATE user_model  SET name=? WHERE id = ?", "xavier", 1)

	//查询符合条件的两条数据
	db.Debug().Where("name = ? ", "xingwei").Find(&UserModel{})

	//and查询
	// db.Debug().Where("name = ? AND address = ?", "xavier", "tencent binhai").Find(&UserModel{})

	type Result struct {
		id      int
		name    string
		address string
	}

	// var results []Result
	//这里不能用Result，因为他没有gorm的标签，所以返回的结果集为空
	results := make([]*UserModel, 0)
	//获取结果集
	db.Debug().Raw("SELECT * from user_model").Find(&results)
	fmt.Println(results)
	for _, result := range results {

		fmt.Println(result.Id, result.Name, result.Address)
	}

}
