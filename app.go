package main

import(
  "encoding/json"
  "fmt"
  "log"
  "labix.org/v2/mgo"
  "labix.org/v2/mgo/bson"
  "net/http"
  "strconv"
)

type Room struct {
  Id bson.ObjectId `bson:"_id"`
  Name string `bson:"name"`
  Type string `bson:"type"`
  Status int `bson:"status"`
  Price string `bson:"price"`
  Area string `bson:"area"`
  BuildId int `bson:"buildId"`
}

//查询售楼信息
func rooms(rw http.ResponseWriter, req *http.Request){
  req.ParseForm() //解析参数，默认是不会解析的 
  sess, err := mgo.Dial("127.0.0.1")
  if err != nil {
   fmt.Printf("连接数据库失败");
  }
  defer sess.Close()
  log.Println("id的值:", req.FormValue("id"))
  sess.SetSafe(&mgo.Safe{})
  col := sess.DB("estates").C("room") 
  rooms := []Room{}
  buildId, err := strconv.Atoi(req.FormValue("id"))
  err = col.Find(bson.M{"buildId":buildId}).All(&rooms)
  result,_ := json.Marshal(rooms)
  fmt.Fprintf(rw, string(result));
}

//更新售房信息
func updateRoom(rw http.ResponseWriter, req *http.Request) {
  req.ParseForm();

  sess, err := mgo.Dial("127.0.0.1")
  if err != nil {
    fmt.Fprintf(rw, "连接数据库失败");  
  }
  defer sess.Close()
  sess.SetSafe(&mgo.Safe{})
  status,err := strconv.Atoi(req.FormValue("status"))
  _id := req.FormValue("id")
  log.Println(_id,status)
  col := sess.DB("estates").C("room")
  err = col.Update(bson.M{"_id":bson.ObjectIdHex(_id)}, bson.M{"$set":bson.M{"status":status}})
  if err != nil {
    log.Println(err)
    fmt.Fprintf(rw, "false")
  } else {
    fmt.Fprintf(rw, "true")
  }
 }

func main() {
  http.HandleFunc("/build", rooms)
  http.HandleFunc("/updateRoom", updateRoom)
  log.Fatal(http.ListenAndServe(":8002", nil))
}
