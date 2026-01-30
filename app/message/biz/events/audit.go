package events

import "context"

// auditKey 是用于在 context 中存储审计条目的私有 key
type auditKey struct{}

// AuditEntry 定义了审计日志的结构
type AuditEntry struct {
	Status  string         // "Success" 或 "Failure"
	Error   string         // 如果失败，记录错误信息
	Details map[string]any // 业务处理器填充的具体细节
}

// NewAuditEntry 创建一个新的审计条目
func NewAuditEntry() *AuditEntry {
	return &AuditEntry{
		Details: make(map[string]any),
	}
}

// WithAuditEntry 将一个新的 AuditEntry 放入 context 中
func WithAuditEntry(ctx context.Context, entry *AuditEntry) context.Context {
	return context.WithValue(ctx, auditKey{}, entry)
}

// GetAuditEntry 从 context 中获取 AuditEntry
// 处理器将使用此函数来获取并填充审计信息
func GetAuditEntry(ctx context.Context) (*AuditEntry, bool) {
	entry, ok := ctx.Value(auditKey{}).(*AuditEntry)
	return entry, ok
}
