package main

import (
	//"database/sql"
	"htyx/config"
	"htyx/lib/errno"
	"htyx/lib/token"
	"htyx/model"
	//"htyx/handler/category"
	"fmt"
	//"htyx/model/categorymodel"
	//"htyx/model/articlemodel"
	//"htyx/model/audiomodel"
	//"htyx/model/likelistmodel"
	//"htyx/model/usermodel"
	//"htyx/model/commentsmodel"

	//"github.com/jmoiron/sqlx"
	//"encoding/json"

	"github.com/lexkong/log"
	"github.com/spf13/viper"
)

func main() {
	err := config.Init("")
	if err != nil {
		panic(err)
	}

	fmt.Println(viper.GetString("log.writers"))

	err = errno.NewLogErr(errno.ErrUserNotFound, fmt.Errorf("username can not found in db: xx.xx.xx.xx")).Add("This is add message.")
	log.Errorf(err, "Get an error")
	code, message := errno.DecodeErr(err)
	log.Infof("jiexijieguo %d %s", code, message)
	if errno.IsErrUserNotFound(err) {
		log.Debug("err type is ErrUserNotFound")
	}
	fmt.Println(fmt.Sprintf(viper.GetString("wxurl"), "你大爷"))
	fmt.Println("以上测试config log 库")

	//以上测试config log 库
	//以下测试sql连接和token库
	fmt.Println("以下测试sql连接和token库")
	py := token.PayLoad{Uid: 1}
	ts, _ := token.Sign(py, "Rtg8BPKNEf2mB4mgvKONGPZZQSaJWNLijxR42qRgq0iBb5")
	fmt.Println(ts)
	//time.Sleep(9 * time.Second)
	py1, err := token.Parse(ts, "Rtg8BPKNEf2mB4mgvKONGPZZQSaJWNLijxR42qRgq0iBb5")
	fmt.Println(py1, err)
	fmt.Println("以下数据库初始化并建立一张place表")
	model.Init()
	/*
		xx := `CREATE TABLE place(
		country text);`
		ss, err := model.DB.Self.Exec(xx)
		if err != nil {
			fmt.Println(ss, err)
		}
	*/
	//一定要大写 建表时default 不能是null，可以用''处理
	/*
		type Ss struct {
			ShengShi string
			DiQu     string
		}
		dx := Ss{}
		err = model.DB.Self.Get(&dx, "select shengshi from test1 limit1")
		fmt.Println(err)
		err = model.DB.Self.Get(&dx, "select diqu from test3 where id=?", 2)
		fmt.Println(err)
		fmt.Println(dx)
		fmt.Println(dx.ShengShi)
		if dx.DiQu == "" {
			fmt.Println("kongge")
		}
	*/
	/*
		data, err := classicmodel.GetLatestClassic(1)
		fmt.Println(data, err)
		datastring, _ := json.Marshal(data)
		fmt.Printf("%s\n", datastring)
		fmt.Println("--------------------------")
		data1, err1 := classicmodel.GetFavorClassics(1, 0)
		fmt.Println(data1, err1)
		datastring1, _ := json.Marshal(data1)
		fmt.Printf("%s", datastring1)*/
	/*
		fmt.Println("----测试--categorymodel.GetCategory--------------------")
		data, err := categorymodel.GetCategory(1)
		datastring, _ := json.Marshal(data)
		fmt.Printf("%s\n", datastring)
		fmt.Println("--------------------------")
	*/
	/*
		fmt.Println("----测试--articlemodel--------------------")
		data, err := articlemodel.GetListByCateId(2, 0)
		fmt.Println(err)
		datastring, _ := json.Marshal(data)
		fmt.Printf("%s\n", datastring)
		fmt.Println("--------------------------")
	*/
	/*data1, err1 := articlemodel.GetOne(2)
	fmt.Println(err1)
	datastring1, _ := json.Marshal(data1)
	fmt.Printf("%s\n", datastring1)
	fmt.Println("--------------------------")*/
	/*
		fmt.Println("----测试--audiomodel--------------------")
		data, err := audiomodel.GetListByCateId(5, 0)
		fmt.Println(err)
		datastring, _ := json.Marshal(data)
		fmt.Printf("%s\n", datastring)
		fmt.Println("--------------------------")

		data1, err1 := audiomodel.GetOne(2)
		fmt.Println(err1)
		datastring1, _ := json.Marshal(data1)
		fmt.Printf("%s\n", datastring1)
		fmt.Println("--------------------------")
	*/
	/*
		fmt.Println("----测试--likemodel--------------------")
		data, err := likelistmodel.GetListByCltype(1, 2, 0)
		fmt.Println(err)
		datastring, _ := json.Marshal(data)
		fmt.Printf("%s\n", datastring)
		fmt.Println("--------------------------")*/

	/*fmt.Println("----测试--usermodel--------------------")
	usermodel.ChangeLikeStatus(3, 1, 2)

	usermodel.ChangeCommentLikeStatus(3, 3)
	fmt.Println("--------------------------")*/

	/*
			fmt.Println("----测试--commentsmodel--------------------")
			data1, err1 := commentsmodel.GetNewComments(1, 1, 1, 0)
			fmt.Println(err1)
			datastring1, _ := json.Marshal(data1)
			fmt.Printf("%s\n", datastring1)
			fmt.Println("--------------------------")
		data, err := commentsmodel.GetHotComments(1, 1, 1, 0)
		fmt.Println(err)
		datastring, _ := json.Marshal(data)
		fmt.Printf("%s\n", datastring)
		commentsmodel.CreateComment(1, 1, "小米3", "小米的评论3", 4)
		fmt.Println("--------------------------")*/

}
