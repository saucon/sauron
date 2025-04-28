package cmd

type Secure interface {
	Encrypt(plainText string, secretKey string) (string, error)
	Decrypt(cipherText string, secretKey string) (string, error)
}
