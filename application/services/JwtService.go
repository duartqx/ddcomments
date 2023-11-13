package services

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	m "github.com/duartqx/ddcomments/domains/models"
	r "github.com/duartqx/ddcomments/domains/repositories"
)

type claimsUser struct {
	Id    uuid.UUID
	Email string
	Name  string
}

type JwtAuthService struct {
	userRepository    r.IUserRepository
	sessionRepository r.ISessionRepository
	secret            *[]byte
}

func NewJwtAuthService(
	userRepository r.IUserRepository, sessionStore r.ISessionRepository, secret *[]byte,
) *JwtAuthService {
	return &JwtAuthService{
		userRepository:    userRepository,
		sessionRepository: sessionStore,
		secret:            secret,
	}
}

func (jas JwtAuthService) keyFunc(t *jwt.Token) (interface{}, error) {
	return *jas.secret, nil
}

func (jas JwtAuthService) generateToken(user *claimsUser, expiresAt time.Time) (string, error) {

	claims := jwt.MapClaims{
		"user": map[string]interface{}{
			"id":    user.Id,
			"email": user.Email,
			"name":  user.Name,
		},
		"exp": expiresAt.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err := token.SignedString(*jas.secret)
	if err != nil {
		return "", fmt.Errorf("Bad secret key")
	}

	return tokenStr, nil
}

func (jas JwtAuthService) getUnparsedToken(authorization string, cookie *http.Cookie) string {
	if authorization != "" {
		token, found := strings.CutPrefix(authorization, "Bearer ")
		if found {
			return token
		}
	}
	if cookie != nil {
		return cookie.Value
	}
	return ""
}

func (jas JwtAuthService) ValidateAuth(authorization string, cookie *http.Cookie) (interface{}, error) {

	unparsedToken := jas.getUnparsedToken(authorization, cookie)
	if unparsedToken == "" {
		return nil, fmt.Errorf("Missing Token")
	}

	if _, err := jas.sessionRepository.Get(unparsedToken); err != nil {
		return nil, fmt.Errorf("Missing session")
	}

	parsedToken, err := jwt.Parse(unparsedToken, jas.keyFunc)
	if err != nil || !parsedToken.Valid {

		jas.sessionRepository.Delete(unparsedToken)

		return nil, fmt.Errorf("Expired session")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("Could not parse claims")
	}

	return claims["user"], nil
}

func (jas JwtAuthService) Login(user m.User) (token string, expiresAt time.Time, err error) {

	if user.GetEmail() == "" || user.GetPassword() == "" {
		return token, expiresAt, fmt.Errorf("Invalid Email or Password")
	}

	dbUser, err := jas.userRepository.FindByEmail(user.GetEmail())
	if err != nil {
		return token, expiresAt, fmt.Errorf("Invalid Email")
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(dbUser.GetPassword()), []byte(user.GetPassword()),
	); err != nil {
		return token, expiresAt, fmt.Errorf("Invalid Password")
	}

	createdAt := time.Now()
	expiresAt = createdAt.Add(time.Hour * 12)

	token, err = jas.generateToken(
		&claimsUser{
			Id:    dbUser.GetId(),
			Email: dbUser.GetEmail(),
			Name:  dbUser.GetName(),
		},
		expiresAt,
	)
	if err != nil {
		return "", expiresAt, fmt.Errorf("Could not generate token")
	}

	if err := jas.sessionRepository.Set(token, createdAt, dbUser.GetId()); err != nil {
		return "", expiresAt, err
	}

	return token, expiresAt, nil
}

func (jas *JwtAuthService) Logout(authorization string, cookie *http.Cookie) error {
	unparsedToken := jas.getUnparsedToken(authorization, cookie)
	if unparsedToken == "" {
		return fmt.Errorf("Missing Token")
	}

	if err := jas.sessionRepository.Delete(unparsedToken); err != nil {
		return fmt.Errorf("Invalid Token")
	}

	return nil
}
