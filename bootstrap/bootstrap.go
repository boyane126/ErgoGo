// 按顺序安装组件
package bootstrap

func Setup() {
	// 初始化 Logger
	SetupLogger()
	// 初始化数据库
	SetupDB()
}
