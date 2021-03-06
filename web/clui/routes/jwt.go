/*
Copyright <holder> All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package routes

import (
	"crypto/rsa"
	"fmt"
	"time"

	"math/rand"

	"github.com/IBM/cloudland/web/clui/model"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

var (
	_publicKey  *rsa.PublicKey
	_privateKey *rsa.PrivateKey
)

type HypercubeClaims struct {
	jwt.StandardClaims
	UID  int64      `json:"uid,omitempty"`
	OID  int64      `json:"oid,omitempty"`
	Role model.Role `json:"r,omitempty"`
}

func NewClaims(u, o string, uid, oid int64, role model.Role) (claims jwt.Claims) {
	now := time.Now()
	issuedAt := now.Unix()
	return &HypercubeClaims{
		StandardClaims: jwt.StandardClaims{
			Audience:  u,
			ExpiresAt: now.Add(time.Hour * 2).Unix(),
			Id:        claimsID(now),
			IssuedAt:  issuedAt,
			Issuer:    "Cloudland",
			NotBefore: issuedAt,
			Subject:   o,
		},
		UID:  uid,
		OID:  oid,
		Role: role,
	}
}

func claimsID(now time.Time) string {
	return fmt.Sprintf("%d", now.UnixNano()+rand.Int63())
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func publicKey() *rsa.PublicKey {
	if _publicKey == nil {
		key := viper.GetString("key.public")
		if key == "" {
			panic("No public key provided")
		}
		var err error
		_publicKey, err = jwt.ParseRSAPublicKeyFromPEM([]byte(key))
		if err != nil {
			panic(err)
		}
	}
	return _publicKey
}

func privateKey() *rsa.PrivateKey {
	if _privateKey == nil {
		key := viper.GetString("key.private")
		if key == "" {
			panic("No private key provided")
		}
		var err error
		_privateKey, err = jwt.ParseRSAPrivateKeyFromPEM([]byte(key))
		if err != nil {
			panic(err)
		}

	}
	return _privateKey
}

func NewToken(u, o string, uid, oid int64, role model.Role) (signed string, err error) {
	claims := NewClaims(u, o, uid, oid, role)
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	signed, err = token.SignedString(privateKey())
	return
}
