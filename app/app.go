package app

import (
	"database/sql"
	"flag"
	"fmt"
	"kitstart_goweb/app/router"
	"kitstart_goweb/app/utils"
	"sync"
	"time"

	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttprouter"

	// Register mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// 参数值
var configDir string // 配置文件路径
var env string       // 运行环境 prod | dev

// App -
type App struct {
	CommandArgs map[string]interface{}
	Db          *sql.DB // 数据库连接对象
	Port        string
}

var app *App // 系统app单例

var once sync.Once

// InitApp - 初始化App 整个项目  / 线程安全 可以多次用
// 加载配置
// 初始化路由表
func InitApp() *App {

	once.Do(func() {
		app = &App{
			CommandArgs: make(map[string]interface{}),
		}
		// 参数配置
		flag.StringVar(&configDir, "c", "", "config directory")
		flag.StringVar(&env, "e", "dev", "enviroment of running")

		app.CommandArgs["config"] = configDir
		app.CommandArgs["env"] = env

		flag.Parse()

		if configDir == "" {
			fmt.Println("config directory is required to start app")
			return
		}

		err := utils.LoadConfig(configDir)

		if err != nil {
			fmt.Println(err)
		}

		port, err := utils.ConfigString("app.port")
		app.Port = port

		// 初始化数据库
		InitDb(app)

		// 初始化web
		InitWeb(app)
	})

	return app
}

// InitDb - 初始化database
func InitDb(app *App) (db *sql.DB) {
	env := app.CommandArgs["env"].(string)

	baseConfigKey := "database." + env
	driverName, err := utils.ConfigString(baseConfigKey + ".driver")
	if err != nil {
		utils.Error(err)
		return
	}
	if driverName == "mysql" {
		name, err := utils.ConfigString(baseConfigKey + ".user")
		if err != nil {
			utils.Error(err)
			return
		}
		password, err := utils.ConfigString(baseConfigKey + ".password")
		if err != nil {
			utils.Error(err)
			return
		}

		if name == "" || password == "" {
			utils.Info("数据库账户名或密码未配置")
			return
		}

		host, err := utils.ConfigString(baseConfigKey + ".host")
		if err != nil {
			utils.Error(err)
			return
		}

		dbName, err := utils.ConfigString(baseConfigKey + ".db")
		if err != nil {
			utils.Error(err)
			return
		}
		charset, _ := utils.ConfigString(baseConfigKey + ".charset")
		if charset == "" {
			charset = "utf8mb4"
		}

		port, _ := utils.ConfigString(baseConfigKey + ".port")

		dsn := name + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbName + "?charset=" + charset
		db, err = sql.Open(driverName, dsn)
		if err != nil {
			utils.Error(err)
			return
		}

		db.SetConnMaxIdleTime(time.Minute * 10) // 10分钟的空闲断开时间
		db.SetConnMaxLifetime(time.Minute * 3)  // 3分钟的连接时间
		db.SetMaxIdleConns(20)                  // 20个最大闲置连接
		db.SetMaxOpenConns(50)                  // 50个最大可连接数

		app.Db = db // 把 db 放到 app下

	}

	return

}

// InitWeb - 初始化web 组件
func InitWeb(app *App) {
	routerInstance := fasthttprouter.New()
	handler := router.InitRoutes(routerInstance)

	go fasthttp.ListenAndServe(":"+app.Port, handler)
}
