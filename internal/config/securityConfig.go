package config

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
)

func CreateHeader() string {
	header := map[string]interface{}{
		"alg": "HS256",
		"typ": "JWT",
	}

	headerBytes, err := json.Marshal(header)
	if err != nil {
		fmt.Println("Erreur lors de l'encodage du header")
		return ""
	}

	return base64.RawURLEncoding.EncodeToString(headerBytes)
}

func EncodePayload(payload interface{}) string {
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Erreur lors de l'encodage du payload")
		return ""
	}
	return base64.RawURLEncoding.EncodeToString(payloadBytes)
}

func SignToken(header, payload, secretKey string) string {
	dataToSign := fmt.Sprintf("%s.%s", header, payload)
	signature := hmac.New(sha256.New, []byte(secretKey))
	signature.Write([]byte(dataToSign))
	return base64.RawURLEncoding.EncodeToString(signature.Sum(nil))
}

func CheckToken(header, payload, signature, secretKey string) bool {
	dataToVerify := fmt.Sprintf("%s.%s", header, payload)
	signatureBytes, err := base64.RawURLEncoding.DecodeString(signature)
	if err != nil {
		fmt.Println("Erreur lors de la décodage de la signature")
		return false
	}

	mac := hmac.New(sha256.New, []byte(secretKey))
	mac.Write([]byte(dataToVerify))
	expectedSignature := mac.Sum(nil)

	return hmac.Equal(expectedSignature, signatureBytes)
}

func BuildJWT(header, payload, secretKey string) string {
	tokenParts := []string{header, payload, SignToken(header, payload, secretKey)}
	return strings.Join(tokenParts, ".")
}

func SplitToken(token string) (header, payload, signature string) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return "", "", ""
	}
	return parts[0], parts[1], parts[2]
}
