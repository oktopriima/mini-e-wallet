/*
 * Name : Okto Prima Jaya
 * Github : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 18/11/2019, 20:52
 * Copyright (c) 2019
 */

package helper

import (
	"encoding/json"
	"errors"
	"net/http"
)

type FacebookResponse struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Birthday string `json:"birthday"`
	Email    string `json:"email"`
	Gender   string `json:"gender"`
	Avatar   string `json:"avatar"`
}

func GetFacebookData(token string) (*FacebookResponse, error) {
	var uri string
	uri = "https://graph.facebook.com/v3.1/me?fields=id,name,about,birthday,gender,email&access_token=" + token

	resp, err := http.Get(uri)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New("failed get facebook data")
	}

	byteResp, err := ResponseToByte(resp)
	if err != nil {
		return nil, err
	}

	data := new(FacebookResponse)

	if err = json.Unmarshal(byteResp, &data); err != nil {
		return nil, err
	}

	/** create avatar url **/
	data.Avatar = "https://graph.facebook.com/v3.1/" + data.ID + "/picture?type=large"
	return data, nil
}
