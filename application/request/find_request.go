/*
* project mini-e-wallet
* created by oktoprima
* email : octoprima93@gmail.com
* github : https://github.com/oktopriima
* created at 00.57
**/

package request

type FindPagedRequest struct {
	Page int `form:"page"` // get value from query parameter
	Size int `form:"size"`
}

type FindByIDRequest struct {
	ID int `uri:"id"`
}

type FindBySlugRequest struct {
	Slug string `uri:"slug"`
}
