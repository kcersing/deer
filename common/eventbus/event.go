/*
 * @Author: kcersing wt4@live.cn
 * @Date: 2025-12-12 10:21:48
 * @LastEditors: kcersing wt4@live.cn
 * @LastEditTime: 2026-01-16 11:36:55
 * @FilePath: \api\common\eventbus\event.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package eventbus

import (
	"time"

	"github.com/google/uuid"
)

// Event 事件结构体
type Event struct {
	Id        string    // 唯一标识，自动生成UUID
	Topic     string    // 事件主题（如"user_registered"）
	Payload   any       // 事件载体（可以是任何类型）
	Source    string    // 事件来源（"service"或"amqp"）
	Version   int64     // 版本号
	Timestamp time.Time // 时间戳
	Priority  int64     // 优先级
}

func NewEvent(topic string, payload any) *Event {
	return &Event{
		Id:        uuid.New().String(),
		Topic:     topic,
		Payload:   payload,
		Version:   0,
		Timestamp: time.Time{},
		Priority:  0,
	}
}
