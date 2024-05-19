package storage

import (
	"bufio"
	"os"
	"strings"
)

const (
	EmailFile = "storage-db/emails.txt"
)

func IsEmailExists(email string) bool {
	file, err := os.Open(EmailFile)
	if err != nil {
		return false
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.TrimSpace(scanner.Text()) == email {
			return true
		}
	}

	return false
}

func AddEmail(email string) error {
	file, err := os.OpenFile(EmailFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := file.WriteString(email + "\n"); err != nil {
		return err
	}
	return nil
}

func GetEmails() ([]string, error) {
	file, err := os.Open(EmailFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var emails []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		emails = append(emails, strings.TrimSpace(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return emails, nil
}
