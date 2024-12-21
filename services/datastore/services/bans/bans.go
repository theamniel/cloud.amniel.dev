package bans

import "gorm.io/gorm"

type Ban struct {
	IP     string `gorm:"primaryKey" json:"ip"`
	Reason string `json:"reason"`
}

type BansService struct {
	db *gorm.DB
}

func NewBansService(db *gorm.DB) *BansService {
	return &BansService{db: db}
}

func (s *BansService) Ban(ip, reason string) error {
	return s.db.Create(&Ban{IP: ip, Reason: reason}).Error
}

func (s *BansService) Unban(ip string) error {
	return s.db.Unscoped().Delete(&Ban{}, ip).Error
}

func (s *BansService) IsBanned(ip string) (bool, error) {
	var count int64
	err := s.db.Model(&Ban{}).Select("active").Where("ip = ?", ip).Count(&count).Error
	return count > 0, err
}
