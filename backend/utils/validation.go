package utils

//----------------------------------Need more conidition for validation--------------------//

import (
	"errors"
	"strings"
	"Forum/backend/structs"
)



// ValidateInput checks for empty strings and strings with only whitespace
func ValidateInput(fields map[string]string) error {
	for key, value := range fields {
		if strings.TrimSpace(value) == "" {
			return errors.New(key + " cannot be empty or whitespace")
		}
	}
	return nil
}

// ValidateEmail checks if the email is in the correct format
func ValidateEmail(email string) error {
	if !strings.Contains(email, "@") {
		return errors.New("invalid email")
	}
	return nil
}

// ValidatePassword checks if the password is at least 8 characters long
func ValidatePassword(password string) error {
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}
	return nil
}

// ValidateCategory checks if the category is in the correct format
func ValidateCategory(category string, categories []structs.Category) error {
	if strings.TrimSpace(category) == ""  {
		return errors.New("category cannot be empty or whitespace")
	} else {
		for i := 0; i < len(categories); i++ {
			if strings.TrimSpace(category) == categories[i].Category {
				{
					return errors.New("invalid category")
				}
			}
		}

	}
	return nil
}

// ValidatePost checks if the post is in the correct format
func ValidatePost(title, content string, category []string) error {
	if strings.TrimSpace(title) == "" {
		return errors.New("title cannot be empty or whitespace")
	}
	if strings.TrimSpace(content) == "" {
		return errors.New("content cannot be empty or whitespace")
	}
	if len(category) == 0 {
		return errors.New("category cannot be empty")
	}
	return nil
}
