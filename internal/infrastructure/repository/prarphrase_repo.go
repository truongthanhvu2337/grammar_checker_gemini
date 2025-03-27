package repository

import (
	"database/sql"
	"grammar_checker/internal/domain"
)

type ParaphraseRepo struct {
	DB *sql.DB
}

func (r *ParaphraseRepo) Create(item *domain.Paraphrase) (int64, error) {
	query := "INSERT INTO paraphrases (OriginalText, ParaphrasedText, CreatedAt) VALUES (?, ?)"
	result, err := r.DB.Exec(query, item.OriginalText, item.ParaphrasedText, item.CreatedAt)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}


func (r *ParaphraseRepo) GetById(id int64) (*domain.ParaphraseRepo, error) {
	query := "SELECT * FROM grammar_checks WHERE id = ?"
	row := r.DB.QueryRow(query, id)

	var item domain.GrammarCheck
	err := row.Scan(&item.ID, &item.OriginalText, &item.CorrectedText, &item.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, errors.New("record not found")
	} else if err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *ParaphraseRepo) Update(grammarCheck *domain.ParaphraseRepo) error {
	query := "UPDATE grammar_checks SET OriginalText = ?, CorrectedText = ? WHERE id = ?"
	_, err := r.DB.Exec(query, grammarCheck.OriginalText, grammarCheck.CorrectedText, grammarCheck.ID)
	return err
}

func (r *ParaphraseRepo) Delete(id int64) error {
	query := "DELETE FROM grammar_checks WHERE id = ?"
	_, err := r.DB.Exec(query, id)
	return err
}

func (r * ParaphraseRepo) GetAll() ([]domain.ParaphraseRepo, error){
	query := "SELECT * FROM grammar_checks"
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var grammarChecks []domain.GrammarCheck
	for rows.Next() {
		var item domain.GrammarCheck
		err := rows.Scan(&item.ID, &item.OriginalText, &item.CorrectedText, &item.CreatedAt)
		if err != nil {
			return nil, err
		}
		grammarChecks = append(grammarChecks, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return grammarChecks, nil

}