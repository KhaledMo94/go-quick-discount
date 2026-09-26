package interfaces

import (
	"time"

	"github.com/KhaledMo94/quick-discount/internal/models"
)

type TokenIssuer interface{
	Create(userID int64, name string, expiresAt *time.Time) (string, error)
}

type TokenResolver interface{
	FindByToken(token string) (*models.PersonalAccessToken, error)
}

type TokenRevoker interface{
	DeleteToken(token string) error
	DeleteTokenById(tokenID int64) error
	DeleteAllUserTokens(userID int64) error
	LogOutFromAllDevicesExceptCurrent(token string) error
}

type PersonalAccessTokenRepository interface{
	TokenIssuer
	TokenResolver
	TokenRevoker
	FindByID(id int64) (*models.PersonalAccessToken, error)
}