/*
* project mini-e-wallet
* created by oktoprima
* email : octoprima93@gmail.com
* github : https://github.com/oktopriima
* created at 02.56
**/

package authentication

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/oktopriima/mini-e-wallet/application/request"
	"github.com/oktopriima/mini-e-wallet/domain/middleware"
)

func (a *authController) LoginController(request request.LoginRequest) (interface{}, error) {
	user, err := a.userRepo.FindOne(map[string]interface{}{
		"email": request.Email,
	})
	if err != nil {
		return nil, fmt.Errorf("email %s %v", request.Email, err)
	}

	if !a.comparePassword(user.Password, request.Password) {
		return nil, fmt.Errorf("password doesn't match")
	}

	// update last login
	tx := a.db.Begin()
	defer tx.Rollback()
	user.LastLogin = time.Now()
	if err := a.userRepo.Update(user, tx); err != nil {
		return nil, err
	}
	tx.Commit()

	param := new(middleware.TokenStructure)
	param.Email = user.Email
	param.UserID = int64(user.ID)

	customAuth := middleware.NewCustomAuth([]byte(a.conf.GetString("app.signature")))
	resp, err := customAuth.GenerateToken(*param)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *authController) comparePassword(hashed string, plain string) bool {
	byteHash := []byte(hashed)
	bytePlain := []byte(plain)
	err := bcrypt.CompareHashAndPassword(byteHash, bytePlain)
	if err != nil {
		return false
	}
	return true
}
