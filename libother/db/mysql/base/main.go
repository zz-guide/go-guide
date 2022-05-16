package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/tal-tech/go-zero/core/timex"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func init() {
	db, err = sql.Open("mysql", "root:xl123456?@tcp(47.105.50.31:3306)/casbin?charset=utf8mb4&parseTime=True")
	if err != nil {
		log.Println("open mysql failed,", err)
		return
	}

	db.SetConnMaxLifetime(100 * time.Second) //最大连接周期，超过时间的连接就close
	db.SetMaxOpenConns(100)                  //设置最大连接数
	db.SetMaxIdleConns(16)                   //设置闲置连接数

	// 尝试与数据库建立连接（校验dsn是否正确）
	err = db.Ping()
	if err != nil {
		log.Println("ping mysql failed,", err)
		return
	}

	//defer func() {
	//	_ = database.Close()
	//}()
}

type User struct {
	Id   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

/**
	1.查询使用query系列
	2.修改试用exec系列
	3.事务使用tx
	4.sql统一使用预处理

 问题：
	1.打印真实执行的sql
	2.sql执行时长
	3.sql->struct或者sql->map转换
	4.sql字段类型映射，例如id,bigint,时间类型
	5.建议使用QueryContext方法封装所有查询，QueryRow系列方法无法获取原生sql
	6.QueryRow系列方法只适用于明确知道返回结果只有一条记录的时候使用
*/

func main() {
	_query()
	//_queryContext()
	//_queryRow()
	//_queryRowContext()
	//_execInsert()
	//_execUpdate()
	//_execDelete()
	//_prepare()
	//_transaction()
}

const defaultSlowThreshold = time.Millisecond * 0

func ReprOfDuration(duration time.Duration) string {
	return fmt.Sprintf("%.1fms", float32(duration)/float32(time.Millisecond))
}

func _query() {
	// 获取一行或者多行数据
	var user User
	sqlStr := "select `id`,`name` from `user`"
	// 1.表不存在的话，rows是nil，
	// 2.如果不存在数据的话，err也是nil

	startTime := timex.Now() // 开始时间
	// 执行一次select,结果可能是多行
	rows, err := db.Query(sqlStr)
	duration := timex.Since(startTime)
	if duration > defaultSlowThreshold {
		log.Printf("[SQL] exec: slowcall duration=%s - %s\n", ReprOfDuration(duration), sqlStr)
	} else {
		log.Printf("sql exec: %s\n", sqlStr)
	}

	if err != nil {
		log.Println("查询失败, ", err, rows)
		return
	}

	defer rows.Close()

	// todo::需要一个判断数据集合为空的方法
	if !rows.NextResultSet() {
		log.Println("没有数据")
		return
	}

	for rows.Next() {
		err = rows.Scan(&user.Id, &user.Name)
		if err != nil {
			log.Println("执行失败, ", err)
			return
		}

		log.Printf("user=%+v\n", user)
	}
}

func _queryContext() {
	// 获取一行或者多行数据
	var user User
	ctx := context.Background()
	sqlStr := "select `id`,`name` from `user_copy`"
	// 执行一次select,结果可能是多行
	rows, err := db.QueryContext(ctx, sqlStr)
	if err != nil {
		log.Println("查询失败, ", err)
		return
	}

	defer rows.Close()

	isEmpty := true
	// todo::需要一个判断数据集合为空的方法
	for rows.Next() {
		isEmpty = false
		err = rows.Scan(&user.Id, &user.Name)
		if err != nil {
			log.Println("执行失败, ", err)
			return
		}

		log.Printf("user=%+v\n", user)
	}

	if isEmpty {
		log.Println("数据为空")
	}
}

func _queryRow() {
	// 获取单行数据
	var user User
	sqlStr := "select `id`,`name` from `user_copy`"
	row := db.QueryRow(sqlStr)
	err = row.Scan(&user.Id, &user.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("数据为空")
		} else {
			log.Println("查询失败, ", err)
		}

		return
	}

	log.Printf("user=%+v\n", user)
}

