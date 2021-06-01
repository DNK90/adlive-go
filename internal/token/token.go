package token

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/dnk90/adlive/config"
	l "github.com/dnk90/adlive/helpers/log"
	"github.com/google/uuid"
	"time"
)

var ll = l.New()

type Token interface {
	CreateToken(id string) error
	ValidateToken() (string, bool)
	Expire() int64
	Secret() string
}

type JwtToken struct {
	Id           string
	MacAddress   string
	AccessToken  string
	AtExpireAt   int64
	RefreshToken string
	RtExpireAt   int64
	AtUUID       string
	RtUUID       string
}

func (t *JwtToken) Create() (err error) {
	t.AtExpireAt = t.AtExpire()
	t.RtExpireAt = t.RtExpire()
	t.AtUUID = uuid.NewString()
	t.AccessToken, err = t.createToken(t.AtExpireAt, t.AtUUID)
	if err != nil {
		return
	}
	t.RtUUID = uuid.NewString()
	t.RefreshToken, err = t.createToken(t.RtExpireAt, t.RtUUID)
	return
}

func (t JwtToken) createToken(expire int64, uuid string) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["id"] = t.Id
	atClaims["uuid"] = uuid
	atClaims["exp"] = expire
	atClaims["mac"] = t.MacAddress
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(t.Secret()))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (t JwtToken) AtExpire() int64 {
	return time.Now().Add(time.Minute * time.Duration(config.Cfg.JWT.Expire)).Unix()
}

func (t JwtToken) RtExpire() int64 {
	return time.Now().Add(time.Minute * time.Duration(config.Cfg.JWT.RefreshExpire)).Unix()
}

func (t JwtToken) Secret() string {
	return config.Cfg.JWT.Secret
}

func (t JwtToken) RefreshSecret() string {
	return config.Cfg.JWT.RefreshSecret
}

func (t JwtToken) validate(token, secret string) (string, bool) {
	parsedToken, claims, err := t.ParseToken(token, secret)
	if err != nil {
		ll.S.Errorw("error while parsing token", "err", err)
		return "", false
	}
	if !parsedToken.Valid {
		return "", false
	}
	id, ok := claims["id"].(string)
	if !ok {
		return "", false
	}
	return id, true
}

func (t JwtToken) ParseToken(token, secret string) (*jwt.Token, jwt.MapClaims, error) {
	claims := make(jwt.MapClaims)
	parsedToken, err := jwt.ParseWithClaims(token, claims, func(*jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, nil, err
	}
	return parsedToken, claims, nil
}

func (t JwtToken) ValidateToken() (string, bool) {
	return t.validate(t.AccessToken, t.Secret())
}

func (t *JwtToken) ValidateRefreshToken() bool {
	parsedToken, claims, err := t.ParseToken(t.RefreshToken, t.RefreshSecret())
	if err != nil {
		return false
	}
	if !parsedToken.Valid {
		return false
	}
	id, ok := claims["id"].(string)
	if !ok {
		return false
	}
	t.Id = id
	mac, ok := claims["mac"].(string)
	if ok {
		t.MacAddress = mac
	}
	return true
}
