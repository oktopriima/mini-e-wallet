/*
* project mini-e-wallet
* created by oktoprima
* email : octoprima93@gmail.com
* github : https://github.com/oktopriima
* created at 02.55
**/

package request

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
