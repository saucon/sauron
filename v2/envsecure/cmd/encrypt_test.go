package cmd

import (
	"bytes"
	"github.com/saucon/sauron/v2/pkg/env"
	"github.com/saucon/sauron/v2/pkg/secure"
	"github.com/saucon/sauron/v2/sample"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
	"os"
	"testing"
)

func TestRootCmdWithEncrypt(t *testing.T) {
	// Create a buffer to capture the output
	output := &bytes.Buffer{}
	rootCmd.SetOut(output)
	rootCmd.SetErr(output)

	// Set up the arguments for the root command
	rootCmd.SetArgs([]string{
		"encrypt",
		"--file", "../../sample/env.sample.yml",
		"--keyfile", "../../sample/public_key.pem",
		"--algo", "rsa",
	})

	// Execute the root command
	err := rootCmd.Execute()

	// dump the output
	outputString := output.String()

	// Print the captured output for debugging
	t.Log(outputString)

	// Assert no errors occurred
	assert.NoError(t, err)

	// Assert the output contains the expected message
	assert.Contains(t, output.String(), "Config encrypted and saved successfully!")

	// try decrypt the config
	secureRsa := secure.NewSecureRSA()
	privSec, err := secure.ReadKeyFromFile("../../sample/private_key.pem")

	assert.NoError(t, err)

	var cfgSec sample.ConfigSample
	_, err = env.NewSecure(secureRsa, privSec, "secure.env.sample.yml", &cfgSec)

	assert.NoError(t, err)

	var cfg sample.ConfigSample
	data, err := os.ReadFile("../../sample/env.sample.yml")

	assert.NoError(t, err)

	err = yaml.Unmarshal(data, &cfg)
	assert.NoError(t, err)

	assert.Equal(t, cfgSec, cfg)
}
