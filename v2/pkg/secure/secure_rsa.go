package secure

type SecureImplRSA struct{}

func NewSecureRSA() SecureImplRSA {
	return SecureImplRSA{}
}

func (SecureImplRSA) Encrypt(plainText string, secretKey string) (string, error) {
	chipperText, err := encryptRSA(plainText, secretKey)
	if err != nil {
		return "", err
	}
	return chipperText, nil
}

func (SecureImplRSA) Decrypt(cipherText string, secretKey string) (string, error) {
	plainText, err := decryptRSA(cipherText, secretKey)
	if err != nil {
		return "", err
	}
	return plainText, nil
}
