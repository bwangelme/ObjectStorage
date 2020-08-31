package conf

import (
	"log"
	"os"
	"strings"
)

var (
	ApiServersExchange  = "api_servers"
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

	IsDataNode = strings.ToLower(os.Getenv("IS_DATA_NODE")) == "true"
}

func ErrOnEmpty(envVal, envName string) {
	if envVal == "" {
		log.Fatalf("Need To Set %s\n", envName)
	}
}
