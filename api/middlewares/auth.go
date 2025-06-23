package middlewares

import (
	"context"
	"errors"
	"os"
	"strings"
	"todo/apperrors"
	"todo/common"

	"github.com/labstack/echo/v4"
	"google.golang.org/api/idtoken"
)

var (
	googleClientID = os.Getenv("ClientId")
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		// headerからAuthorizationフィールドを抜き出す
		authorization := ctx.Request().Header.Get("Authorization")

		// Authorization フィールドが"Bearer [ID トークン]"の形になっているか検証
		authHeaders := strings.Split(authorization, " ")
		if len(authHeaders) != 2 {
			err := apperrors.RequiredAuthorizationHeader.Wrap(errors.New("invalid req header"), "invalid header")
			apperrors.ErrorHandler(ctx, err)
			return err
		}

		bearer, idToken := authHeaders[0], authHeaders[1]
		if bearer != "Bearer" || idToken == "" {
			err := apperrors.RequiredAuthorizationHeader.Wrap(errors.New("invalid req header"), "invalid header")
			apperrors.ErrorHandler(ctx, err)
			return err
		}

		tokenValidator, err := idtoken.NewValidator(context.Background())
		if err != nil {
			err := apperrors.CannotMakeValidator.Wrap(err, "internal auth err")
			apperrors.ErrorHandler(ctx, err)
			return err
		}

		payload, err := tokenValidator.Validate(context.Background(), idToken, googleClientID)
		if err != nil {
			err = apperrors.CannotMakeValidator.Wrap(err, "invalid id token")
			apperrors.ErrorHandler(ctx, err)
			return err
		}

		// name フィールドをpayloadから抜き出す
		name, ok := payload.Claims["name"]
		if !ok {
			err = apperrors.Unauthorizated.Wrap(err, "invalid id token")
			apperrors.ErrorHandler(ctx, err)
			return err
		}

		// contextにユーザー名をセット
		common.SetUserName(ctx.Request().Context(), name.(string))

		// Handlerへ
		err = next(ctx)

		return err
	}
}