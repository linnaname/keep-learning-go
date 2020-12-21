package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

var (
	id      int64
	name    string
	country string
	icon    string
)

func main() {

	//open and ping
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/pubg_diary")
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		return
	}

	//query
	rows, err := db.Query("select id, name, country, icon from team where id = ?", 1)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	/**
	结果集(rows)未关闭前，底层的连接处于繁忙状态。
	当遍历读到最后一条记录时，会发生一个内部EOF错误，自动调用rows.Close()，但是如果提前退出循环，rows不会关闭，
	连接不会回到连接池中，连接也不会关闭。所以手动关闭非常重要。rows.Close()可以多次调用，是无害操作
	*/
	for rows.Next() {
		err := rows.Scan(&id, &name, &country, &icon)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, name, country, icon)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
		return
	}

	//NULL
	var s sql.NullString
	err = db.QueryRow("select name from team where id=?", 6).Scan(&s)
	if s.Valid {
		fmt.Println(s)
	} else {
		fmt.Println("null string")
	}

	//queryrow
	err = db.QueryRow("select name from team where id = ?", 5).Scan(&name)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(name)

	/**
	事务:*sql.Tx一旦释放，连接就回到连接池中，这里stmt在关闭时就无法找到连接。所以必须在Tx commit或rollback之前关闭statement。
	*/
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()
	stmt, err := tx.Prepare("INSERT INTO team VALUES (?,?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	for i := 10; i < 20; i++ {
		_, err = stmt.Exec(i, "na"+strconv.Itoa(i), "CN", "www.baidu.com")
		if err != nil {
			log.Fatal(err)
		}
	}
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
	stmt.Close()

	//insert

	//预编译之后可以调用exec/query/queryrow
	stm, err := db.Prepare("INSERT INTO team(id,name,country,icon) VALUES(?,?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stm.Exec(4, "Dolly", "China", "www.linnana.me")
	if err != nil {
		log.Fatal(err)
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ID = %d, affected = %d\n", lastId, rowCnt)

}
