/*
* project mini-e-wallet
* created by oktoprima
* email : octoprima93@gmail.com
* github : https://github.com/oktopriima
* created at 02.21
**/

package request

type TopUpRequest struct {
	Amount            float64 `json:"amount"`
	IP                string  `json:"ip"`
	UserAgent         string  `json:"user_agent"`
	Location          string  `json:"location"`
	AuthenticatedUser int     `json:"-"`
}
