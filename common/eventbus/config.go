/*
 * @Author: kcersing wt4@live.cn
 * @Date: 2026-01-22 17:26:00
 * @LastEditors: kcersing wt4@live.cn
 * @LastEditTime: 2026-01-22 17:26:06
 * @FilePath: \api\common\eventbus\config.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package eventbus

import "time"

// Config 可配置项，用于调整池和处理器行为
type Config struct {
	QueueSize      int           // ConsumerPool 队列大小
	HandlerTimeout time.Duration // 每个 handler 的最大处理时长（0 表示无超时）
}

// DefaultConfig 默认配置
var DefaultConfig = &Config{
	QueueSize:      1000,
	HandlerTimeout: 0,
}
