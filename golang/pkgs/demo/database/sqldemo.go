package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"strconv"
	"time"
)

func main() {
	// 创建连接
	db, err := sql.Open("sqlite3", "./demo.db")
	checkErr(err)
	// 结束后关闭
	defer db.Close()

	// 插入
	stmt, err := db.Prepare("insert into articles (title,content,timeline) values (?,?,?)")
	checkErr(err)
	defer stmt.Close()

	res, err := stmt.Exec("title1", "这个才是内容！", timeline())
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Print("\n\n插入数据 id : ", id)

	res, err = stmt.Exec("title2", "这个才是内容2222222！", timeline())
	checkErr(err)
	id, err = res.LastInsertId()
	checkErr(err)
	fmt.Print("\n\n插入数据 id : ", id)

	// 更新
	stmt, err = db.Prepare("update articles set content=? where id=?")
	checkErr(err)
	res, err = stmt.Exec("修改内容2.。。", id)
	checkErr(err)
	affect, err := res.RowsAffected()
	checkErr(err)
	fmt.Printf("\n\n更新数据 %d ,影响数：%v ", id, affect)

	// select
	stmt, err = db.Prepare("select * from articles limit ?,?")
	checkErr(err)
	rows, err := stmt.Query(2, 5)
	checkErr(err)
	defer rows.Close()
	// 可以合并 rows, err := db.Query("select ...")
	fmt.Printf("\n\n选择id 大于2 ，5条数据")

	articles := []Article{}
	for rows.Next() {
		article := Article{}
		err = rows.Scan(&article.Id, &article.Title, &article.Content, &article.Timeline)
		checkErr(err)
		articles = append(articles, article)
	}

	for _, article := range articles {
		fmt.Printf("\n\n id => %d \n title => %s \n content => %s \n timeline => %d ", article.Id, article.Title, article.Content, article.Timeline)
	}

	// del
	stmt, err = db.Prepare("delete from articles where id=?")
	checkErr(err)
	res, err = stmt.Exec(id)
	checkErr(err)
	affect, err = res.RowsAffected()
	checkErr(err)

	fmt.Printf("\n\n删除id:%d, 影响%v\n", id, affect)

}

type Article struct {
	Id       int
	Title    string
	Content  string
	Timeline int
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func timeline() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}
