package utils

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB
var RDB *redis.Client

const PublishKey = "websocket"

func InitConfig() {
	viper.SetConfigName("index")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("index:", viper.Get("index"))
	fmt.Println("mysql:", viper.Get("mysql"))
}

func InitMysql() {
	newLogger := logger.New(

		//自定义打印SQL语句
		// 创建一个新的 log.Logger 实例，将日志输出到标准输出（os.Stdout）。
		//使用 \r\n 作为日志的换行符。
		// log.LstdFlags 设置了日志的格式，包括日期和时间。
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second, // 慢查询的阈值，超过这个时间将被记录
			LogLevel:                  logger.Info, // 日志级别，这里设置为 Silent 表示不输出日志
			IgnoreRecordNotFoundError: true,        // 是否忽略 ErrRecordNotFound 错误
			ParameterizedQueries:      true,        // 是否在 SQL 日志中包含参数
			Colorful:                  true,        // 是否启用彩色日志
		},
	)

	var err error
	DB, err = gorm.Open(mysql.Open(viper.GetString("mysql.dns")), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		fmt.Println("数据库连接失败", err)
	} else {
		fmt.Println("数据库连接成功")
	}

}

func InitRedis() {
	RDB = redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.addr"),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.db"),
	})

	_, err := RDB.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println("redis启动失败：", err)
	} else {
		fmt.Println("redis启动成功")
	}

}

// 发布消息
func RedisPublish(ctx context.Context, channel string, msg string) error {
	err := RDB.Publish(ctx, channel, msg).Err()
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("消息发布成功", msg)
	return err
}

// 订阅消息
func RedisSubscribe(ctx context.Context, channel string) (string, error) {
	pubsub := RDB.Subscribe(ctx, channel)
	fmt.Println("ctx传入成功", ctx)

	msg, err := pubsub.ReceiveMessage(ctx)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	fmt.Println("订阅成功", msg.Payload)
	return msg.Payload, err

}
