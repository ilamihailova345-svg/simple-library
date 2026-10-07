package main

import "errors"

func GetPortFromConfig(config map[string]string) (string, error) {
	port, norm := config["PORT"]
	if !norm {
		return "", errors.New("ключ PORT отсутствует в конфигурации")
	}
	return port, nil
}