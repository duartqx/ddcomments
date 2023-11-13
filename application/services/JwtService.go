package services

import (
	"fmt"
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
	userRepository r.IUserRepository
	sessionStore   r.ISessionRepository
	secret         *[]byte
}

func NewJwtAuthService(
	userRepository r.IUserRepository, sessionStore r.ISessionRepository, secret *[]byte,
) *JwtAuthService {
	return &JwtAuthService{
		userRepository: userRepository,
		sessionStore:   sessionStore,
		secret:         secret,
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

func (jas JwtAuthService) getUnparsedToken(authorization, cookie string) string {
	if authorization != "" {
		token, found := strings.CutPrefix(authorization, "Bearer ")
		if found {
			return token
		}
	}
	return cookie
}

func (jas JwtAuthService) ValidateAuth(authorization, cookie string) (interface{}, error) {

	unparsedToken := jas.getUnparsedToken(authorization, cookie)
	if unparsedToken == "" {
		return nil, fmt.Errorf("Missing Token")
	}

	if _, err := jas.sessionStore.Get(unparsedToken); err != nil {
		return nil, fmt.Errorf("Missing session")
	}

	parsedToken, err := jwt.Parse(unparsedToken, jas.keyFunc)
	if err != nil || !parsedToken.Valid {

		jas.sessionStore.Delete(unparsedToken)

		return nil, fmt.Errorf("Expired session")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("Could not parse claims")
	}

	return claims["user"], nil
}

func (jas JwtAuthService) Login(user m.User) (token string, err error) {

	if user.GetEmail() == "" || user.GetPassword() == "" {
		return token, fmt.Errorf("Invalid Email or Password")
	}

	dbUser, err := jas.userRepository.FindByEmail(user.GetEmail())
	if err != nil {
		return token, fmt.Errorf("Invalid Email")
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(dbUser.GetPassword()), []byte(user.GetPassword()),
	); err != nil {
		return token, fmt.Errorf("Invalid Password")
	}

	createdAt := time.Now()
	expiresAt := createdAt.Add(time.Hour * 12)

	token, err = jas.generateToken(
		&claimsUser{
			Id:    dbUser.GetId(),
			Email: dbUser.GetEmail(),
			Name:  dbUser.GetName(),
		},
		expiresAt,
	)
	if err != nil {
		return "", fmt.Errorf("Could not generate token")
	}

	if err := jas.sessionStore.Set(token, createdAt, dbUser.GetId()); err != nil {
		return "", err
	}

	return token, nil
}

func (jas *JwtAuthService) Logout(authorization, cookie string) error {
	unparsedToken := jas.getUnparsedToken(authorization, cookie)
	if unparsedToken == "" {
		return fmt.Errorf("Missing Token")
	}

	if err := jas.sessionStore.Delete(unparsedToken); err != nil {
		return fmt.Errorf("Invalid Token")
	}

	return nil
}
