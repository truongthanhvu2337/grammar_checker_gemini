
type GrammarError struct {
	ID           int       `json:"id"`
	CheckID      int       `json:"check_id"`      // Liên kết với GrammarCheck
	ErrorText    string    `json:"error_text"`    // Đoạn văn bản có lỗi
	SuggestedFix string    `json:"suggested_fix"` // Gợi ý sửa lỗi
	ErrorType    string    `json:"error_type"`    // Loại lỗi (vd: tense, spelling, punctuation)
	CreatedAt    time.Time `json:"created_at"`
}
