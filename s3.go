package main

import (
	"github.com/minio/minio-go"
	log "github.com/sirupsen/logrus"
	"io"
	"net/url"
	"time"
)

type S3Connection struct {
	con    *minio.Client
	bucket string
}

func (c *S3Connection) Init(host, access, secret, bucket string) {
	s3, err := minio.NewV2(host, access, secret, true)
	if err != nil {
		log.Fatal(err)
	}
	c.con = s3
	c.bucket = bucket
}

func (c S3Connection) PresignPlug(plug Plug) *url.URL {
	presignedURL, err := c.con.PresignedGetObject(c.bucket, plug.S3ID, time.Duration(60)*time.Second, make(url.Values))
	if err != nil {
		log.Fatal(err)
	}

	return presignedURL
}

func (c S3Connection) AddFile(plug Plug, data io.Reader, mime string) {
	opts := new(minio.PutObjectOptions)
	opts.ContentType = mime
	_, err := c.con.PutObject(c.bucket, plug.S3ID, data, -1, *opts)
	if err != nil {
		log.Error(err)
	}
}

func (c S3Connection) DelFile(plug Plug) {
	err := c.con.RemoveObject(c.bucket, plug.S3ID)
	if err != nil {
		log.Error(err)
	}
}
