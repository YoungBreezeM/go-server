package services

import (
	"errors"
	"server/config"
	"server/constant"
	"server/db"
	"server/log"
	"server/models"
	"server/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func Subscribe(openId string) {
	//find user in db
	user := db.User{OpenId: openId}
	b := user.UserIsExistByOpenId()
	//add new user
	if !b {
		user.Integral = config.INIT_INTEGRAL
		user.IsVip = false
		user.Status = constant.SUBSCRIBE
		if err := user.AddUser(); err != nil {
			log.Log.Errorln(err)
		}
	} else {
		user.Status = constant.SUBSCRIBE
		if err := user.UpdateUserStatus(); err != nil {
			log.Log.Errorln(err)
		}

	}
}

func Unsubscribe(openId string) {
	user := db.User{OpenId: openId}
	b := user.UserIsExistByOpenId()

	//update user status
	if b {
		user.Status = constant.UNSUBSCRIBE
		if err := user.UpdateUserStatus(); err != nil {
			log.Log.Errorln(err)
		}
	}
}

func VerifyToken(token string) bool {
	tokenClaims, err := utils.ParseToken(token)
	//
	if err != nil {
		log.Log.Errorln(err)
		return false
	}
	//
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*models.JWTClaims); ok && tokenClaims.Valid {
			user := db.User{
				OpenId: claims.OpenId,
			}
			return user.UserIsExistByOpenId()

		}

	}

	return false
}

func VerifyKey(c *gin.Context, token string) bool {
	key := token[:32]
	openId := token[32:]
	val, err := db.Redis.Get(c, openId).Result()
	if err != nil {
		if err == redis.Nil {
			return false
		}
		return false
	}
	return key == val
}
func GetUserInfoByToken(token string) (*db.User, error) {
	tokenClaims, err := utils.ParseToken(token)
	//
	if err != nil {
		return nil, err
	}
	//
	if claims, ok := tokenClaims.Claims.(*models.JWTClaims); ok && tokenClaims.Valid {
		user := db.User{
			OpenId: claims.OpenId,
		}
		if err := user.FindUserByOpenId(); err != nil {
			return nil, err
		}
		return &user, nil

	}

	return nil, errors.New("token parassing is error")
}
