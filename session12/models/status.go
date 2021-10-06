package models

type Status struct {
	StatusID  uint `gorm:"primaryKey"`
	StatusTxt string
}

type StatusStorage interface {
	Read() ([]Status, error)
	ReadOneByStatus(txt string) (Status, error)
}
