package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/skip2/go-qrcode"
)

var (
	bucketName = os.Getenv("lambdas3images")
	region     = os.Getenv("eu-north-1")
	s3Client   *s3.S3
)

// Generate QR Code for UPI ID
func generateQRCode(upiID string) ([]byte, error) {
	upiURL := fmt.Sprintf("upi://pay?pa=%s", upiID)
	qrCode, err := qrcode.Encode(upiURL, qrcode.Medium, 256)
	if err != nil {
		return nil, err
	}
	return qrCode, nil
}

// Upload QR Code to S3
func uploadToS3(fileName string, qrCode []byte) (string, error) {
	_, err := s3Client.PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(bucketName),
		Key:         aws.String(fileName),
		Body:        bytes.NewReader(qrCode),
		ContentType: aws.String("image/png"),
		ACL:         aws.String("public-read"),
	})

	if err != nil {
		return "", err
	}

	fileURL := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", bucketName, region, fileName)
	return fileURL, nil
}

// Lambda Handler
func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	upiID := request.QueryStringParameters["upi_id"]
	if upiID == "" {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       `{"error": "upi_id is required"}`,
		}, nil
	}

	qrCode, err := generateQRCode(upiID)
	if err != nil {
		log.Println("QR Code Generation Error:", err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       `{"error": "Failed to generate QR code"}`,
		}, nil
	}

	fileName := fmt.Sprintf("qrcodes/%d.png", time.Now().Unix())
	fileURL, err := uploadToS3(fileName, qrCode)
	if err != nil {
		log.Println("S3 Upload Error:", err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       `{"error": "Failed to upload QR code"}`,
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       fmt.Sprintf(`{"download_url": "%s"}`, fileURL),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil
}

func main() {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(region),
	}))
	s3Client = s3.New(sess)

	lambda.Start(handler)
}
