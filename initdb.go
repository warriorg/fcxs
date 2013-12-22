package main

import (
	"fmt"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"os"
)

type Build struct {
	Id   int    `bson:"_id"`
	Name string `bson:"name"`
}

type Room struct {
	Id      bson.ObjectId `bson:"_id"`
	Name    string        `bson:"name"`
	Type    string        `bson:"type"`
	Status  int           `bson:"status"`
	Price   string        `bson:"price"`
	Area    string        `bson:"area"`
	BuildId int           `bson:"buildId"`
}

func main() {
	sess, err := mgo.Dial("127.0.0.1")
	if err != nil {
		fmt.Printf("连接数据库失败")
		os.Exit(1)
	}
	defer sess.Close()

	sess.SetSafe(&mgo.Safe{})
	err = sess.DB("estates").DropDatabase()
	if err != nil {
		fmt.Printf("删除数据库失败:%v\n", err)
		os.Exit(1)
	}
	collection := sess.DB("estates").C("build")
	for i := 1; i < 22; i++ {
		doc := Build{Id: i, Name: fmt.Sprintf("%d号楼", i)}
		err = collection.Insert(doc)
		if err != nil {
			fmt.Print("新增记录失败:%v\n", err)
			os.Exit(1)
		}
	}

	fmt.Print("初始化1号楼\n")
	col := sess.DB("estates").C("room")
	for i := 1; i < 25; i++ {
		doc := Room{Id: bson.NewObjectId(), Name: fmt.Sprintf("%d01", i), Type: "2室2厅1卫", Status: 0, Area: "91.11", BuildId: 1}
		err = col.Insert(doc)
		printfErr("新增1号楼01室", err)
	}
	for i := 1; i < 25; i++ {
		doc := Room{Id: bson.NewObjectId(), Name: fmt.Sprintf("%d02", i), Type: "2室2厅1卫", Status: 0, Area: "89.39", BuildId: 1}
		err = col.Insert(doc)
		printfErr("新增1号楼02室", err)
	}
	for i := 1; i < 25; i++ {
		doc := room(fmt.Sprintf("%d03", i), "2室2厅1卫", "89.39", 1)
		err = col.Insert(doc)
		printfErr("新增1号楼03室", err)
	}
	for i := 1; i < 25; i++ {
		doc := room(fmt.Sprintf("%d04", i), "2室2厅1卫", "91.11", 1)
		err = col.Insert(doc)
		printfErr("新增1号楼04室", err)
	}
	fmt.Print("初始化2号楼\n")
	doc := room(fmt.Sprintf("%d01", 1), "4室2厅1卫", "135.65", 2)
	err = col.Insert(doc)
	printfErr("新增2号楼101", err)
	for i := 2; i < 29; i++ {
		insertCol(col, fmt.Sprintf("%d01", i), "3室2厅1卫", "124.65", 2)
	}
	for i := 2; i < 29; i++ {
		insertCol(col, fmt.Sprintf("%d02", i), "2室2厅1卫", "88.86", 2)
	}
	insertCol(col, fmt.Sprintf("%d03", 1), "3室2厅1卫", "121.99", 2)
	for i := 2; i < 29; i++ {
		insertCol(col, fmt.Sprintf("%d03", i), "3室2厅1卫", "124.65", 2)
	}

	fmt.Print("初始化3号楼\n")
	insertCol(col, fmt.Sprintf("%d01", 1), "4室2厅2卫", "148.42", 3)
	for i := 2; i < 28; i++ {
		insertCol(col, fmt.Sprintf("%d01", i), "3室2厅2卫", "137.71", 3)
	}
	insertCol(col, "2801", "5室3厅3卫", "221.9", 3)
	for i := 2; i < 28; i++ {
		insertCol(col, fmt.Sprintf("%d02", i), "2房2厅1卫", "86.48", 3)
	}
	insertCol(col, "2802", "3室3厅2卫", "172.95", 3)
	insertCol(col, fmt.Sprintf("%d03", 1), "3室2厅2卫", "135.12", 3)
	for i := 2; i < 28; i++ {
		insertCol(col, fmt.Sprintf("%d03", i), "3室2厅2卫", "137.71", 3)
	}
	insertCol(col, "2803", "5室3厅3卫", "221.9", 3)

	fmt.Print("初始化4号楼\n")
	insertCol(col, "101", "4室2厅2卫", "141.6", 4)
	for i := 2; i < 28; i++ {
		insertCol(col, fmt.Sprintf("%d01", i), "3室2厅2卫", "130.7", 4)
	}
	insertCol(col, "2801", "5室3厅3卫", "130.7", 4)
	for i := 2; i < 28; i++ {
		insertCol(col, fmt.Sprintf("%d02", i), "2房2厅1卫", "88.04", 4)
	}
	insertCol(col, "2802", "3室3厅2卫", "88.04", 4)
	insertCol(col, "103", "3室2厅2卫", "128.07", 4)
	for i := 2; i < 28; i++ {
		insertCol(col, fmt.Sprintf("%d03", i), "3室2厅2卫", "130.7", 4)
	}
	insertCol(col, "2803", "5室3厅3卫", "130.7", 4)

	fmt.Print("初始化5号楼\n")
	insertCol(col, "101", "2房2厅1卫", "112.91", 5)
	for i := 2; i < 30; i++ {
		insertCol(col, fmt.Sprintf("%d,01"), "2室2厅1卫", "140.29", 5)
	}
	insertCol(col, "2001", "5房2厅2卫", "240.14", 5)
	insertCol(col, "102", "2房2厅1卫", "120.34", 5)
	for i := 2; i < 30; i++ {
		insertCol(col, fmt.Sprintf("%d02", i), "3房2厅2卫", "137.41", 5)
	}
	insertCol(col, "3002", "5房2厅2卫", "242.35", 5)
	insertCol(col, "103", "2房2厅2卫", "120.34", 5)
	for i := 2; i < 30; i++ {
		insertCol(col, fmt.Sprintf("%d03", i), "3房2厅2卫", "137.41", 5)
	}
	insertCol(col, "3003", "5房2厅2卫", "242.35", 5)
	insertCol(col, "104", "2房2厅1卫", "112.91", 5)
	for i := 2; i < 30; i++ {
		insertCol(col, fmt.Sprintf("%d03", i), "3房2厅2卫", "140.29", 5)
	}
	insertCol(col, "3004", "5房2厅2卫", "240.14", 5)

	fmt.Print("初始化6号楼\n")
	insertCol(col, "101", "4室2厅2卫", "141.6", 6)
	for i := 2; i < 27; i++ {
		insertCol(col, fmt.Sprintf("%d01", i), "3室2厅2卫", "130.7", 6)
	}
	insertCol(col, "1701", "5室3厅3卫", "130.7", 6)
	for i := 2; i < 27; i++ {
		insertCol(col, fmt.Sprintf("%d02", i), "2房2厅1卫", "88.04", 6)
	}
	insertCol(col, "2702", "3室3厅2卫", "88.04", 6)
	insertCol(col, "103", "3室2厅2卫", "128.07", 6)
	for i := 2; i < 27; i++ {
		insertCol(col, fmt.Sprintf("%d03", i), "3室2厅2卫", "130.7", 6)
	}
	insertCol(col, "2703", "5室3厅3卫", "130.7", 6)

	fmt.Print("初始化7号楼\n")
	insertCol(col, "101", "2房2厅1卫", "113.18", 7)
	for i := 2; i < 28; i++ {
		insertCol(col, fmt.Sprintf("%d01", i), "3房2厅2卫", "140.65", 7)
	}
	insertCol(col, "2801", "5房2厅2卫", "240.74", 7)
	insertCol(col, "102", "2房2厅2卫", "113.18", 7)
	for i := 2; i < 28; i++ {
		insertCol(col, fmt.Sprintf("%d02", i), "3房2厅2卫", "140.65", 7)
	}
	insertCol(col, "2802", "5房2厅2卫", "240.74", 7)

	fmt.Print("初始化8号楼\n")
	for i := 1; i < 28; i++ {
		insertCol(col, fmt.Sprintf("%d01", i), "2房2厅1卫", "90.97", 8)
	}
	for i := 1; i < 28; i++ {
		insertCol(col, fmt.Sprintf("%d02", i), "2房2厅1卫", "89.25", 8)
	}
	for i := 1; i < 28; i++ {
		insertCol(col, fmt.Sprintf("%d03", i), "2房2厅1卫", "89.25", 8)
	}
	for i := 1; i < 28; i++ {
		insertCol(col, fmt.Sprintf("%d04", i), "2房2厅1卫", "90.97", 8)
	}

	fmt.Print("初始化9号楼\n")
	insertCol(col, "101", "2房2厅1卫", "112.91", 9)
	for i := 2; i < 28; i++ {
		insertCol(col, fmt.Sprintf("%d01", i), "3房2厅2卫", "140.29", 9)
	}
	insertCol(col, "2801", "5房2厅2卫", "240.14", 9)
	insertCol(col, "102", "2房2厅2卫", "120.34", 9)
	for i := 2; i < 28; i++ {
		insertCol(col, fmt.Sprintf("%d02", i), "3房2厅2卫", "137.41", 9)
	}
	insertCol(col, "2802", "5房2厅2卫", "242.35", 9)
	insertCol(col, "103", "2房2厅2卫", "120.34", 9)
	for i := 2; i < 28; i++ {
		insertCol(col, fmt.Sprintf("%d03", i), "3房2厅2卫", "137.41", 9)
	}
	insertCol(col, "2803", "5房2厅2卫", "242.35", 9)
	insertCol(col, "104", "2房2厅1卫", "112.91", 9)
	for i := 2; i < 28; i++ {
		insertCol(col, fmt.Sprintf("%d04", i), "3房2厅2卫", "140.29", 9)
	}
	insertCol(col, "2804", "5房2厅2卫", "240.14", 9)

	fmt.Print("初始化10号楼\n")
	for i := 1; i < 25; i++ {
		insertCol(col, fmt.Sprintf("%d01", i), "2房2厅1卫", "91.11", 10)
	}
	for i := 1; i < 25; i++ {
		insertCol(col, fmt.Sprintf("%d02", i), "2房2厅1卫", "89.39", 10)
	}
	for i := 1; i < 25; i++ {
		insertCol(col, fmt.Sprintf("%d03", i), "2房2厅1卫", "89.39", 10)
	}
	for i := 1; i < 25; i++ {
		insertCol(col, fmt.Sprintf("%d04", i), "2房2厅1卫", "91.11", 10)
	}
	for i := 1; i < 25; i++ {
		insertCol(col, fmt.Sprintf("%d05", i), "2房2厅1卫", "91.11", 10)
	}
	for i := 1; i < 25; i++ {
		insertCol(col, fmt.Sprintf("%d06", i), "2房2厅1卫", "89.39", 10)
	}
	for i := 1; i < 25; i++ {
		insertCol(col, fmt.Sprintf("%d07", i), "2房2厅1卫", "89.39", 10)
	}
	for i := 1; i < 25; i++ {
		insertCol(col, fmt.Sprintf("%d08", i), "2房2厅1卫", "91.1", 10)
	}

	fmt.Println("初始化11号楼")
	insertCol(col, "101", "4室2厅1卫", "141.6", 11)
	for i := 2; i < 27; i++ {
		insertCol(col, fmt.Sprintf("%d01", i), "3室2厅2卫", "130.7", 11)
	}
	for i := 2; i < 27; i++ {
		insertCol(col, fmt.Sprintf("%d02", i), "2房2厅1卫", "88.04", 11)
	}
	insertCol(col, "103", "3室2厅2卫", "128.07", 11)
	for i := 2; i < 27; i++ {
		insertCol(col, fmt.Sprintf("%d03", i), "3室2厅2卫", "130.7", 11)
	}

	fmt.Println("初始化12号楼")
	insertCol(col, "101", "4室2厅2卫", "148.42", 12)
	for i := 2; i < 28; i++ {
		insertCol(col, fmt.Sprintf("%d01", i), "3室2厅2卫", "137.71", 12)
	}
	insertCol(col, "2801", "5室3厅3卫", "221.9", 12)
	for i := 2; i < 28; i++ {
		insertCol(col, fmt.Sprintf("%d02", i), "2房2厅1卫", "86.48", 12)
	}
	insertCol(col, "2802", "3室3厅2卫", "172.95", 12)
	insertCol(col, "103", "3室2厅2卫", "135.12", 12)
	for i := 2; i < 28; i++ {
		insertCol(col, fmt.Sprintf("%d03", i), "3室2厅2卫", "137.71", 12)
	}
	insertCol(col, "2803", "5室3厅3卫", "221.9", 12)

	fmt.Println("初始化13号楼")
	insertCol(col, "101", "4室2厅2卫", "141.6", 13)
	for i := 2; i < 29; i++ {
		insertCol(col, fmt.Sprintf("%d01", i), "3室2厅2卫", "130.7", 13)
	}
	for i := 2; i < 29; i++ {
		insertCol(col, fmt.Sprintf("%d02", i), "2室2厅2卫", "130.7", 13)
	}
	insertCol(col, "103", "3室2厅2卫", "128.07", 13)
	for i := 2; i < 29; i++ {
		insertCol(col, fmt.Sprintf("%d01", i), "3室2厅2卫", "130.7", 13)
	}

	fmt.Println("初始化14号楼")
	insertCol(col, "101", "3房2厅2卫", "135.87", 14)
	for i := 2; i < 30; i++ {
		insertCol(col, fmt.Sprintf("%d01", i), "3房2厅2卫", "138.47", 14)
	}
	insertCol(col, "3001", "5室2厅3卫", "223.12", 14)
	for i := 2; i < 30; i++ {
		insertCol(col, fmt.Sprintf("%d02", i), "2房2厅1卫", "86.95", 14)
	}
	insertCol(col, "3002", "3室3厅2卫", "173.91", 14)
	insertCol(col, "103", "4房2厅2卫", "139.85", 14)
	for i := 2; i < 30; i++ {
		insertCol(col, fmt.Sprintf("%d03", i), "3房2厅2卫", "129.08", 14)
	}
	insertCol(col, "3003", "5室2厅3卫", "223.55", 14)
	insertCol(col, "104", "4房2厅2卫", "139.85", 14)
	for i := 2; i < 30; i++ {
		insertCol(col, fmt.Sprintf("%d04", i), "3房2厅2卫", "129.08", 14)
	}
	insertCol(col, "3004", "5室2厅3卫", "223.55", 14)
	for i := 2; i < 30; i++ {
		insertCol(col, fmt.Sprintf("%d05", i), "2房2厅1卫", "86.95", 14)
	}
	insertCol(col, "3005", "3室3厅2卫", "173.91", 14)
	insertCol(col, "106", "3房2厅2卫", "135.87", 14)
	for i := 2; i < 30; i++ {
		insertCol(col, fmt.Sprintf("%d06", i), "3房2厅2卫", "138.47", 14)
	}
	insertCol(col, "3006", "5室2厅3卫", "223.12", 14)

	/***15号楼***/
	fmt.Println("初始化15号楼")
	insertCol(col, "101", "4室2厅2卫", "141.6", 15)
	for i := 2; i < 28; i++ {
		insertCol(col, fmt.Sprintf("%d01", i), "3室2厅2卫", "130.7", 15)
	}
	insertCol(col, "2801", "5室3厅3卫", "130.7", 15)
	for i := 2; i < 28; i++ {
		insertCol(col, fmt.Sprintf("%d02", i), "2室2厅2卫", "88.04", 15)
	}
	insertCol(col, "2802", "3室3厅2卫", "88.04", 15)
	insertCol(col, "103", "3室2厅2卫", "128.07", 15)
	for i := 2; i < 28; i++ {
		insertCol(col, fmt.Sprintf("%d03", i), "3室2厅2卫", "130.7", 15)
	}
	insertCol(col, "2803", "5室3厅3卫", "130.7", 15)

	/**********16号楼*******/
	fmt.Println("初始化16号楼")
	insertCol(col, "101", "2房2厅1卫", "113.2", 16)
	for i := 2; i < 28; i++ {
		insertCol(col, fmt.Sprintf("%d01", i), "3房2厅2卫", "140", 16)
	}
	insertCol(col, "2801", "5房2厅2卫", "240.76", 16)
	insertCol(col, "102", "2房2厅2卫", "113.2", 16)
	for i := 2; i < 28; i++ {
		insertCol(col, fmt.Sprintf("%d02", i), "3房2厅2卫", "140.64", 16)
	}
	insertCol(col, "2802", "5房2厅2卫", "240.75", 16)

	/**************17号楼***************/
	fmt.Println("初始化17号楼")
	insertCol(col, "101", "2房2厅2卫", "123.07", 17)
	for i := 2; i < 28; i++ {
		insertCol(col, fmt.Sprintf("%d01", i), "3房2厅2卫", "123.07", 17)
	}
	insertCol(col, "102", "2房2厅2卫", "123.07", 17)
	for i := 2; i < 28; i++ {
		insertCol(col, fmt.Sprintf("%d02", i), "3房2厅2卫", "123.07", 17)
	}
	insertCol(col, "103", "2房2厅2卫", "123.07", 17)
	for i := 2; i < 28; i++ {
		insertCol(col, fmt.Sprintf("%d03", i), "3房2厅2卫", "123.07", 17)
	}
	insertCol(col, "104", "2房2厅2卫", "123.07", 17)
	for i := 2; i < 28; i++ {
		insertCol(col, fmt.Sprintf("%d04", i), "3房2厅2卫", "123.07", 17)
	}

	/**********18号楼*******/
	fmt.Println("初始化18号楼")
	insertCol(col, "101", "2房2厅1卫", "113.2", 18)
	for i := 2; i < 28; i++ {
		insertCol(col, fmt.Sprintf("%d01", i), "3房2厅2卫", "140", 18)
	}
	insertCol(col, "2801", "5房2厅2卫", "240.76", 18)
	insertCol(col, "102", "2房2厅2卫", "113.2", 18)
	for i := 2; i < 28; i++ {
		insertCol(col, fmt.Sprintf("%d02", i), "3房2厅2卫", "140.64", 18)
	}
	insertCol(col, "2802", "5房2厅2卫", "240.75", 18)

	/**********19号楼*******/
	fmt.Println("初始化19号楼")
	insertCol(col, "101", "2房2厅2卫", "123.07", 19)
	for i := 2; i < 28; i++ {
		insertCol(col, fmt.Sprintf("%d01", i), "3房2厅2卫", "143.55", 19)
	}
	insertCol(col, "2801", "3房2厅2卫", "142.52", 19)
	insertCol(col, "102", "2房2厅2卫", "123.07", 19)
	for i := 2; i < 28; i++ {
		insertCol(col, fmt.Sprintf("%d02", i), "3房2厅2卫", "143.55", 19)
	}
	insertCol(col, "2802", "3房2厅2卫", "142.52", 19)
	insertCol(col, "103", "2房2厅2卫", "123.07", 19)
	for i := 2; i < 28; i++ {
		insertCol(col, fmt.Sprintf("%d03", i), "3房2厅2卫", "143.55", 19)
	}
	insertCol(col, "2803", "3房2厅2卫", "142.52", 19)
	insertCol(col, "104", "2房2厅2卫", "123.07", 19)
	for i := 2; i < 28; i++ {
		insertCol(col, fmt.Sprintf("%d04", i), "3房2厅2卫", "143.55", 19)
	}
	insertCol(col, "2804", "3房2厅2卫", "142.52", 19)

	/**********20号楼*******/
	fmt.Println("初始化20号楼")
	insertCol(col, "101", "商铺", "156.70", 20)
	insertCol(col, "102", "商铺", "156.70", 20)
	insertCol(col, "103", "商铺", "156.70", 20)
	insertCol(col, "104", "商铺", "156.70", 20)
	insertCol(col, "105", "商铺", "117.56", 20)
	insertCol(col, "106", "商铺", "117.56", 20)
	insertCol(col, "107", "商铺", "117.56", 20)
	insertCol(col, "108", "商铺", "117.56", 20)
	insertCol(col, "109", "商铺", "59.05", 20)
	insertCol(col, "110", "商铺", "59.05", 20)
	insertCol(col, "111", "商铺", "63.59", 20)
	insertCol(col, "112", "商铺", "63.59", 20)
	insertCol(col, "113", "商铺", "63.59", 20)
	insertCol(col, "114", "商铺", "63.59", 20)
	insertCol(col, "115", "商铺", "63.59", 20)
	insertCol(col, "116", "商铺", "63.59", 20)
	insertCol(col, "117", "商铺", "63.59", 20)
	insertCol(col, "118", "商铺", "63.59", 20)
	insertCol(col, "119", "商铺", "63.59", 20)
	insertCol(col, "120", "商铺", "63.59", 20)
	insertCol(col, "121", "商铺", "63.59", 20)
	insertCol(col, "122", "商铺", "63.59", 20)
	insertCol(col, "123", "商铺", "63.59", 20)
	insertCol(col, "124", "商铺", "63.59", 20)
	insertCol(col, "125", "商铺", "537.69", 20)
	insertCol(col, "126", "商铺", "149.36", 20)
	insertCol(col, "127", "商铺", "149.36", 20)
	insertCol(col, "128", "商铺", "113.55", 20)
	insertCol(col, "129", "商铺", "113.55", 20)
	insertCol(col, "130", "商铺", "113.55", 20)
	insertCol(col, "131", "商铺", "113.55", 20)
	insertCol(col, "132", "商铺", "113.55", 20)
	insertCol(col, "133", "商铺", "113.55", 20)
	insertCol(col, "134", "商铺", "113.55", 20)
	insertCol(col, "135", "商铺", "113.55", 20)
	/*****-----------------**/
	insertCol(col, "340", "商铺", "965.53", 20)
	insertCol(col, "230", "商铺", "36.38", 20)

	insertCol(col, "201", "酒店公寓", "63.44", 20)
	insertCol(col, "202", "酒店公寓", "63.44", 20)
	insertCol(col, "203", "酒店公寓", "63.44", 20)
	insertCol(col, "204", "酒店公寓", "53.54", 20)
	insertCol(col, "205", "酒店公寓", "63.44", 20)
	insertCol(col, "206", "酒店公寓", "53.54", 20)
	insertCol(col, "207", "酒店公寓", "63.44", 20)
	insertCol(col, "208", "酒店公寓", "53.54", 20)
	insertCol(col, "209", "酒店公寓", "63.44", 20)
	insertCol(col, "210", "酒店公寓", "53.54", 20)
	insertCol(col, "211", "酒店公寓", "63.44", 20)
	insertCol(col, "212", "酒店公寓", "53.54", 20)
	insertCol(col, "213", "酒店公寓", "63.44", 20)
	insertCol(col, "214", "酒店公寓", "53.54", 20)
	insertCol(col, "215", "酒店公寓", "63.44", 20)
	insertCol(col, "216", "酒店公寓", "53.54", 20)
	insertCol(col, "217", "酒店公寓", "63.44", 20)
	insertCol(col, "218", "酒店公寓", "53.54", 20)
	insertCol(col, "219", "酒店公寓", "63.44", 20)
	insertCol(col, "220", "酒店公寓", "53.54", 20)
	insertCol(col, "221", "酒店公寓", "63.44", 20)
	insertCol(col, "222", "酒店公寓", "53.54", 20)
	insertCol(col, "223", "酒店公寓", "63.44", 20)
	insertCol(col, "224", "酒店公寓", "53.54", 20)
	insertCol(col, "225", "酒店公寓", "63.44", 20)
	insertCol(col, "226", "酒店公寓", "53.54", 20)
	insertCol(col, "227", "酒店公寓", "63.44", 20)
	insertCol(col, "228", "酒店公寓", "59.05", 20)
	insertCol(col, "229", "酒店公寓", "63.44", 20)

	insertCol(col, "301", "酒店公寓", "53.79", 20)
	insertCol(col, "302", "酒店公寓", "54.38", 20)
	insertCol(col, "303", "酒店公寓", "53.79", 20)
	insertCol(col, "304", "酒店公寓", "91.86", 20)
	insertCol(col, "305", "酒店公寓", "53.79", 20)
	insertCol(col, "306", "酒店公寓", "53.79", 20)
	insertCol(col, "307", "酒店公寓", "74.21", 20)
	insertCol(col, "308", "酒店公寓", "63.75", 20)
	insertCol(col, "309", "酒店公寓", "63.75", 20)
	insertCol(col, "310", "酒店公寓", "63.75", 20)
	insertCol(col, "311", "酒店公寓", "63.75", 20)
	insertCol(col, "312", "酒店公寓", "63.75", 20)
	insertCol(col, "313", "酒店公寓", "63.75", 20)
	insertCol(col, "314", "酒店公寓", "54.38", 20)
	insertCol(col, "315", "酒店公寓", "63.75", 20)
	insertCol(col, "316", "酒店公寓", "54.38", 20)
	insertCol(col, "317", "酒店公寓", "63.75", 20)
	insertCol(col, "318", "酒店公寓", "44.38", 20)
	insertCol(col, "319", "酒店公寓", "63.44", 20)
	insertCol(col, "320", "酒店公寓", "54.11", 20)
	insertCol(col, "321", "酒店公寓", "63.75", 20)
	insertCol(col, "322", "酒店公寓", "54.38", 20)
	insertCol(col, "323", "酒店公寓", "63.75", 20)
	insertCol(col, "324", "酒店公寓", "54,38", 20)
	insertCol(col, "325", "酒店公寓", "63.75", 20)
	insertCol(col, "326", "酒店公寓", "54.38", 20)
	insertCol(col, "327", "酒店公寓", "63.75", 20)
	insertCol(col, "328", "酒店公寓", "54.38", 20)
	insertCol(col, "329", "酒店公寓", "63.75", 20)
	insertCol(col, "330", "酒店公寓", "54.38", 20)
	insertCol(col, "331", "酒店公寓", "63.74", 20)
	insertCol(col, "332", "酒店公寓", "54.38", 20)
	insertCol(col, "333", "酒店公寓", "63.75", 20)
	insertCol(col, "334", "酒店公寓", "54.38", 20)
	insertCol(col, "335", "酒店公寓", "63.75", 20)
	insertCol(col, "336", "酒店公寓", "54.38", 20)
	insertCol(col, "337", "酒店公寓", "63.75", 20)
	insertCol(col, "338", "酒店公寓", "54.38", 20)
	insertCol(col, "339", "酒店公寓", "63.75", 20)

	/**********21号楼*******/
	fmt.Println("初始化21号楼")
	insertCol(col, "101", "商铺", "113.87", 21)
	insertCol(col, "102", "商铺", "113.87", 21)
	insertCol(col, "103", "商铺", "113.49", 21)
	insertCol(col, "104", "商铺", "113.49", 21)
	insertCol(col, "105", "商铺", "114.41", 21)
	insertCol(col, "106", "商铺", "114.41", 21)
	insertCol(col, "107", "商铺", "114.41", 21)
	insertCol(col, "108", "商铺", "114.41", 21)
	insertCol(col, "109", "商铺", "150.49", 21)
	insertCol(col, "110", "商铺", "150.49", 21)
	insertCol(col, "111", "商铺", "541.74", 21)
	insertCol(col, "112", "商铺", "118.98", 21)
	insertCol(col, "113", "商铺", "118.98", 21)
	insertCol(col, "114", "商铺", "118.98", 21)
	insertCol(col, "115", "商铺", "118.98", 21)
	insertCol(col, "201", "商铺", "44.48", 21)
	insertCol(col, "301", "商铺", "1428.87", 21)

}

func insertCol(col *mgo.Collection, name string, t string, area string, buildId int) {
	doc := room(name, t, area, buildId)
	err := col.Insert(doc)
	printfErr(name, err)
}

func room(name string, t string, area string, buildId int) Room {
	return Room{Id: bson.NewObjectId(), Name: name, Type: t, Status: 0, Area: area, BuildId: buildId}
}

func printfErr(str string, err error) {
	if err != nil {
		fmt.Print("%s:%v\n", str, err)
		os.Exit(1)
	}
}
