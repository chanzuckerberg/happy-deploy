package workspace_repo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/jeremywohl/flatten"
	"github.com/pkg/errors"
)

const tfrcFileName = ".terraform.d/credentials.tfrc.json"
const terraformHostName = "si.prod.tfe.czi.technology"

func GetTfeToken() (string, error) {
	token, ok := os.LookupEnv("TFE_TOKEN")
	if !ok {
		token, err := readTerraformTokenFile()
		if err == nil {
			return token, nil
		}

		composeArgs := []string{"terraform", "login", terraformHostName}

		docker, err := exec.LookPath("terraform")
		if err != nil {
			return "", errors.New("Please set env var TFE_TOKEN")
		}

		cmd := &exec.Cmd{
			Path:   docker,
			Args:   composeArgs,
			Stdout: os.Stdout,
			Stderr: os.Stderr,
			Stdin:  os.Stdin,
		}
		err = cmd.Run()
		if err != nil {
			return "", errors.New("Please set env var TFE_TOKEN")
		}
		token, err = readTerraformTokenFile()
		if err != nil {
			return "", errors.New("Please set env var TFE_TOKEN")
		}
		return token, nil

	}
	log.Println("TFE_TOKEN environment variable is set.")

	return token, nil
}

func readTerraformTokenFile() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", errors.Wrap(err, "cannot locate home directory")
	}

	absolutePath := filepath.Join(homeDir, tfrcFileName)

	jsonFile, err := os.Open(absolutePath)
	if err != nil {
		return "", errors.Wrap(err, "cannot open terraform credentials file")
	}

	defer jsonFile.Close()

	var tfeConfig map[string]interface{}
	bytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return "", errors.Wrap(err, "cannot read terraform credentials file")
	}

	err = json.Unmarshal(bytes, &tfeConfig)
	if err != nil {
		log.Println("Cannot read terraform credentials file.")
	}

	tfeConfig, err = flatten.Flatten(tfeConfig, "", flatten.RailsStyle)
	if err == nil {
		token, ok := tfeConfig[fmt.Sprintf("credentials[%s][token]", terraformHostName)]
		if ok {
			return token.(string), nil
		}
	}

	log.Println("Cannot read a token from the tfrc file.")
	return "", errors.New("unable to read the TFE token")
}
