package main

import (
	"context"
	"fmt"
	"go-guide/libother/db/gorm/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var db *gorm.DB

func init() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second,   // 慢 SQL 阈值
			LogLevel:                  logger.Silent, // 日志级别
			IgnoreRecordNotFoundError: true,          // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  false,         // 禁用彩色打印
		},
	)

	dsn := "root:123456@tcp(127.0.0.1:3306)/orm?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true, // 禁用默认事务
		PrepareStmt:            true, // 缓存预编译
		Logger:                 newLogger,
	})
	if err != nil {
		fmt.Println("数据库连接错误:", err)
	}
}

func main() {
	//Select()
	//SelectWhere()
	//Select2()
	//Select1()

	//Create()
	//ScanExec()
	Preload()
	//HasOne()
}

func Select() {
	var st model.Student
	// 1.limit 1,order by id
	// res := db.Debug().First(&st, 1)
	// 2. limit 1,没有排序
	//res := db.Debug().Take(&st, 1)
	//res := db.Debug().Take(&st, "1")
	res := db.Debug().Take(&st, "id = ?", "1")
	// 使用 切片时，如果只有一个元素会自动转为 equal
	//res := db.Debug().Take(&st, []int{1, 2})
	if res.Error == nil {
		fmt.Printf("查询到了:%+v\n", st)
	} else {
		fmt.Println("记录不存在:", res.Error)
	}
}

func Select1() {
	// 如果Model没有primary key，则根据第一个字段排序
	result := map[string]interface{}{}
	res := db.Debug().Model(&model.Student{}).Take(&result, 1)
	if res.Error == nil {
		fmt.Printf("查询到了:%+v\n", result)
	} else {
		fmt.Println("记录不存在:", res.Error)
	}
}

func Select2() {
	var st = make([]model.Student, 0)
	// 获取全部记录, SELECT * FROM `student`
	res := db.Debug().Find(&st)
	if res.Error == nil {
		fmt.Printf("查询到了:%+v\n", st)
	} else {
		fmt.Println("记录不存在:", res.Error)
	}
}

func SelectWhere() {
	var st = make([]model.Student, 0)
	// 当使用结构体当做查询条件时，如果字段为零值，不会当做查询条件,SELECT * FROM `student` WHERE `student`.`name` = 'xulei'
	// res := db.Debug().Where(&model.Student{Name: "xulei", Age: 0}).Find(&st)
	//  SELECT * FROM `student` WHERE name = 'xulei' AND age = 0
	// res := db.Debug().Where("name = ? AND age = ?", "xulei", 0).Find(&st)
	// Not条件，!=, not in 可能不会用
	// SELECT * FROM `student` WHERE NOT name = 'xulei'
	// res := db.Debug().Not("name = ?", "xulei").Find(&st)
	// SELECT * FROM `student` WHERE `student`.`id` NOT IN (1,2)
	// res := db.Debug().Not([]int{1, 2}).Find(&st)
	// Or条件，有Where就有Or,SELECT * FROM `student` WHERE `student`.`id` = 1 OR `student`.`id` = 2
	// res := db.Debug().Where(1).Or(2).Find(&st)
	// Select()用来选取特定的字段
	// SELECT `id`,`name` FROM `student` WHERE `student`.`id` = 1
	res := db.Debug().Select("id", "name").Where(1).Find(&st)
	if res.Error == nil {
		fmt.Printf("查询到了:%+v\n", st)
	} else {
		fmt.Println("记录不存在:", res.Error)
	}
}

func Create() {
	st := model.Student{Name: "hah", Age: 18}
	fmt.Printf("创建前:%+v\n", st)
	// Omit可以忽略插入的字段
	//res := db.Debug().Omit("birthday").Create(&st)
	res := db.Debug().Create(&st)
	if res.Error == nil {
		fmt.Printf("查询到了:%+v\n", st)
	} else {
		fmt.Println("创建失败:", res.Error)
	}
}

func Update() {
	var st model.Student
	db.First(&st)
	st.Name = "徐哈哈"
	// Save 会保存所有的字段，即使字段是零值
	res := db.Save(&st)
	// Update用来更新指定column
	// res := db.Model(&st).Update("name", "hello")

	if res.Error == nil {
		fmt.Printf("查询到了:%+v\n", st)
	} else {
		fmt.Println("创建失败:", res.Error)
	}
}

func Delete() {
	var st model.Student
	res := db.Where("name = ?", "jinzhu").Delete(&st)
	// res := db.Delete(&model.Student{}, "10")
	if res.Error == nil {
		fmt.Printf("查询到了:%+v\n", st)
	} else {
		fmt.Println("创建失败:", res.Error)
	}
}

func ScanExec() {
	//var st model.Student
	//fmt.Println("st:", st)
	/*rows, _ := db.Raw("SELECT id, name, age FROM student WHERE name = ?", "xulei").Rows()
	defer rows.Close()
	for rows.Next() {
		var st model.Student
		rows.Scan(&st)

		// 业务逻辑...
	}*/

	// UPDATE student SET name = '你说呢' WHERE id IN (1,2,3)
	// res := db.Debug().Exec("UPDATE student SET name = ? WHERE id IN ?", "你说呢", []int64{1, 2, 3})
	// DryRun 模式,在不执行的情况下生成 SQL ，可以用于准备或测试生成的 SQL
	/*res := db.Debug().Session(&gorm.Session{DryRun: true}).Exec("UPDATE student SET name = ? WHERE id IN ?", "你说呢", []int64{1, 2, 3})
	if res.Error == nil {
		fmt.Printf("执行成功:%+v\n", st)
	} else {
		fmt.Println("执行失败:", res.Error)
	}*/
}

func Preload() {
	var st model.Student
	var sts []model.Student
	// Preload  SELECT * FROM `student` WHERE id = '1' LIMIT 1;SELECT * FROM `store` WHERE `store`.`id` = 1;
	// 会查询两次，如果StoreId是零值则不会查询
	//res := db.Debug().Preload("Store").Take(&st, "id = ?", "1")
	//res := db.Debug().Preload("Cars").Find(&sts)
	res := db.Debug().Preload("Cars").Where("id IN ?", []int{3, 4}).Find(&sts)
	if res.Error == nil {
		fmt.Printf("查询到了:%+v\n %+v\n", st, sts)
	} else {
		fmt.Println("记录不存在:", res.Error)
	}
}

func HasOne() {
	// Association 可以分步骤执行
	// Preload必须和student一块查询
	var st model.Student
	db.Debug().Where("id = ?", "4").Take(&st)
	// 使用关联，需要提前先查询好student
	err := db.Model(&st).Debug().Association("Car").Find(&st.Car)
	if err == nil {
		fmt.Printf("查询到了:\n%+v \n", st)
	} else {
		fmt.Println("记录不存在:", err)
	}
}

func Context() {
	ctx := context.Background()
	// 持续会话模式 或者单会话模式
	tx := db.WithContext(ctx)
	var st model.Student
	res := tx.Debug().Select("id", "name").Where(1).Find(&st)
	if res.Error == nil {
		fmt.Printf("查询到了:%+v\n", st)
	} else {
		fmt.Println("记录不存在:", res.Error)
	}
}
