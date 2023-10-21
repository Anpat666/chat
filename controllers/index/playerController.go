package index

import (
	"fmt"
	"lottery/models"
	"lottery/utils"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)

type PlayerController struct{}

func (con PlayerController) Register(c *gin.Context) {
	c.HTML(200, "register.html", gin.H{
		"code": 1,
	})
}

func (con PlayerController) RegisterConfirm(c *gin.Context) {
	player := models.Player{}
	username := c.PostForm("username")
	password := c.PostForm("password")
	repassword := c.PostForm("repassword")
	data := models.FindUserByName(username)

	if len(username) <= 4 {
		c.JSON(200, gin.H{
			"code":    0,
			"message": "用户名不能少于5位",
		})
		return
	}

	if len(password) <= 4 {
		c.JSON(200, gin.H{
			"code":    0,
			"message": "密码不能少于5位",
		})
		return
	}

	if data.UserName != "" {
		c.JSON(200, gin.H{
			"code":    0,
			"message": "该用户已存在",
		})
		return
	}

	if password != repassword {
		fmt.Println(password)
		fmt.Println(repassword)
		c.JSON(200, gin.H{
			"code":    0,
			"message": "两次密码不一致",
		})
		return
	}

	random := fmt.Sprintf("%06d", rand.Int31())
	player.UserName = username
	player.Password = password
	player.LoginTime = time.Now()
	player.PlayerId = rand.Intn(1000) + 10000
	player.Random = random
	player.HassPassword = utils.MakePasswrod(password, random)
	models.DB.Create(&player)

	c.JSON(200, gin.H{
		"code":    1,
		"message": "创建成功",
	})

}

func (con PlayerController) Login(c *gin.Context) {
	player := models.Player{}
	username := c.PostForm("username")
	password := c.PostForm("password")

	models.DB.Where("user_name=?", username).Find(&player)
	if player.UserName == "" {
		c.JSON(200, gin.H{
			"code":    0,
			"message": "用户名错误",
		})
		return
	}

	hashpdw := utils.VaildPassword(password, player.Random, player.HassPassword)
	if hashpdw == false {
		c.JSON(200, gin.H{
			"code":    0,
			"message": "密码错误",
		})
		return
	}

	c.JSON(200, gin.H{
		"code":    1,
		"message": "登录成功",
	})
}
