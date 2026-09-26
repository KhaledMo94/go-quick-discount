package helper

import (
	"crypto/rand"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/KhaledMo94/quick-discount/internal/models"
)

var (
	ErrInvalidSanctumToken = errors.New("invalid sanctum token")
	ErrInvalidToken        = errors.New("invalid token")
	ErrTokenExpired        = errors.New("token expired")
)

// GenerateSecret creates a cryptographically random hex secret.
func GenerateSecret() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

// HashToken returns the SHA-256 hex digest of a plain token secret.
func HashToken(plain string) string {
	sum := sha256.Sum256([]byte(plain))
	return hex.EncodeToString(sum[:])
}

// FormatBearerToken builds a Sanctum-style "id|plain" bearer token.
func FormatBearerToken(id int64, plain string) string {
	return fmt.Sprintf("%d|%s", id, plain)
}

// SplitSanctumToken parses a bearer token into its id and plain secret.
func SplitSanctumToken(token string) (int64, string, error) {
	parts := strings.SplitN(token, "|", 2)
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		return 0, "", ErrInvalidSanctumToken
	}

	id, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		return 0, "", ErrInvalidSanctumToken
	}

	return id, parts[1], nil
}

// TokenHashMatches compares a stored hash with the hash of a plain secret
// using constant-time comparison.
func TokenHashMatches(storedHash, plain string) bool {
	return subtle.ConstantTimeCompare([]byte(storedHash), []byte(HashToken(plain))) == 1
}

// IsTokenValid reports whether the token has not expired.
func IsTokenValid(token *models.PersonalAccessToken) bool {
	if token == nil {
		return false
	}
	if token.ExpiresAt != nil && token.ExpiresAt.Before(time.Now()) {
		return false
	}
	return true
}
