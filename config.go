package models

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

//******************************************************************************

// DeploymentType enum type
type DeploymentType string

const (
	// DeploymentTypeDevelopment enum value
	DeploymentTypeDevelopment DeploymentType = "development"
	// DeploymentTypeProduction enum value
	DeploymentTypeProduction DeploymentType = "production"
)

//******************************************************************************

// Config type.
type Config struct {
	DeploymentType DeploymentType `json:"deployment-type"`
	DeploymentName string         `json:"deployment-name"`
	WebserverPort  string         `json:"webserver-port"`
	SQLHostname    string         `json:"sql-hostname"`
	SQLPort        string         `json:"sql-port"`
	SQLUsername    string         `json:"sql-username"`
	SQLPassword    string         `json:"sql-password"`
	SQLDatabase    string         `json:"sql-database"`
}

//******************************************************************************

// LoadConfig loads the json config file
func LoadConfig(path string, dtStr, dn string) (cfg Config, err error) {
	dt := DeploymentType(dtStr)

	fileData, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	if len(fileData) <= 0 {
		err = errors.New("config file is empty")
		return
	}

	type Configs []Config
	var configs Configs

	err = json.Unmarshal(fileData, &configs)
	if err != nil {
		return
	}

	found := false
	for _, obj := range configs {
		if obj.DeploymentType == dt && obj.DeploymentName == dn {
			cfg = obj
			found = true
			break
		}
	}

	if !found {
		err = errors.New("config type is not found")
		return
	}

	updateConfig(cfg)

	return
}

//******************************************************************************

func updateConfig(cfg Config) {
	if cfg.DeploymentType == DeploymentTypeDevelopment {
		return
	}

	deploymentType := os.Getenv("DEPLOYMENT_TYPE")
	deploymentName := os.Getenv("DEPLOYMENT_NAME")
	webserverPort := os.Getenv("WEBSERVER_PORT")
	sqlHostname := os.Getenv("HOST_NAME")
	sqlUsername := os.Getenv("USERNAME")
	sqlPassword := os.Getenv("PASSWORD")
	sqlDatabase := os.Getenv("DATABASE")

	if deploymentType != "" {
		cfg.DeploymentType = DeploymentType(deploymentType)
	}

	if deploymentName != "" {
		cfg.DeploymentName = deploymentName
	}

	if webserverPort != "" {
		cfg.WebserverPort = webserverPort
	}

	if sqlHostname != "" {
		cfg.SQLHostname = sqlHostname
	}

	if sqlUsername != "" {
		cfg.SQLUsername = sqlUsername
	}

	if sqlPassword != "" {
		cfg.SQLPassword = sqlPassword
	}

	if sqlDatabase != "" {
		cfg.SQLDatabase = sqlDatabase
	}
}

//******************************************************************************