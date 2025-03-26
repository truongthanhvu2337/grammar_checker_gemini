package entity

import "time"

type GrammarCheck struct {
	ID            int       `json:"id"`
	OriginalText  string    `json:"original_text"`
	CorrectedText string    `json:"corrected_text"`
	CreatedAt     time.Time `json:"created_at"`
}
