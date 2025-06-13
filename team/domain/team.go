package domain

type TeamStatus string

const (
	StatusEnabled  TeamStatus = "enabled"
	StatusDisabled TeamStatus = "disabled"
)

type Team struct {
	ID          int        `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string     `json:"name"`
	Avatar      string     `json:"avatar"`
	Designation string     `json:"designation"`
	Bio         string     `json:"bio"`
	SortOrder   int        `json:"sort_order"`
	Status      TeamStatus `json:"status"`
}
