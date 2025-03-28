package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// 连接到数据库
	db, err := sql.Open("mysql", "root:admin@/test?charset=utf8")
	checkErr(err)

	// 插入数据
	stmt, err := db.Prepare("INSERT INTO userinfo SET username=?,department=?,created=?")
	checkErr(err)

	// 执行插入操作
	res, err := stmt.Exec("tonya", "研发部门", "2012-12-06")
	checkErr(err)

	// 获取最后插入的ID
	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println(id)

	// 更新数据
	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)

	// 执行更行操作
	res, err = stmt.Exec("tonyaupdate", id)
	checkErr(err)

	// 获取受影响的行数
	affect, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(affect)

	// 查询数据--方法一
	stmt, err = db.Prepare("select * from userinfo")
	checkErr(err)
	rows, err := stmt.Query()
	checkErr(err)
	for rows.Next() {
		var uid int
		var username, department, created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid, username, department, created)
	}

	// 查询数据--方法二
	rows, err = db.Query("select * from userinfo")
	checkErr(err)
	for rows.Next() {
		var uid int
		var username, department, created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid, username, department, created)
	}

	// 删除数据
	stmt, err = db.Prepare("delete from userinfo where uid=?")
	checkErr(err)

	// 执行删除操作
	res, err = stmt.Exec(id)
	checkErr(err)

	// 获取受影响的行数
	affect, err = res.RowsAffected()
	checkErr(err)
	fmt.Println(affect) // 输出受影响的行数

	// 关闭数据库连接
	db.Close()
}
