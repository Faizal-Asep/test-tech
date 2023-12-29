package repositorie

import (
	"context"
	"log"
	"time"

	"gorm.io/gorm"
)

type handPhoneClient interface {
	Get(ctx context.Context) ([]HandPhone, error)
	GetById(ctx context.Context, id string) (HandPhone, error)
	Save(ctx context.Context, in HandPhone) (HandPhone, error)
	Update(ctx context.Context, id string, in HandPhone) (HandPhone, error)
	Delete(ctx context.Context, id string) error
}

func NewHandPhone() handPhoneClient {
	return &HandPhone{}
}

// Table Structur
type HandPhone struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Number    string         `gorm:"column:phonenumber;size:25;not null" json:"phonenumber" binding:"required,numeric"`
	Provider  string         `gorm:"column:provider;size:125;not null" json:"provider" binding:"required"`
}

func (HandPhone) TableName() string {
	return "hand_phone"
}

func (HandPhone) Get(ctx context.Context) (out []HandPhone, err error) {

	if err = db.Find(&out).Error; err != nil {
		log.Println(err)
		return
	}
	return
}

func (HandPhone) GetById(ctx context.Context, id string) (out HandPhone, err error) {

	// out = &HandPhone{}
	if err = db.WithContext(ctx).First(&out, "id = ?", id).Error; err != nil {
		log.Println(err)
		return
	}
	return
}

func (HandPhone) Save(ctx context.Context, in HandPhone) (out HandPhone, err error) {

	if err = db.WithContext(ctx).Create(&in).Error; err != nil {
		log.Println(err)
		return
	}
	out = in
	return
}

func (HandPhone) Update(ctx context.Context, id string, in HandPhone) (out HandPhone, err error) {

	if err = db.WithContext(ctx).First(&out, "id = ?", id).Error; err != nil {
		log.Println(err)
		return
	}

	if in.Number != "" {
		out.Number = in.Number
	}
	if in.Provider != "" {
		out.Provider = in.Provider
	}

	if err = db.WithContext(ctx).Save(&out).Error; err != nil {
		log.Println(err)
		return
	}

	return
}

func (HandPhone) Delete(ctx context.Context, id string) (err error) {
	if err = db.WithContext(ctx).Delete(&HandPhone{}, id).Error; err != nil {
		log.Println(err)
		return
	}
	return
}
