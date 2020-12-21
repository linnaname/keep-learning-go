package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
	"io"
	"log"
	"net/http"
	"os"
)

var secrets = gin.H{
	"foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
	"austin": gin.H{"email": "austin@example.com", "phone": "666"},
	"lena":   gin.H{"email": "lena@guapa.com", "phone": "523443"},
}

func main() {
	r := gin.Default()

	gin.DisableConsoleColor()

	// 记录到文件。
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}

	//example
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	//json
	r.GET("/someJSON", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "GO语言",
			"tag":  "<br>",
		}

		// 输出 : {"lang":"GO\u8bed\u8a00","tag":"\u003cbr\u003e"}
		c.AsciiJSON(http.StatusOK, data)
	})

	r.LoadHTMLFiles("./templates/index.tmpl")

	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})

	r.GET("/JSONP", func(c *gin.Context) {
		data := map[string]interface{}{
			"foo": "bar",
		}

		// /JSONP?callback=x
		// 将输出：x({\"foo\":\"bar\"})
		c.JSONP(http.StatusOK, data)
	})

	r.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")

		c.JSON(200, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})

	r.GET("/json", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})

	// 提供字面字符
	r.GET("/purejson", func(c *gin.Context) {
		c.PureJSON(200, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})

	r.POST("/post", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		message := c.PostForm("message")
		fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
	})

	r.GET("/secureJSON", func(c *gin.Context) {
		names := []string{"lena", "austin", "foo"}

		// 将输出：while(1);["lena","austin","foo"]
		c.SecureJSON(http.StatusOK, names)
	})

	r.GET("/moreJSON", func(c *gin.Context) {
		// 你也可以使用一个结构体
		var msg struct {
			Name    string `json:"user"`
			Message string
			Number  int
		}
		msg.Name = "Lena"
		msg.Message = "hey"
		msg.Number = 123
		// 注意 msg.Name 在 JSON 中变成了 "user"
		// 将输出：{"user": "Lena", "Message": "hey", "Number": 123}
		c.JSON(http.StatusOK, msg)
	})

	r.GET("/someXML", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	r.GET("/someYAML", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	r.GET("/someProtoBuf", func(c *gin.Context) {
		reps := []int64{int64(1), int64(2)}
		label := "test"
		// protobuf 的具体定义写在 testdata/protoexample 文件中。
		data := &protoexample.Test{
			Label: &label,
			Reps:  reps,
		}
		// 请注意，数据在响应中变为二进制数据
		// 将输出被 protoexample.Test protobuf 序列化了的数据
		c.ProtoBuf(http.StatusOK, data)
	})

	r.MaxMultipartMemory = 8 << 20 // 8 MiB
	r.POST("/upload", func(c *gin.Context) {
		// 单文件
		file, _ := c.FormFile("file")
		log.Println(file.Filename)

		// 上传文件至指定目录
		c.SaveUploadedFile(file, "/Users/goranka/linnana/go/keep-learning-go/")
		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})

	r.GET("/someDataFromReader", func(c *gin.Context) {
		response, err := http.Get("https://raw.githubusercontent.com/gin-gonic/logo/master/color.png")
		if err != nil || response.StatusCode != http.StatusOK {
			c.Status(http.StatusServiceUnavailable)
			return
		}

		reader := response.Body
		contentLength := response.ContentLength
		contentType := response.Header.Get("Content-Type")

		extraHeaders := map[string]string{
			"Content-Disposition": `attachment; filename="gopher.png"`,
		}

		c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
	})

	authorized := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"foo":    "bar",
		"austin": "1234",
		"lena":   "hello2",
		"manu":   "4321",
	}))

	// /admin/secrets 端点
	// 触发 "localhost:8080/admin/secrets
	authorized.GET("/secrets", func(c *gin.Context) {
		// 获取用户，它是由 BasicAuth 中间件设置的
		user := c.MustGet(gin.AuthUserKey).(string)
		if secret, ok := secrets[user]; ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
		}
	})

	r.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname") // c.Request.URL.Query().Get("lastname") 的一种快捷方式

		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})

	r.Run()

}
