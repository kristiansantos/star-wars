package hash

import "golang.org/x/crypto/bcrypt"

type PasswordHash struct{}

func New() *PasswordHash {
	return &PasswordHash{}
}

func (hash *PasswordHash) Create(payload string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(payload), 14)
	return string(bytes)
}

func (hash *PasswordHash) Compare(hashed, payload string) bool {
	check := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(payload))
	return check == nil
}
