package jwt

import (
	"XinAPI/pkg/path"
	"io/ioutil"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTStandard struct {
	Id        string
	Name      string
	ValidTime time.Duration
	Issuer    string
	Subject   string
}

// get time for Local
func timeIn(name string) time.Time {
	loc, err := time.LoadLocation(name)
	if err != nil {
		panic(err)
	}
	return time.Now().In(loc)
}

func (Stand *JWTStandard) Create() (string, error) {

	timestamp := timeIn("Asia/Taipei")

	claims := jwt.StandardClaims{
		Audience:  Stand.Name,                            // 授予者名稱
		Id:        Stand.Id,                              // 授予者ID
		IssuedAt:  timestamp.Unix(),                      // 簽發時間
		NotBefore: timestamp.Unix(),                      // 生效時間
		ExpiresAt: timestamp.Add(Stand.ValidTime).Unix(), // 失效時間
		Issuer:    Stand.Issuer,                          // 發簽者
		Subject:   Stand.Subject,                         // 簽章用途
	}

	bytes, _ := ioutil.ReadFile(path.BasePath("/storage/key/graphql/api/oauth-private.key"))
	rsaPrivate, _ := jwt.ParseRSAPrivateKeyFromPEM(bytes)

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	token, err := tokenClaims.SignedString(rsaPrivate)

	return token, err

}

func Validate() (*jwt.StandardClaims, error) {
	token := "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJoZWxsbyB3b3JsZCIsImV4cCI6MTU4MjQ0NTI4MywianRpIjoiMSIsImlhdCI6MTU4MjQ0MTY4MywiaXNzIjoiTGlvbiBGMkUiLCJuYmYiOjE1ODI0NDE2ODMsInN1YiI6ImxvZ2luIn0.ImCW6wEQ82rv1oStilBMrxvuYK7P5x9FtvlgKDfQq56RbzxzHJZear4W3snyRV_pfASiH9WYS8ffYoVTbjPb6JmTf9mlIrHAkMCw7HU8NGsK8EBPf5-THbGkOr2JYOu-qfLJ-JLjnTgWv64dyrVs8RM8DqxSLCehuj7QAFBHKTJrmJ9NXyG3Uu0rmtBij0-c6GPibA1nyFniSnshelnK_locvGNy1-ngS3hKma4fhtaXaF4nMgrheMppfw9BbVvvoyBZY5Sjd4V4nLrfl_kJqxKD7EJBhV_dyn4mZTMgE2CXVFfAhTHrvu4-YOAOEdktshXrMVgHY7450JWFW1cGWA"

	bytes, _ := ioutil.ReadFile(path.BasePath("/storage/key/graphql/api/oauth-public.key"))
	rsaPublic, err := jwt.ParseRSAPublicKeyFromPEM(bytes)

	jwtToken, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return rsaPublic, nil
	})

	if err == nil && jwtToken != nil {
		if claim, ok := jwtToken.Claims.(*jwt.StandardClaims); ok && jwtToken.Valid {
			return claim, nil
		}
	}
	return nil, err

}

// func Revoke() {

// }
