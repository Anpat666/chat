package service

import (
	"chat/models"
	utils "chat/utils"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// GetUserList
// @Summary 用户列表查询
// @Tags 用户模块
// @Success 200 {string} json{"code","message"}
// @Router /user [get]
func GetUserList(c *gin.Context) {
	user := models.GetUserList()
	c.JSON(200, gin.H{
		"code":    0,
		"message": "查询成功",
		"data":    user,
	})
}

// UserCreate
// @Summary 新增用户
// @Tags 用户模块
// @param name query string false "用户名"
// @param phone query string false "手机号码"
// @param password query string false "密码"
// @param repassword query string false "确认密码"
// @Success 200 {string} json{"code","message"}
// @Router /user/create [get]
func UserCreate(c *gin.Context) {
	var user models.UserBasic
	user.Name = c.Query("name")
	password := c.Query("password")
	repassword := c.Query("repassword")
	user.Phone = c.Query("phone")
	resName := models.FindUserByName(user.Name)
	resPhone := models.FindUserByPhone(user.Phone)

	if resName.Name != "" {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "用户名已存在",
		})
		return
	}

	if resPhone.Phone != "" {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "手机号码已存在",
		})
		return
	}

	if password != repassword {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "两次密码输入不一致",
		})
		return
	}

	user.Salt = utils.MakeRandomSalt(16)
	user.HashPassword = utils.EncryptionPassword(user.Salt, password)

	models.UserCreate(user)
	c.JSON(200, gin.H{
		"code":    0,
		"message": "新增用户成功",
		"data":    user,
	})
}

// UserDelete
// @Summary 删除用户
// @Tags 用户模块
// @param id query string false "id"
// @Success 200 {string} json{"code","message"}
// @Router /user/delete [get]
func UserDelete(c *gin.Context) {
	var user models.UserBasic
	idstr := c.Query("id")
	id, err := strconv.ParseUint(idstr, 10, 64)
	if err != nil {
		c.JSON(-1, gin.H{
			"code":    -1,
			"message": "ID输入有误",
		})
		return
	}
	user.ID = uint(id)
	models.UserDelete(user)
	c.JSON(200, gin.H{
		"code":    0,
		"message": "删除成功",
	})

}

// UserUpdata
// @Summary 修改用户手机和邮箱
// @Tags 用户模块
// @param id formData string false "id"
// @param phone formData string false "phone"
// @param email formData string false "email"
// @Success 200 {string} json{"code","message"}
// @Router /user/updata [post]
func UserUpdata(c *gin.Context) {
	var user models.UserBasic
	id, _ := strconv.Atoi(c.PostForm("id"))
	user.ID = uint(id)
	user.Phone = c.PostForm("phone")
	user.Email = c.PostForm("email")

	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		fmt.Println("正则表达式出错：", err)
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "手机或邮箱格式错误",
		})
		return
	} else {
		models.UserUpdataPhoneAndEmail(user)
		c.JSON(200, gin.H{
			"code":    0,
			"message": "用户信息修改成功",
			"data":    user,
		})
	}

}

// UserLogin
// @Summary 用户登录
// @Tags 用户模块
// @param name formData string false "用户名"
// @param password formData string false "登录密码"
// @Success 200 {string} json{"code","message"}
// @Router /user/login [post]
func UserLogin(c *gin.Context) {
	name := c.PostForm("name")
	password := c.PostForm("password")
	user := models.FindUserByName(name)
	if user.Name == "" {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "用户名不存在",
		})
		return
	}

	CheckHashPwd := utils.DecryptPassword(user.Salt, password, user.HashPassword)
	if !CheckHashPwd {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "密码错误",
		})
		return
	}

	models.LoginUpdataUserInfo(user)

	c.JSON(200, gin.H{
		"code":    0,
		"message": "用户登录成功",
		"data":    user,
	})

}

// UserChangePassword
// @Summary 修改密码
// @Tags 用户模块
// @param name formData string false "用户名"
// @param oldPassword formData string false "旧密码"
// @param newPassword formData string false "新密码"
// @param checkNewPassword formData string false "确认新密码"
// @Success 200 {string} json{"code","message"}
// @Router /user/changepwd [post]
func UserChangePassword(c *gin.Context) {

	name := c.PostForm("name")
	oldPassword := c.PostForm("oldPassword")
	newPassword := c.PostForm("newPassword")
	checkNewPassword := c.PostForm("checkNewPassword")
	user := models.FindUserByName(name)
	if user.Name == "" {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "用户名不存在",
		})
		return
	}

	CheckHashPwd := utils.DecryptPassword(user.Salt, oldPassword, user.HashPassword)
	if !CheckHashPwd {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "原密码错误",
		})
		return
	}

	if newPassword != checkNewPassword {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "确认新密码输入有误",
		})
		return
	}

	models.UpdataPasswordById(user.ID, newPassword)
	user = models.FindUserByName(user.Name)

	c.JSON(200, gin.H{
		"code":    0,
		"message": "用户密码修改成功",
		"data":    user,
	})

}

var upGrade = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func MsgHandler(ws *websocket.Conn, c *gin.Context) {
	for {
		msg, err := utils.RedisSubscribe(c, utils.PublishKey)
		if err != nil {
			fmt.Println(err)
		}
		tm := time.Now().Format("2006-01-02 15:04:05")
		strMsg := fmt.Sprintf("[ws][%s]:%s", tm, msg)
		err = ws.WriteMessage(1, []byte(strMsg))
		if err != nil {
			fmt.Println(err)
		}
	}
}

func SendMsg(c *gin.Context) {
	ws, err := upGrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
	}
	defer func(ws *websocket.Conn) {
		err = ws.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(ws)

	MsgHandler(ws, c)
}
