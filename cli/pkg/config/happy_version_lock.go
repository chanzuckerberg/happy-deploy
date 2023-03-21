package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

type HappyVersionLockFile struct {
	HappyVersion string
	Path         string
}

func NewHappyVersionLockFile(projectRoot string, requiredVersion string) (*HappyVersionLockFile, error) {

	if projectRoot == "" {
		return nil, errors.New("No projectRoot specified")
	}

	if requiredVersion == "" {
		return nil, errors.New("No requiredVersion specified")
	}

	path := calcHappyVersionPath(projectRoot)

	return &HappyVersionLockFile{
		HappyVersion: requiredVersion,
		Path:         path,
	}, nil
}

func DoesHappyVersionLockFileExist(projectRoot string) bool {
	filePath := calcHappyVersionPath(projectRoot)
	_, err := os.Stat(filePath)
	return err == nil
}

func LoadHappyVersionLockFile(projectRoot string) (*HappyVersionLockFile, error) {

	filePath := calcHappyVersionPath(projectRoot)

	versionFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer versionFile.Close()

	contents, err := ioutil.ReadAll(versionFile)
	if err != nil {
		return nil, err
	}

	hvlf := HappyVersionLockFile{}

	err = json.Unmarshal(contents, &hvlf)
	if err != nil {
		return nil, err
	}

	return &hvlf, nil
}

func (v *HappyVersionLockFile) Save() error {

	if v.Path == "" {
		return errors.New("Path is not set!")
	}

	happyVersionFile, err := os.Create(v.Path)

	if err != nil {
		return errors.New(fmt.Sprintf("Could not create %s: %v", v.Path, err))
	}

	contents, err := json.MarshalIndent(&v, "", " ")
	if err != nil {
		return err
	}

	happyVersionFile.WriteString(string(contents))
	happyVersionFile.Close()

	return nil
}

func calcHappyVersionPath(projectRoot string) string {
	return filepath.Join(projectRoot, ".happy", "version.lock")
}
