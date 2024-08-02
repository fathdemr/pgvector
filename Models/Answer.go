package Models

type Answer struct {
	ID     uint64 `gorm:"primary_key" json:"id"`
	Answer string `json:"answer"`
}
