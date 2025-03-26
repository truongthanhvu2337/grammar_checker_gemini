type Paraphrase struct {
	ID             int       `json:"id"`
	CheckID        int       `json:"check_id"`         // Liên kết với GrammarCheck
	Level          string    `json:"level"`            // A1, A2, B1, B2, C1, C2 hoặc theo lớp học
	ParaphrasedText string    `json:"paraphrased_text"` // Văn bản được diễn đạt lại
	CreatedAt      time.Time `json:"created_at"`
}
