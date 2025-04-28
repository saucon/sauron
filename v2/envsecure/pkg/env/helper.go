package env

import (
	"fmt"
)

// Recursive function to print key-value
func printMap(m map[string]interface{}, prefix string) {
	for k, v := range m {
		switch child := v.(type) {
		case map[string]interface{}:
			// If nested map, go deeper
			printMap(child, prefix+k+".")
		default:
			// Otherwise, print key and value
			fmt.Printf("%s%s: %v\n", prefix, k, v)
		}
	}
}

// encryptInterface encrypts values inside maps or slices recursively
func encryptInterface(secure Secure, i interface{}, secretKey string) {
	switch v := i.(type) {
	case map[string]interface{}:
		for k, val := range v {
			switch child := val.(type) {
			case map[string]interface{}, []interface{}:
				// Recurse deeper
				encryptInterface(secure, child, secretKey)
			case string:
				enc, err := secure.Encrypt(child, secretKey)
				if err != nil {
					panic(fmt.Errorf("failed to encrypt key %s: %w", k, err))
				}
				v[k] = enc
			case nil: // do nothing for nil
			default:
				s := fmt.Sprintf("%v", child)
				enc, err := secure.Encrypt(s, secretKey)
				if err != nil {
					panic(fmt.Errorf("failed to encrypt key %s: %w", k, err))
				}
				v[k] = enc
			}
		}
	case []interface{}:
		for idx, item := range v {
			switch child := item.(type) {
			case map[string]interface{}, []interface{}:
				encryptInterface(secure, child, secretKey)
			case string:
				enc, err := secure.Encrypt(child, secretKey)
				if err != nil {
					panic(fmt.Errorf("failed to encrypt index %d: %w", idx, err))
				}
				v[idx] = enc
			case nil: // do nothing for nil
			default:
				s := fmt.Sprintf("%v", child)
				enc, err := secure.Encrypt(s, secretKey)
				if err != nil {
					panic(fmt.Errorf("failed to encrypt index %d: %w", idx, err))
				}
				v[idx] = enc
			}
		}
	default:
		// do nothing for other types
	}
}

// decryptInterface encrypts values inside maps or slices recursively
func decryptInterface(secure Secure, i interface{}, secretKey string) {
	switch v := i.(type) {
	case map[string]interface{}:
		for k, val := range v {
			switch child := val.(type) {
			case map[string]interface{}, []interface{}:
				// Recurse deeper
				decryptInterface(secure, child, secretKey)
			case string:
				enc, err := secure.Decrypt(child, secretKey)
				if err != nil {
					panic(fmt.Errorf("failed to Decrypt key %s: %w", k, err))
				}
				v[k] = enc
			case nil: // do nothing for nil
			default:
				s := fmt.Sprintf("%v", child)
				enc, err := secure.Decrypt(s, secretKey)
				if err != nil {
					panic(fmt.Errorf("failed to Decrypt key %s: %w", k, err))
				}
				v[k] = enc
			}
		}
	case []interface{}:
		for idx, item := range v {
			switch child := item.(type) {
			case map[string]interface{}, []interface{}:
				decryptInterface(secure, child, secretKey)
			case string:
				enc, err := secure.Decrypt(child, secretKey)
				if err != nil {
					panic(fmt.Errorf("failed to Decrypt index %d: %w", idx, err))
				}
				v[idx] = enc
			case nil: // do nothing for nil
			default:
				s := fmt.Sprintf("%v", child)
				enc, err := secure.Decrypt(s, secretKey)
				if err != nil {
					panic(fmt.Errorf("failed to Decrypt index %d: %w", idx, err))
				}
				v[idx] = enc
			}
		}
	default:
		// do nothing for other types
	}
}
