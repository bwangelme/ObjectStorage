package heartbeat

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/bwangelme/ObjectStorage/conf"
	"github.com/bwangelme/ObjectStorage/rabbitmq"
)

type DataNodes struct {
	DataNodesMap map[string]time.Time
	mutex        *sync.Mutex
}

func NewDataNodes() *DataNodes {
	return &DataNodes{
		DataNodesMap: make(map[string]time.Time),
		mutex:        new(sync.Mutex),
	}
}

func (d *DataNodes) Add(serverNode string) {
	d.mutex.Lock()
	d.DataNodesMap[serverNode] = time.Now()
	d.mutex.Unlock()
}

func (d *DataNodes) RemoveExpiredNodes() {
	d.mutex.Lock()
	for s, t := range d.DataNodesMap {
		if t.Add(10 * time.Second).Before(time.Now()) {
			delete(d.DataNodesMap, s)
		}
	}
	d.mutex.Unlock()
}

func (d *DataNodes) All() []string {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	ds := make([]string, 0)
	for s, _ := range d.DataNodesMap {
		ds = append(ds, s)
	}
	return ds
}

var defaultDataNodes *DataNodes

func init() {
	defaultDataNodes = NewDataNodes()
}

//ListenHeartBeat 监听 RabbitMQ 的消息，将在线的数据节点及时更新到 defaultDataNodes 中
func ListenHeartBeat() {
	q := rabbitmq.New(conf.RabbitMQServer)
	defer q.Close()

	q.Bind(conf.ApiServersExchange)
	c := q.Consume()
	go removeExpiredDataNodes()
	for msg := range c {
		dataNode, e := strconv.Unquote(string(msg.Body))
		if e != nil {
			//TODO 加入日志框架，将日志分等级输出
			log.Printf("Parse HeartBeat msg, %s\n", e)
			continue
		}
		defaultDataNodes.Add(dataNode)
	}
}

//ChooseRandomDataNode 随机选择一个在线的数据节点，用于做存储对象的节点
func ChooseRandomDataNode() string {
	ds := defaultDataNodes.All()
	n := len(ds)
	if n == 0 {
		return ""
	}
	return ds[rand.Intn(n)]
}

//removeExpiredDataNodes 删除掉10秒未刷新的数据节点
func removeExpiredDataNodes() {
	for {
		time.Sleep(5 * time.Second)
		defaultDataNodes.RemoveExpiredNodes()
	}
}

func Handler(w http.ResponseWriter, r *http.Request) {
	m := r.Method
	if m != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	nodes := defaultDataNodes.All()
	res := map[string]interface{}{
		"nodes": nodes,
	}
	b, _ := json.Marshal(res)
	w.Write(b)
}
