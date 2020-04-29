/*
* project mini-e-wallet
* created by oktoprima
* email : octoprima93@gmail.com
* github : https://github.com/oktopriima
* created at 14.16
**/

package request

type TransferRequest struct {
	UserDestination   int     `json:"user_destination"`
	Amount            float64 `json:"amount"`
	IP                string  `json:"-"`
	UserAgent         string  `json:"-"`
	Location          string  `json:"-"`
	AuthenticatedUser int     `json:"-"`
}
