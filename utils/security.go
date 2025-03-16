package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"log"
	"os"
)

// LoadPrivateKey carga la clave privada desde un archivo cuyo path está en .env.local
func LoadPrivateKey() (*rsa.PrivateKey, error) {
	privateKeyPath := os.Getenv("PRIVATE_KEY_PATH")
	if privateKeyPath == "" {
		return nil, fmt.Errorf("la variable de entorno PRIVATE_KEY_PATH no está definida")
	}

	keyData, err := os.ReadFile(privateKeyPath)
	if err != nil {
		return nil, fmt.Errorf("error al leer el archivo de clave privada: %v", err)
	}

	block, _ := pem.Decode(keyData)
	if block == nil {
		return nil, fmt.Errorf("error al decodificar la clave privada: bloque PEM inválido")
	}

	// Intentar parsear como PKCS#1
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err == nil {
		return privateKey, nil
	}

	// Si falla, intentar como PKCS#8
	privateKeyInterface, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("error al parsear la clave privada en PKCS#1 o PKCS#8: %v", err)
	}

	// Convertir a *rsa.PrivateKey
	privateKey, ok := privateKeyInterface.(*rsa.PrivateKey)
	if !ok {
		return nil, fmt.Errorf("la clave privada no es de tipo RSA")
	}

	return privateKey, nil
}

// LoadPublicKey carga la clave pública desde un archivo cuyo path está en .env.local
func LoadPublicKey() (*rsa.PublicKey, error) {
	publicKeyPath := os.Getenv("PUBLIC_KEY_PATH")
	log.Println("publicKeyPath: " + publicKeyPath)
	if publicKeyPath == "" {
		return nil, fmt.Errorf("error: PUBLIC_KEY_PATH no está definido en .env.local")
	}

	publicKeyBytes, err := os.ReadFile(publicKeyPath)
	if err != nil {
		return nil, fmt.Errorf("error al leer la clave pública: %v", err)
	}

	block, _ := pem.Decode(publicKeyBytes)
	if block == nil {
		return nil, fmt.Errorf("error al decodificar la clave pública")
	}

	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("error al parsear la clave pública: %v", err)
	}

	return publicKey.(*rsa.PublicKey), nil
}

func EncryptMessage(plainText string) (string, error) {
	publicKey, err := LoadPublicKey()
	if err != nil {
		return "", fmt.Errorf("error loading public key: %v", err)
	}

	encryptedBytes, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, publicKey, []byte(plainText), nil)
	if err != nil {
		return "", fmt.Errorf("error encrypting message: %v", err)
	}

	// Convertir a Base64 antes de enviarlo
	encoded := base64.StdEncoding.EncodeToString(encryptedBytes)
	return encoded, nil
}

// DecryptMessage desencripta un mensaje usando la clave privada
func DecryptMessage(encryptedBase64 string) (string, error) {
	privateKey, err := LoadPrivateKey()
	if err != nil {
		return "", fmt.Errorf("error loading private key: %v", err)
	}

	// Decodificar el mensaje de Base64 primero
	encryptedData, err := base64.StdEncoding.DecodeString(encryptedBase64)
	if err != nil {
		return "", fmt.Errorf("error decoding base64 data: %v", err)
	}

	// Desencriptar el mensaje con OAEP
	decryptedBytes, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, privateKey, encryptedData, nil)
	if err != nil {
		return "", fmt.Errorf("error decrypting data: %v", err)
	}

	return string(decryptedBytes), nil
}
