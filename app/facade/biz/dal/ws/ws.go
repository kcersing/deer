package ws

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/json"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/hertz-contrib/websocket"
	"hardware/biz/dal/mq"
	"net/http"
	"sync/atomic"
	"time"
)

// 全局指标变量
var (
	activeConnections  int32
	totalConnections   int64
	messagesReceived   int64
	messagesSent       int64
	bytesReceived      int64
	bytesSent          int64
	lastConnectionTime int64 // Unix 时间戳
)
var upgrader = websocket.HertzUpgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(ctx *app.RequestContext) bool {
		return true
	},
}

func Handler(sub mq.Subscriber) app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		err := upgrader.Upgrade(ctx, func(conn *websocket.Conn) {

			// 更新连接计数
			atomic.AddInt32(&activeConnections, 1)
			atomic.AddInt64(&totalConnections, 1)
			atomic.StoreInt64(&lastConnectionTime, time.Now().Unix())
			defer atomic.AddInt32(&activeConnections, -1)

			msgs, cleanUp, err := sub.Subscribe(c)
			defer cleanUp()
			if err != nil {
				klog.Error("cannot subscribe", err)
				ctx.String(http.StatusInternalServerError, "")
				return
			}
			done := make(chan struct{})
			go func() {
				for {
					messageType, message, err := conn.ReadMessage()
					if err != nil {
						if !websocket.IsCloseError(err, websocket.CloseGoingAway, websocket.CloseNormalClosure) {
							klog.Warn("unexpected read error", err)
						}
						done <- struct{}{}
						break
					}

					// 更新接收消息指标
					atomic.AddInt64(&messagesReceived, 1)
					atomic.AddInt64(&bytesReceived, int64(len(message)))

					// 处理客户端发送的消息
					klog.Debug("received message from client: ", string(message))

					// 根据 messageType 做不同处理
					if messageType == websocket.TextMessage {

					}

				}
			}()

			for {
				select {
				case msg := <-msgs:
					//if msg == nil {
					//	klog.Warn("received nil message from MQ")
					//	continue
					//}
					err := conn.WriteJSON(msg)
					if err != nil {
						klog.Warn("cannot write JSON %s", err.Error())
					}

					// 序列化为 JSON 以获取字节数
					jsonData, err := json.Marshal(msg)
					if err != nil {
						klog.Warn("cannot marshal message: %s", err.Error())
						continue
					}

					err = conn.WriteMessage(websocket.TextMessage, jsonData)
					if err != nil {
						klog.Warn("cannot write JSON %s", err.Error())
						continue
					}
					// 更新发送消息指标
					atomic.AddInt64(&messagesSent, 1)
					atomic.AddInt64(&bytesSent, int64(len(jsonData)))

				case <-done:
					return
				}
			}
		})
		if err != nil {
			klog.Warnf("Upgrade err: %s", err.Error())
		}

	}
}

//import (
//    "github.com/prometheus/client_golang/prometheus"
//    "github.com/prometheus/client_golang/prometheus/promauto"
//)
//
//// 创建 Prometheus 指标
//var (
//    wsActiveConnections = promauto.NewGauge(prometheus.GaugeOpts{
//        Name: "websocket_active_connections",
//        Help: "当前活跃的 WebSocket 连接数",
//    })
//
//    wsTotalConnections = promauto.NewCounter(prometheus.CounterOpts{
//        Name: "websocket_total_connections",
//        Help: "总 WebSocket 连接数",
//    })
//
//    wsMessagesReceived = promauto.NewCounter(prometheus.CounterOpts{
//        Name: "websocket_messages_received_total",
//        Help: "接收的 WebSocket 消息总数",
//    })
//
//    wsMessagesSent = promauto.NewCounter(prometheus.CounterOpts{
//        Name: "websocket_messages_sent_total",
//        Help: "发送的 WebSocket 消息总数",
//    })
//)
//
//// 然后在 Handler 中更新这些指标
//func Handler(sub mq.Subscriber) app.HandlerFunc {
//    return func(c context.Context, ctx *app.RequestContext) {
//        err := upgrader.Upgrade(ctx, func(conn *websocket.Conn) {
//            // 更新 Prometheus 指标
//            wsActiveConnections.Inc()
//            wsTotalConnections.Inc()
//
//            defer wsActiveConnections.Dec()
//
//            // 其余代码逻辑...
//        })
//
//        // 错误处理...
//    }
//}
//
//// 然后在应用中注册 Prometheus 指标端点
//// h.GET("/metrics", metricsHandler())
