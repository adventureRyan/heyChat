package utils

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB  *gorm.DB
	Red *redis.Client
)

func InitConfig() {
	viper.SetConfigName("app")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("config.app:", viper.Get("app"))
	fmt.Println("config app inited")
}

func InitMySQL() {
	// 自定义日志模板,打印SQL语句
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, // 慢SQL阈值
			LogLevel:      logger.Info, // 级别
			Colorful:      true,        // 彩色
		},
	)
	DB, _ = gorm.Open(mysql.Open(viper.GetString("mysql.dns")), &gorm.Config{Logger: newLogger})
	fmt.Println("config.mysql:", viper.Get("mysql"))
	fmt.Println("Mysql inited")
	//user := models.UserBasic{}
	//DB.Find(&user)
	//fmt.Println(user)
}

func InitRedis() {
	Red = redis.NewClient(&redis.Options{
		Addr:         viper.GetString("redis.addr"),
		Password:     viper.GetString("redis.password"),
		DB:           viper.GetInt("redis.DB"),
		PoolSize:     viper.GetInt("redis.PoolSize"),
		MinIdleConns: viper.GetInt("redis.MinIdleConns"),
	})
	// ctx := context.Background()
	// pong, err := Red.Ping(ctx).Result()
	// if err != nil {
	// 	fmt.Println("init redis .", err)
	// } else {
	// 	fmt.Println("init redis success", pong)
	// }
}

const (
	PublishKey = "websocket"
)

// Publish发布消息到Redis
func Publish(ctx context.Context, channel string, msg string) error {
	var err error
	err = Red.Publish(ctx, channel, msg).Err()
	return err
}

// Subscribe订阅消息到Redis
func Subscribe(ctx context.Context, channel string) (string, error) {
	sub := Red.Subscribe(ctx, channel)
	fmt.Println("Subscribe", ctx)
	msg, err := sub.ReceiveMessage(ctx)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	fmt.Println("Subscribe success:", msg.Payload)
	return msg.Payload, err
}
