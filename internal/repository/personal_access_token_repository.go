package repository

import (
	"time"

	"github.com/KhaledMo94/quick-discount/internal/db"
	"github.com/KhaledMo94/quick-discount/internal/helper"
	"github.com/KhaledMo94/quick-discount/internal/models"
)

type PersonalAccessTokenRepository struct {
	db *db.GormDB
}

func NewPersonalAccessRepository(db *db.GormDB) *PersonalAccessTokenRepository {
	return &PersonalAccessTokenRepository{db: db}
}

func (r *PersonalAccessTokenRepository) Create(userID int64, name string, expiresAt *time.Time) (string, error) {
	plain, err := helper.GenerateSecret()
	if err != nil {
		return "", err
	}

	row := models.PersonalAccessToken{
		TokenableID:   userID,
		TokenableType: models.TOKENABLE_TYPE_USER,
		Name:          name,
		Token:         helper.HashToken(plain),
		ExpiresAt:     expiresAt,
	}

	if err := r.db.Create(&row).Error; err != nil {
		return "", err
	}

	return helper.FormatBearerToken(row.ID, plain), nil
}

func (r *PersonalAccessTokenRepository) FindByID(id int64) (*models.PersonalAccessToken, error) {
	var token models.PersonalAccessToken

	if err := r.db.Find(&token, id).Error; err != nil {
		return nil, err
	}

	if !helper.IsTokenValid(&token) {
		return nil, helper.ErrTokenExpired
	}

	return &token, nil
}

func (r *PersonalAccessTokenRepository) FindByToken(token string) (*models.PersonalAccessToken, error) {
	row, err := r.resolveBearer(token)
	if err != nil {
		return nil, err
	}

	if !helper.IsTokenValid(row) {
		return nil, helper.ErrTokenExpired
	}

	return row, nil
}

func (r *PersonalAccessTokenRepository) DeleteTokenById(tokenID int64) error {
	return r.db.Delete(&models.PersonalAccessToken{}, tokenID).Error
}

func (r *PersonalAccessTokenRepository) DeleteToken(token string) error {
	row, err := r.resolveBearer(token)
	if err != nil {
		return err
	}
	return r.db.Delete(row).Error
}

func (r *PersonalAccessTokenRepository) DeleteAllUserTokens(userID int64) error {
	return r.db.Where("tokenable_id = ?", userID).
		Where("tokenable_type = ?", models.TOKENABLE_TYPE_USER).
		Delete(&models.PersonalAccessToken{}).Error
}

func (r *PersonalAccessTokenRepository) LogOutFromAllDevicesExceptCurrent(token string) error {
	row, err := r.resolveBearer(token)
	if err != nil {
		return err
	}

	return r.db.
		Where("tokenable_id = ? AND tokenable_type = ?", row.TokenableID, models.TOKENABLE_TYPE_USER).
		Where("id <> ?", row.ID).
		Delete(&models.PersonalAccessToken{}).Error
}

func (r *PersonalAccessTokenRepository) resolveBearer(token string) (*models.PersonalAccessToken, error) {
	id, plain, err := helper.SplitSanctumToken(token)
	if err != nil {
		return nil, err
	}

	var row models.PersonalAccessToken
	if err := r.db.First(&row, id).Error; err != nil {
		return nil, err
	}

	if !helper.TokenHashMatches(row.Token, plain) {
		return nil, helper.ErrInvalidToken
	}

	return &row, nil
}
