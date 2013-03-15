package main
import (
  "fmt"
  "net/http"
  "encoding/json"
  "os"
)

type Message struct {
  Name string
  Message string
}

// 模仿存储，用一个全局变量来存储这些东西
var Messages []Message

func main() {
  
  // 测试我们的存储是否可以正常运行
  hello := Message{"mj", "你好吗？"}
  Messages = append(Messages, hello)
  fmt.Println(Messages)
  
  fmt.Println(os.Getwd())
  
  //我们先吧路由实现了
  http.HandleFunc("/messages.json", MessagesController)
  http.Handle("/", http.FileServer(http.Dir("./src/message/public")))
  err := http.ListenAndServe(":9001", nil)
  if err != nil {
    panic(err)
  }
}

func MessagesController(w http.ResponseWriter, r *http.Request) {
  if r.Method == "GET" {
    // 如果是get的话，需要什么
    messages, err := json.Marshal(Messages)
    if err != nil {
      panic(err)
    }
    fmt.Fprint(w, string(messages))
  } else if r.Method == "POST" {
    // 如果是post是什么
    message := Message{r.FormValue("Name"), r.FormValue("Message")}
    Messages = append(Messages, message)
    message_json, err := json.Marshal(message)
    if err != nil {
      panic(err)
    }
    fmt.Fprint(w, string(message_json))
  }
}