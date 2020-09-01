package conf

import (
	"log"
	"os"
	"path"
	"strings"
)

var (
	// APIServersExchange 维护数据节点信息的 Exchange
	APIServersExchange = "api_servers"
	// DataServersExchange 从数据节点定位文件的 Exchange
	DataServersExchange = "data_servers"
)

// 从外部环境变量读取的配置
var (
	ListenAddress  string
	StorageRoot    string
	RabbitMQServer string
	IsDataNode     bool
)

func init() {
	ListenAddress = os.Getenv("LISTEN_ADDRESS")
	ErrOnEmpty(ListenAddress, "LISTEN_ADDRESS")

	StorageRoot = os.Getenv("STORAGE_ROOT")
	ErrOnEmpty(StorageRoot, "STORAGE_ROOT")

	RabbitMQServer = os.Getenv("RABBITMQ_SERVER")
	ErrOnEmpty(RabbitMQServer, "RABBITMQ_SERVER")

	IsDataNodeStr := os.Getenv("IS_DATA_NODE")
	ErrOnEmpty(IsDataNodeStr, "IS_DATA_NODE")
	IsDataNode = strings.ToLower(IsDataNodeStr) == "true"

	createStoreRoot()
}

// ErrOnEmpty 目标值为 nil 时，程序退出，初始化失败
func ErrOnEmpty(envVal, envName string) {
	if envVal == "" {
		log.Fatalf("Need To Set %s\n", envName)
	}
}

// createStoreRoot 创建存储文件的文件夹
func createStoreRoot() {
	root := path.Join(StorageRoot, "objects")
	err := os.MkdirAll(root, 0755)
	if err != nil {
		log.Println("Create %s failed: %s", root, err)
	}
}
