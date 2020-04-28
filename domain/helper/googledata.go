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

type GoogleResponse struct {
	Kind   string `json:"kind"`
	Etag   string `json:"etag"`
	Gender string `json:"gender"`
	Emails []struct {
		Value string `json:"value"`
		Type  string `json:"type"`
	} `json:"emails"`
	ObjectType  string `json:"objectType"`
	ID          string `json:"id"`
	DisplayName string `json:"displayName"`
	Name        struct {
		FamilyName string `json:"familyName"`
		GivenName  string `json:"givenName"`
	} `json:"name"`
	PlusURL string `json:"url"`
	Image   struct {
		URL       string `json:"url"`
		IsDefault bool   `json:"isDefault"`
	} `json:"image"`
	IsPlusUser     bool   `json:"isPlusUser"`
	Language       string `json:"language"`
	CircledByCount int    `json:"circledByCount"`
	Verified       bool   `json:"verified"`
	Cover          struct {
		Layout     string `json:"layout"`
		CoverPhote struct {
			URL    string `json:"url"`
			Height int    `json:"height"`
			Width  int    `json:"width"`
		} `json:"coverPhoto"`
		CoverInfo struct {
			TopImageOffset  int `json:"topImageOffset"`
			LeftImageOffset int `json:"leftImageOffset"`
		} `json:"coverInfo"`
	} `json:"cover"`
	Domain string `json:"domain"`
}

func GetGoogleData(token string) (*GoogleResponse, error) {
	var uri string
	uri = "https://www.googleapis.com/plus/v1/people/me?access_token=" + token

	/** make request to google */
	resp, err := http.Get(uri)
	if err != nil {
		return nil, err
	}

	/** check status code */
	if resp.StatusCode != 200 {
		panic(resp.Status)
		return nil, errors.New("failed get google data")
	}

	/** convert response to byte */
	bodyByte, err := ResponseToByte(resp)
	if err != nil {
		return nil, err
	}

	data := new(GoogleResponse)
	if err = json.Unmarshal(bodyByte, &data); err != nil {
		return nil, err
	}

	return data, nil
}
