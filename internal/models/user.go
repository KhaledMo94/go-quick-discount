package models

import (
	"time"

	"github.com/KhaledMo94/quick-discount/internal/types"
	"gorm.io/gorm"
)

type User struct {
	Name                             string          `json:"name" gorm:"type:varchar(255);not null"`
	Email                            *string         `json:"email" gorm:"uniqueIndex;type:varchar(255)"`
	EmailVerifiedAt                  *time.Time      `json:"email_verified_at"`
	Password                         string          `json:"-" gorm:"not null"`
	RememberToken                    *string         `json:"-"`
	PhoneNumber                      *string         `json:"phone_number" gorm:"uniqueIndex"`
	CountryCode                      *string         `json:"country_code"`
	PhoneVerifiedAt                  *time.Time      `json:"phone_verified_at"`
	types.HasSex                     `gorm:"embedded"`
	types.HasStatus                  `gorm:"embedded"`
	Image                            types.Image     `json:"image" gorm:"embedded"`
	FcmToken                         *string         `json:"fcm_token"`
	QrCode                           *string         `json:"qr_code"`
	CityID                           *int64          `json:"city_id"`
	CardCode                         *string         `json:"card_code"`
	OtpCode                          *string         `json:"otp_code"`
	OtpExpiresAt                     *time.Time      `json:"otp_expires_at"`
	ServiceProviderModeratorID       *int64          `json:"service_provider_moderator_id"`
	ServiceProviderBranchID          *int64          `json:"service_provider_branch_id"`
	ServiceProviderBranchModeratorID *int64          `json:"service_provider_branch_moderator_id"`
	BirthDate                        *time.Time      `json:"birth_date"`
	Points                           int             `json:"points" gorm:"default:0"`
	ReferredID                       *int64          `json:"referred_id"`

	BaseModel

	//relations
    Roles 				[]Role `json:"roles" gorm:"many2many:model_has_roles;foreignKey:ID;joinForeignKey:ModelID;references:ID;joinReferences:RoleID;"`
	Permissions 		[]Permission `json:"permissions" gorm:"many2many:model_has_permissions;foreignKey:ID;joinForeignKey:ModelID;references:ID;joinReferences:PermissionID;"`
	City				*City			`json:"city" gorm:"foreignKey:CityID"`
}

//------------BEGAIN SCOPES ---------------------------

func ScopeVerified(db *gorm.DB) *gorm.DB {
	return db.Where(func(tx *gorm.DB) *gorm.DB {
		return tx.Where("email_verified_at is not null").Or("phone_verified_at is not null")
	})
}

func ScopeEndUser(db *gorm.DB) *gorm.DB {
	return db.
		Where("service_provider_moderator_id is null").
		Where("service_provider_branch_id is null").
		Where("service_provider_branch_moderator_id is null")
}

func ScopeCashier(db *gorm.DB) *gorm.DB {
	return db.
		Where("service_provider_branch_id is not null")
}

func ScopeBranchModerator(db *gorm.DB) *gorm.DB {
	return db.
		Where("service_provider_branch_moderator_id is not null")
}

func ScopeProviderModerator(db *gorm.DB) *gorm.DB {
	return db.
		Where("service_provider_moderator_id is not null")
}

func ScopeCanPushNotification(db *gorm.DB) *gorm.DB {
	return db.
		Where("fcm_token is not null")
}

func ScopeCanEmail(db *gorm.DB) *gorm.DB {
	return db.
		Where("email is not null")
}

//-----------------END SCOPES--------------------

//-----------------BEGAIN ATTRIBUTES ------------

func (u *User) Phone() string {
	if u.CountryCode == nil || u.PhoneNumber == nil {
		return ""
	}
	return "+" + *u.CountryCode + *u.PhoneNumber
}

//-----------------END ATTRIBUTES ------------

//---------------- BEGAIN RELATIONS ----------