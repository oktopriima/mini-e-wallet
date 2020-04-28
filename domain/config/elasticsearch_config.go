/*
 * Name : Okto Prima Jaya
 * Github : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 18/11/2019, 20:50
 * Copyright (c) 2019
 */

package config

import (
	"context"
	"time"

	"github.com/olivere/elastic"
)

func NewElasticsearchConfig(cfg Config) (*elastic.Client, error) {
	var uri, ip, port string
	ctx := context.Background()

	ip = cfg.GetString("elasticsearch.address")
	port = cfg.GetString("elasticsearch.port")

	uri = "http://" + ip + ":" + port

	client, err := elastic.NewClient(
		elastic.SetURL(uri),
		elastic.SetSniff(false),
		elastic.SetHealthcheckInterval(10*time.Second),
	)
	if err != nil {
		return nil, err
	}

	if _, _, err := client.Ping(uri).Do(ctx); err != nil {
		return nil, err
	}

	return client, nil
}
