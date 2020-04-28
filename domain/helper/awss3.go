/*
 * Name : Okto Prima Jaya
 * Github : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 18/11/2019, 20:52
 * Copyright (c) 2019
 */

package helper

import (
	"bytes"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"

	"github.com/oktopriima/mini-e-wallet/domain/config"
)

func UploadToS3(dest string, filename string) error {
	conf := config.NewConfig()

	awsconf := &aws.Config{
		Region: aws.String(conf.GetString("aws.region")),
		Credentials: credentials.NewStaticCredentials(
			conf.GetString("aws.s3.access_key"),
			conf.GetString("aws.s3.secret_key"),
			"",
		),
	}

	s, err := session.NewSession(awsconf)
	if err != nil {
		return err
	}

	if err = uploadFile(
		s,
		dest,
		conf.GetString("media.storage")+filename,
		conf.GetString("aws.s3.bucket"),
	); err != nil {
		return err
	}

	return nil
}

func uploadFile(s *session.Session, dest, filename, bucket string) error {
	var size int64
	tmpfile, err := os.Open(dest)
	if err != nil {
		return err
	}

	defer tmpfile.Close()

	fileinfo, _ := tmpfile.Stat()
	size = fileinfo.Size()
	buffer := make([]byte, size)
	if _, err = tmpfile.Read(buffer); err != nil {
		return err
	}

	_, err = s3.New(s).PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(bucket),
		Key:                  aws.String(filename),
		ACL:                  aws.String("private"),
		Body:                 bytes.NewReader(buffer),
		ContentLength:        aws.Int64(size),
		ContentType:          aws.String(http.DetectContentType(buffer)),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
	})

	return err
}
