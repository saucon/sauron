package secure

type SecureImplAEAD struct{}

func NewSecureAEAD() SecureImplAEAD {
	return SecureImplAEAD{}
}

func (SecureImplAEAD) Encrypt(plainText string, secretKey string) (string, error) {
	chipperText, err := encryptAEAD(plainText, secretKey)
	if err != nil {
		return "", err
	}
	return chipperText, nil
}

func (SecureImplAEAD) Decrypt(cipherText string, secretKey string) (string, error) {
	plainText, err := decryptAEAD(cipherText, secretKey)
	if err != nil {
		return "", err
	}
	return plainText, nil
}
