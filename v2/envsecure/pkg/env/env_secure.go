package env

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
	"os"
	"strconv"
	"strings"
)

type Secure interface {
	Encrypt(plainText string, secretKey string) (string, error)
	Decrypt(cipherText string, secretKey string) (string, error)
}

func EncryptEnv(secure Secure, filepath string, secretKey string) error {
	var configMap map[string]any

	fp := strings.Split(filepath, "/")
	filename := fp[len(fp)-1]

	// 1. Read YAML file
	data, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}

	// 2. Unmarshal into a map
	err = yaml.Unmarshal(data, &configMap)
	if err != nil {
		return err
	}
	//printMap(configMap, "")

	encryptInterface(secure, configMap, secretKey)

	// Create the output file
	file, err := os.Create("secure." + filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return err
	}
	defer file.Close()

	// Create a YAML encoder
	encoder := yaml.NewEncoder(file)
	encoder.SetIndent(2)

	// Encode the modified node
	if err = encoder.Encode(configMap); err != nil {
		fmt.Println("Error encoding YAML:", err)
		return err
	}

	return nil
}

func decryptEnv(secure Secure, decryptedBuf *bytes.Buffer, filepath string, secretKey string) error {
	var configMap map[string]any

	data, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(data, &configMap)
	if err != nil {
		return err
	}

	decryptInterface(secure, configMap, secretKey)

	// Create a YAML encoder
	encoder := yaml.NewEncoder(decryptedBuf)
	encoder.SetIndent(2)

	// Encode the modified node
	if err = encoder.Encode(configMap); err != nil {
		fmt.Println("Error encoding YAML:", err)
		return err
	}

	return nil
}

func NewSecure(secure Secure, secretKey string, filepath string, config ...any) (ev EnvConfig, err error) {
	var vi *viper.Viper
	for _, c := range config {
		vi, err = loadConfigViperSecure(secure, secretKey, filepath, &c)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"error_cause":   PrintErrorStack(err),
				"error_message": err.Error(),
			}).Fatal("config/env: load config")
		}
	}

	ev.ViperInstance = vi
	ev.Config = &config
	return ev, err
}

func loadConfigViperSecure(secure Secure, secretKey string, filepath string, config any) (vi *viper.Viper, err error) {
	var decryptedBuf bytes.Buffer

	err = decryptEnv(secure, &decryptedBuf, filepath, secretKey)
	if err != nil {
		return nil, err
	}

	// get yaml
	vi = viper.New()
	vi.SetConfigType("yaml")
	if err := vi.ReadConfig(bytes.NewBuffer(decryptedBuf.Bytes())); err != nil {
		_ = fmt.Errorf("error reading config file, %v", err)
	}

	keysConfig := getAllKeys(vi.AllSettings())
	if errr := vi.Unmarshal(&config); errr != nil {
		errr = fmt.Errorf("error Unmarshal config file, %v", errr)
		logrus.WithFields(logrus.Fields{
			"error_cause":   PrintErrorStack(errr),
			"error_message": errr.Error(),
		}).Debug("config/env: load config")
	}

	vi.AllKeys()

	// set env
	envPrefix := vi.GetString("envLib.app.envPrefix")
	vi.SetEnvPrefix(envPrefix)
	vi.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	for _, key := range keysConfig {
		envKey := envPrefix + "_" + strings.ToUpper(strings.ReplaceAll(key, ".", "_"))
		err := os.Setenv(envKey, vi.GetString(key))
		if err != nil {
			return nil, err
		}
		vi.MustBindEnv(key, envKey)
	}
	vi.AutomaticEnv()

	return vi, err
}

func getAllKeys(settings map[string]any) []string {
	allKeys := []string{}

	for keyParent, v := range settings {
		switch v := v.(type) {
		case map[string]any:
			s := getAllKeys(v)
			for _, keyValue := range s {
				allKeys = append(allKeys, keyParent+"."+keyValue)
			}
		case []any:
			s := getAllKeysArr(v)
			for _, keyValue := range s {
				allKeys = append(allKeys, keyParent+"."+keyValue)
			}
		case string:
			allKeys = append(allKeys, keyParent)
		}
	}
	return allKeys
}

func getAllKeysArr(arr []any) []string {
	allKeys := []string{}

	for keyParent, v := range arr {
		switch v := v.(type) {
		case map[string]any:
			s := getAllKeys(v)
			for _, key := range s {
				allKeys = append(allKeys, strconv.Itoa(keyParent)+"."+key)
			}
		case string:
			allKeys = append(allKeys, strconv.Itoa(keyParent))
		}
	}
	return allKeys
}
