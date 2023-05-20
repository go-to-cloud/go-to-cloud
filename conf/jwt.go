package conf

import "sync"

type JWT struct {
	Security string
	Realm    string
	IdKey    string
}

var jwt *JWT

var onceJwt sync.Once

// GetJwtKey 获取JWT私钥
func GetJwtKey() *JWT {
	onceJwt.Do(func() {
		if jwt == nil {
			j := getConf().Jwt
			jwt = &JWT{
				Security: j.Security,
				Realm:    j.Realm,
			}
		}
	})

	return jwt
}
