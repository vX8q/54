package encoding

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Yandex-Practicum/final-project-encoding-go/models"
	"gopkg.in/yaml.v3"
)

type JSONData struct {
	DockerCompose *models.DockerCompose
	FileInput     string
	FileOutput    string
}

type YAMLData struct {
	DockerCompose *models.DockerCompose
	FileInput     string
	FileOutput    string
}

type MyEncoder interface {
	Encoding() error
}

func (j *JSONData) Encoding() error {
	data, err := ioutil.ReadFile(j.FileInput)
	if err != nil {
		return fmt.Errorf("ошибка при чтении файла JSON: %v", err)
	}

	err = json.Unmarshal(data, &j.DockerCompose)
	if err != nil {
		return fmt.Errorf("ошибка при декодировании JSON: %v", err)
	}

	yamlData, err := yaml.Marshal(j.DockerCompose)
	if err != nil {
		return fmt.Errorf("ошибка при кодировании в YAML: %v", err)
	}

	err = ioutil.WriteFile(j.FileOutput, yamlData, os.ModePerm)
	if err != nil {
		return fmt.Errorf("ошибка при записи в файл YAML: %v", err)
	}

	return nil
}

func (y *YAMLData) Encoding() error {
	data, err := ioutil.ReadFile(y.FileInput)
	if err != nil {
		return fmt.Errorf("ошибка при чтении файла YAML: %v", err)
	}

	err = yaml.Unmarshal(data, &y.DockerCompose)
	if err != nil {
		return fmt.Errorf("ошибка при декодировании YAML: %v", err)
	}

	jsonData, err := json.MarshalIndent(y.DockerCompose, "", "  ")
	if err != nil {
		return fmt.Errorf("ошибка при кодировании в JSON: %v", err)
	}

	err = ioutil.WriteFile(y.FileOutput, jsonData, os.ModePerm)
	if err != nil {
		return fmt.Errorf("ошибка при записи в файл JSON: %v", err)
	}

	return nil
}
