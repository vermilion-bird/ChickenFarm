package main

import (
	"chickenFarm/db"
	"chickenFarm/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

//webSocket请求ping 返回pong
func ping(c *gin.Context) {
	//升级get请求为webSocket协议
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer ws.Close()
	for {
		//读取ws中的数据

		mt, message, err := ws.ReadMessage()
		if err != nil {
			break
		}
		if string(message) == "ping" {
			message = []byte("pong")
		}
		//写入ws数据
		for {
			time.Sleep(1000 * time.Millisecond)

			msg := db.GetAllInfo()

			err = ws.WriteMessage(mt, []byte(msg))
			if err != nil {
				panic(err)
				// break
			}
		}
	}
}

func main() {
	db.Connct()
	bindAddress := "0.0.0.0:80"
	r := gin.Default()
	// r.LoadHTMLGlob("templates/**/*")
	// r.GET("/", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "html/index.html", gin.H{})
	// })
	// r.Static("/", "templates/html/")
	r.LoadHTMLGlob("./dist/*.html")
	r.StaticFile("/", "./dist/index.html")
	r.StaticFS("/h5", http.Dir("./dist"))
	// r.Static("/assets", "./assets")
	r.GET("/delHostByIP", delHostByIP)
	r.GET("/ping", ping)
	r.Run(bindAddress)
}

func delHostByIP(c *gin.Context) {
	ip := c.Query("ip")
	db.DeleteOne("ip = ?", ip, model.UpInfo{})
}
