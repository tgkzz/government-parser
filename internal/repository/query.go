package repository

import "parser/internal/models"

// Update
func (r *Repository) UploadOne(lot models.Result) error {
	query := `
INSERT INTO results (
    ID, Number, NameRu, NameKk, NameEn, TenderType, SumTruNoNds, 
    AcceptanceBeginDateTime, AcceptanceEndDateTime, AdvertStatus, 
    FlagApplicationFiled, FlagNegotiationOutside
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
`
	_, err := r.db.Exec(query,
		lot.ID, lot.Number, lot.NameRu, lot.NameKk, lot.NameEn, lot.TenderType,
		lot.SumTruNoNds, lot.AcceptanceBeginDateTime, lot.AcceptanceEndDateTime,
		lot.AdvertStatus, lot.FlagApplicationFiled, lot.FlagNegotiationOutside)
	if err != nil {
		return err
	}

	return nil
}
