package domain

import "time"

type GrammarCheck struct {
	ID            int       `json:"id"`
	OriginalText  string    `json:"original_text"`
	CorrectedText string    `json:"corrected_text"`
	Paraphrases   []Paraphrase `json:"paraphrases"`
	GrammarErrors []GrammarError `json:"grammar_errors"`
	CreatedAt     time.Time `json:"created_at"`
}
