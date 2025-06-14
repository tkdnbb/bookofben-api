package models

import "time"

// Verse represents a single Bible verse
type Verse struct {
	BookID        string `json:"book_id"`
	BookName      string `json:"book_name"`
	Chapter       int    `json:"chapter"`
	Verse         int    `json:"verse"`
	Text          string `json:"text"`
	TranslationID string `json:"translation_id"` // 默认为 "en"
}

// BibleResponse represents the API response for Bible passages
type BibleResponse struct {
	Reference       string  `json:"reference"`
	Verses          []Verse `json:"verses"`
	Text            string  `json:"text"`
	TranslationID   string  `json:"translation_id"`
	TranslationName string  `json:"translation_name"`
	TranslationNote string  `json:"translation_note"`
}

// Translation represents a Bible translation
type Translation struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Note string `json:"note"`
}

// Book represents a Bible book
type Book struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Chapters int    `json:"chapters"`
}
type Comment struct {
	ID            string     `json:"id"`
	Title         string     `json:"title"`
	Content       string     `json:"content"`
	BookID        string     `json:"book_id"` // 关联到具体的书卷
	Chapter       int        `json:"chapter"`
	Verse         int        `json:"verse"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	PinnedAmount  int64      `json:"pinned_amount"` // 置顶金额 (以分为单位，避免浮点数问题)
	PinnedUntil   *time.Time `json:"pinned_until"`  // 置顶到期时间
	IsActive      bool       `json:"is_active"`     // 评论是否有效/显示
	UserID        string     `json:"user_id"`
	Username      string     `json:"username"` // 用户名，便于显示
	TranslationID string     `json:"translation_id"`
	TransactionID string     `json:"transaction_id"` // 关联的置顶支付交易
}

type Transaction struct {
	ID          string     `json:"id"`
	Sender      string     `json:"sender"`       // 发送方地址
	Recipient   string     `json:"recipient"`    // 接收方地址
	Amount      float64    `json:"amount"`       // 转账金额 (USDT)
	TxHash      string     `json:"tx_hash"`      // 区块链交易哈希
	Network     string     `json:"network"`      // "TRC20" 或 "ERC20"
	Status      string     `json:"status"`       // "pending", "confirmed", "failed"
	BlockNumber int64      `json:"block_number"` // 区块高度
	GasFee      float64    `json:"gas_fee"`      // 手续费
	Memo        string     `json:"memo"`         // 备注信息
	CreatedAt   time.Time  `json:"created_at"`
	CompletedAt *time.Time `json:"completed_at"` // 使用指针，因为可能为空
}
