package models

type URL struct {
	LotURL      string
	DocumentULR string
}

type Filter struct {
}

type Result struct {
	ID                      int     `json:"id"`
	Number                  string  `json:"number"`
	NameRu                  string  `json:"nameRu"`
	NameKk                  string  `json:"nameKk"`
	NameEn                  *string `json:"nameEn"`
	TenderType              string  `json:"tenderType"`
	SumTruNoNds             float64 `json:"sumTruNoNds"`
	AcceptanceBeginDateTime string  `json:"acceptanceBeginDateTime"`
	AcceptanceEndDateTime   string  `json:"acceptanceEndDateTime"`
	AdvertStatus            string  `json:"advertStatus"`
	FlagApplicationFiled    bool    `json:"flagApplicationFiled"`
	FlagNegotiationOutside  *bool   `json:"flagNegotiationOutside"`
}
