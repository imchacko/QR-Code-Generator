# QR Code Generator API

## Overview
This API allows users to generate a QR code for a UPI ID and retrieve the QR code image URL. It is built using **AWS Lambda (Golang)**, **API Gateway**, and **S3** for storing QR codes.

## Base URL
```
https://f9ys51ix8i.execute-api.eu-north-1.amazonaws.com/dev
```

## Endpoint
### Generate QR Code
```
GET /QRCodeGenerator
```

### Request Parameters
| Parameter | Description |
|-----------|-------------|
| `upi_id`  | The UPI ID for which the QR code is to be generated (required) |

### Example Request (Postman)
**Method:** `GET`
```
https://f9ys51ix8i.execute-api.eu-north-1.amazonaws.com/dev/QRCodeGenerator?upi_id=rahul.chacko@okhdfc
```

## Response
A successful request returns a JSON response containing the download URL of the generated QR code.

### Example Success Response
```json
{
  "download_url": "https://f9ys51ix8i.execute-api.eu-north-1.amazonaws.com/dev/qrcodes/1700000000.png"
}
```

## Error Handling
If an error occurs, the API returns an appropriate HTTP status code and an error message.

### Example Error Response
```json
{
  "error": "upi_id is required"
}
```

## How to Use
1. Clone the repository:
   ```bash
   git clone https://github.com/imchacko/QR-Code-Generator.git
   ```
2. Navigate to the project directory:
   ```bash
   cd QR-Code-Generator
   ```
3. Run the application locally (if applicable):
   ```bash
   go run main.go
   ```

## Deploying the Lambda Function
### 1. Build the Go Binary
AWS Lambda expects a compiled binary instead of a raw Go file. Run the following command inside your project directory:

```bash
GOOS=linux GOARCH=amd64 go build -o bootstrap main.go
```

### 2. Zip the binary file
```bash
zip deployment.zip bootstrap
```

### 3. Create an AWS Lambda Function
Create a Lambda function using the AWS CLI:

```bash
aws lambda create-function \
  --function-name QRCodeGenerator \
  --runtime provided.al2 \
  --handler bootstrap \
  --memory-size 128 \
  --timeout 30 \
  --role arn:aws:iam::114826118786:role/LambdaExecutionRole \
  --zip-file fileb://deployment.zip
```


### 4. Deploy Updates
If you update your function, recompile and upload:

```bash
GOOS=linux GOARCH=amd64 go build -o bootstrap main.go
zip deployment.zip bootstrap

aws lambda update-function-code \
  --function-name QRCodeGenerator \
  --zip-file fileb://deployment.zip
```

## Setting Up API Gateway
To expose your Lambda function via an HTTP endpoint:
1. Go to the **AWS API Gateway Console**.
2. Create a **REST API**.
3. Create a **GET Method** and link it to your `QRCodeGenerator` Lambda function.
4. Deploy the API and get the **Invoke URL**.

## Executing the API
Test the API using Postman or `curl`:
```bash
curl -X GET "https://f9ys51ix8i.execute-api.eu-north-1.amazonaws.com/dev/QRCodeGenerator?upi_id=rahul.chacko@okhdfc"
```

## AWS Credentials Verification
To verify your AWS credentials, run:
```bash
aws sts get-caller-identity
```
Example output:
```json
{
    "UserId": "AIDARVPBLQ2BD7MGN4LFH",
    "Account": "114826118786",
    "Arn": "arn:aws:iam::114826118786:user/rahul_chacko"
}
```

