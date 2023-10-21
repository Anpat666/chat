package models

import (
	"math/rand"
	"time"

	"gorm.io/gorm"
)

type Player struct {
	gorm.Model
	PlayerId        int       //会员ID
	UserName        string    //登录用户名
	Password        string    //密码
	HassPassword    string    //密码哈希
	SubName         string    //会员昵称
	Random          string    //随机种子
	Balance         int       //余额
	Enable          bool      //启用
	Deposit         int       //存款
	WithdrawMoney   int       //提款
	SumBetAmount    float64   //总投注额度
	SumPayoutAmount float64   //总派彩额度
	LoginTime       time.Time //最近登录时间
}

func (p *Player) TableName() string {
	return "player"
}

func PlayerList() []*Player {
	var players []*Player
	DB.Find(&players)
	return players
}

func FindUserByName(name string) Player {
	user := Player{}
	DB.Where("user_name=?", name).Find(&user)
	return user
}

// 生成会员ID
func CreatePlayerId(*gorm.DB) int {
	player := Player{}
	for {
		randomNum := rand.Intn(900000) + 100000
		res := DB.Where("player_id=?", randomNum).Find(&player)
		if res.Error != nil {
			return randomNum
		}
	}

}
