package middleware

import (
	"errors"
	"fmt"
	"strings"

	"github.com/rochmanramadhani/go-lazisnu-api/internal/config"

	baseModel "github.com/rochmanramadhani/go-lazisnu-api/internal/model/abstraction"
	ctxval "github.com/rochmanramadhani/go-lazisnu-api/pkg/util/ctxval"
	res "github.com/rochmanramadhani/go-lazisnu-api/pkg/util/response"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func Authentication(next echo.HandlerFunc) echo.HandlerFunc {
	var (
		jwtKey = config.Config.JWT.Secret
	)

	return func(c echo.Context) error {
		authToken := c.Request().Header.Get("Authorization")
		if authToken == "" {
			return res.ErrorBuilder(res.Constant.Error.Unauthorized, errors.New("missing auth token")).Send(c)
		}

		splitToken := strings.Split(authToken, "Bearer ")
		if len(splitToken) < 2 {
			return res.ErrorBuilder(res.Constant.Error.Unauthorized, errors.New("invalid auth token")).Send(c)
		}

		token, err := jwt.Parse(splitToken[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method :%v", token.Header["alg"])
			}

			return []byte(jwtKey), nil
		})

		if !token.Valid || err != nil {
			return res.ErrorBuilder(res.Constant.Error.Unauthorized, err).Send(c)
		}

		var userID uint64
		destructID := token.Claims.(jwt.MapClaims)["user_id"]
		if destructID != nil {
			userID = uint64(destructID.(float64))
		}

		var roleID uint64
		destructRoleID := token.Claims.(jwt.MapClaims)["role_id"]
		if destructRoleID != nil {
			roleID = uint64(destructRoleID.(float64))
		}

		var companyID uint64
		destructCompanyID := token.Claims.(jwt.MapClaims)["company_id"]
		if destructCompanyID != nil {
			companyID = uint64(destructCompanyID.(float64))
		}

		var name string
		destructName := token.Claims.(jwt.MapClaims)["name"]
		if destructName != nil {
			name = destructName.(string)
		}

		authCtx := &baseModel.AuthContext{
			UserID:    userID,
			RoleID:    roleID,
			CompanyID: companyID,
			Name:      name,
		}
		ctx := ctxval.SetAuthValue(c.Request().Context(), authCtx)
		c.SetRequest(c.Request().WithContext(ctx))

		return next(c)
	}
}
