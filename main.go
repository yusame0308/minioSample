package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"log"
	"os"
)

func main() {
	// sessionの作成
	ak := "minio"
	sk := "miniopass"
	cfg := aws.Config{
		Credentials:      credentials.NewStaticCredentials(ak, sk, ""),
		Region:           aws.String("ap-northeast-1"),
		Endpoint:         aws.String("http://localhost:9000"),
		S3ForcePathStyle: aws.Bool(true),
	}

	sess := session.Must(session.NewSession(&cfg))
	svc := s3.New(sess, &cfg)

	bucketName := "test"
	objectKey := "test.txt"

	obj, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	})
	if err != nil {
		log.Fatal(err)
	}

	// 最初の10byteだけ読み込んで表示
	rc := obj.Body
	defer rc.Close()
	buf := make([]byte, 15)
	_, err = rc.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s", buf)

	targetFilePath := "badminton.png"
	file, err := os.Open(targetFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	bName := "aaa"
	oKey := "badminton"

	// Uploaderを作成し、ローカルファイルをアップロード
	uploader := s3manager.NewUploader(sess)
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bName),
		Key:    aws.String(oKey),
		Body:   file,
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("done")
}
