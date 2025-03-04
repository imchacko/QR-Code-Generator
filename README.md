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
  "download_url": "https://your-s3-bucket.s3.eu-north-1.amazonaws.com/qrcodes/1700000000.png"
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


