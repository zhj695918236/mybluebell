package main

import (
	"fmt"
	"mybluebell/controller"
	"mybluebell/dao/mysql"
	"mybluebell/dao/redis"
	"mybluebell/logger"
	"mybluebell/pkg/snowflake"
	"mybluebell/router"
	"mybluebell/setting"
)

type user struct {
	name string
	age  int
}

//var logger *zap.Logger

func main() {
	//if len(os.Args) < 2 {
	//	fmt.Println("need config file.eg: bluebell config.yaml")
	//	return
	//}
	// 加载配置
	if err := setting.Init("./conf/config.yaml"); err != nil {
		fmt.Printf("load config failed, err:%v\n", err)
		return
	}
	if err := logger.Init(setting.Conf.LogConfig, setting.Conf.Mode); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}
	if err := mysql.Init(setting.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	defer mysql.Close() // 程序退出关闭数据库连接
	if err := redis.Init(setting.Conf.RedisConfig); err != nil {
		fmt.Printf("init redis failed, err:%v\n", err)
		return
	}
	defer redis.Close()

	if err := snowflake.Init(setting.Conf.StartTime, setting.Conf.MachineID); err != nil {
		fmt.Printf("init snowflake failed, err:%v\n", err)
		return
	}

	// 初始化gin框架内置的校验器使用的翻译器
	if err := controller.InitTrans("zh"); err != nil {
		fmt.Printf("init validator trans failed, err:%v\n", err)
		return
	}

	// 注册路由
	r := router.SetupRouter(setting.Conf.Mode)
	err := r.Run(fmt.Sprintf(":%d", setting.Conf.Port))
	if err != nil {
		fmt.Printf("run server failed, err:%v\n", err)
		return
	}

	//InitLogger()
	//defer logger.Sync()
	//simpleHttpGet("www.google.com")
	//simpleHttpGet("http://www.baidu.com")
}

//func InitLogger() {
//	logger, _ = zap.NewProduction()
//}

//func simpleHttpGet(url string) {
//	resp, err := http.Get(url)
//	if err != nil {
//		logger.Error(
//			"Error fetching url..",
//			zap.String("url", url),
//			zap.Error(err))
//	} else {
//		logger.Info("Success..",
//			zap.String("statusCode", resp.Status),
//			zap.String("url", url))
//		resp.Body.Close()
//	}
//}