func _queryRowContext() {
	// 带上下文的方法，查询一行数据，如果数据为空会返回 sql.ErrNoRows
	ctx := context.Background()
	var user User
	sqlStr := "select `id`,`name` from `user_copy`"
	row := db.QueryRowContext(ctx, sqlStr)
	err = row.Scan(&user.Id, &user.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("数据为空")
		} else {
			log.Println("查询失败, ", err)
		}

		return
	}

	log.Printf("user=%+v\n", user)
}

func _execInsert() {
	// Exec执行一次命令（包括查询、删除、更新、插入等）
	sqlStr := "insert into `user`(`name`) values (?)"
	ret, err := db.Exec(sqlStr, "王五")
	if err != nil {
		log.Printf("插入失败, err:%v\n", err)
		return
	}

	newId, err := ret.LastInsertId() // 新插入数据的id
	if err != nil {
		log.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}

	log.Printf("插入成功, the id is %d.\n", newId)
}

func _execUpdate() {
	// Exec执行一次命令（包括查询、删除、更新、插入等）
	sqlStr := "update `user` set name=? where id = ?"
	ret, err := db.Exec(sqlStr, "赵六", 3)
	if err != nil {
		log.Printf("更新失败, err:%v\n", err)
		return
	}

	rowsAffect, err := ret.RowsAffected() // 影响行数
	if err != nil {
		log.Printf("get rowsAffect failed, err:%v\n", err)
		return
	}

	log.Printf("更新成功, affected rows %d.\n", rowsAffect)
}

func _execDelete() {
	// Exec执行一次命令（包括查询、删除、更新、插入等）
	sqlStr := "delete from `user` where id = ?"
	ret, err := db.Exec(sqlStr, 3)
	if err != nil {
		log.Printf("删除失败, err:%v\n", err)
		return
	}

	rowsAffect, err := ret.RowsAffected() // 影响行数
	if err != nil {
		log.Printf("get rowsAffect failed, err:%v\n", err)
		return
	}

	log.Printf("删除成功, affected rows %d.\n", rowsAffect)
}

func _prepare() {
	// 获取单行数据
	var user User
	sqlStr := "select `id`,`name` from `user_copy`"
	// 为什么要预处理？
	// 1.优化MySQL服务器重复执行SQL的方法，可以提升服务器性能，提前让服务器编译，一次编译多次执行，节省后续编译的成本。
	// 2.避免SQL注入问题。
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		log.Printf("prepare failed, err:%v\n", err)
		return
	}

	defer stmt.Close()
	row := stmt.QueryRow(sqlStr)
	err = row.Scan(&user.Id, &user.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("数据为空")
		} else {
			log.Println("查询失败, ", err)
		}

		return
	}

	log.Printf("user=%+v\n", user)
}

func _transaction() {
	tx, err := db.Begin() // 开启事务
	if err != nil {
		if tx != nil {
			_ = tx.Rollback() // 回滚
		}
		log.Printf("事务开启失败, err:%v\n", err)
		return
	}

	sqlStr1 := "update `user` set name=? where id=?"
	ret1, err := tx.Exec(sqlStr1, "哈哈", 1)
	if err != nil {
		tx.Rollback() // 回滚
		log.Printf("exec sql1 failed, err:%v\n", err)
		return
	}
	affRow1, err := ret1.RowsAffected()
	if err != nil {
		tx.Rollback() // 回滚
		log.Printf("exec ret1.RowsAffected() failed, err:%v\n", err)
		return
	}

	sqlStr2 := "update `user` set name=? where id=?"
	ret2, err := tx.Exec(sqlStr2, "李四", 2)
	if err != nil {
		tx.Rollback() // 回滚
		log.Printf("exec sql2 failed, err:%v\n", err)
		return
	}
	affRow2, err := ret2.RowsAffected()
	if err != nil {
		tx.Rollback() // 回滚
		log.Printf("exec ret1.RowsAffected() failed, err:%v\n", err)
		return
	}

	log.Println(affRow1, affRow2)
	if affRow1 == 1 && affRow2 == 1 {
		log.Println("事务提交啦...")
		tx.Commit() // 提交事务
	} else {
		tx.Rollback()
		log.Println("事务回滚啦...")
	}

	log.Println("事务执行成功!")
}
