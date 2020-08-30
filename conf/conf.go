package conf

import (
	"log"
	"os"
)

var (
	ApiServersExchange = "apiServers"
	DataServersExchange = "dataServers"
)

// 从外部环境变量读取的配置
var (
	ListenAddress string
	StorageRoot   string
	RabbitMQServer string
)

func init() {
	ListenAddress = os.Getenv("LISTEN_ADDRESS")
	ErrOnEmpty(ListenAddress, "LISTEN_ADDRESS")

	StorageRoot = os.Getenv("STORAGE_ROOT")
	ErrOnEmpty(StorageRoot, "STORAGE_ROOT")

	RabbitMQServer = os.Getenv("RABBITMQ_SERVER")
	ErrOnEmpty(RabbitMQServer, "RABBITMQ_SERVER")
}

func ErrOnEmpty(envVal, envName string) {
	if envVal == ""	 {
		log.Fatalf("Need To Set %s\n", envName)
	}
}
