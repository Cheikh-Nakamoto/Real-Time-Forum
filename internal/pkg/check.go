package pkg

import (
	"regexp"
	"strings"
)

func CheckValues(values []string, arg uint) bool {
	for _, value := range values {
		if len(value) < 3 {
			return false
		}
	}
	return arg >= 12
}

// La fonction retourne true si l'email est valide
// Check if the email is valid
func CheckEmail(email string) bool {
	if len(email) >= 10 && len(email) <= 35 {
		parts := strings.Split(email, "@")
		// Vérifier s'il y a bien 2 parties dans l'email
		// Check if there are 2 parts in the email string
		if len(parts) != 2 || len(strings.TrimSpace(parts[0])) == 0 || len(strings.TrimSpace(parts[1])) == 0 {
			return false
		}

		// Vérifier si le nom d'utilisateur est suffisamment long
		// Check if the username is long enough
		username := parts[0]
		regUsername := regexp.MustCompile(`^[a-z0-9]+$`)
		if len(username) < 3 || len(username) > 20 || !regUsername.MatchString(username) {
			return false
		}

		// Vérifier si le nom de domaine est correct
		// Check if the domain name is correct
		domainName := strings.Split(parts[1], ".")
		if len(domainName) != 2 || (len(strings.TrimSpace(domainName[0])) < 3 || len(strings.TrimSpace(domainName[0])) > 10) || (len(strings.TrimSpace(domainName[1])) < 2 || len(strings.TrimSpace(domainName[1])) > 3) {
			return false
		}

		// Si tout est dans les normes, alors l'email est valide
		// If everything is in order, then the email is valid
		return true
	}
	return false
}

// La fonction retourne true si le mot de passe est valide
// Return true if the password is valid
func CheckPassword(password string) bool {
	if len(password) >= 8 && len(password) <= 20 {
		hasUppercase := regexp.MustCompile(`[A-Z]+?`).MatchString(password)
		hasLowercase := regexp.MustCompile(`[a-z]+?`).MatchString(password)
		hasDigit := regexp.MustCompile(`[0-9]+?`).MatchString(password)
		hasSpecialChar := regexp.MustCompile(`[\W_]+?`).MatchString(password)
		return hasUppercase && hasLowercase && hasDigit && hasSpecialChar
	}
	return false
}

// La fonction retourne true si le nom d'utilisateur est valide
// Return true if the username is valid
func CheckNickname(nickname string) bool {
	if len(nickname) >= 3 && len(nickname) <= 20 {
		hasLowercase := regexp.MustCompile(`[a-z]+?`).MatchString(nickname)
		hasDigit := regexp.MustCompile(`[0-9]*`).MatchString(nickname)
		hasUppercase := regexp.MustCompile(`[A-Z]+?`).MatchString(nickname)
		hasBeginWithDigit := regexp.MustCompile(`^[0-9]+?`).MatchString(nickname)
		hasContainsSpecialChar := regexp.MustCompile(`[\W_]+?`).MatchString(nickname)
		return hasLowercase && hasDigit && !hasUppercase && !hasBeginWithDigit && !hasContainsSpecialChar
	}
	return false
}

// La fonction retourne true si le nom est valide
// Return true if the name is valid
func CheckName(name string) bool {
	if len(name) >= 3 && len(name) <= 20 {
		hasSpecialChar := regexp.MustCompile(`^([a-zA-Z\s-]+?)$`).MatchString(name)
		return hasSpecialChar
	}
	return false
}

// La fonction retourne true si l'âge est valide
// Return true if the age is valid
func CheckAge(age uint) bool {
	return age >= 12 && age <= 120
}

// La fonction retourne true si le genre est valide
// Return true if the gender is valid
func CheckGender(gender string) bool {
	// if len(gender) < 4 || len(gender) > 10 {
	// 	return false
	// }
	return gender == "male" || gender == "female"
}

// La fonction retourne true si le titre est valide
// Return true if the title is valid
func CheckTitle(title string) bool {
	if len(title) < 15 || len(title) > 100 {
		return false
	}

	return true
}

// La fonction retourne true si la catégorie est valide
// Return true if the category is valid
func CheckCategory(category string) bool {
	/*if len(category) < 3 || len(category) > 20 {
		return false
	}*/
	return regexp.MustCompile(`^[a-zA-Z0-9\s\W_]{3,20}$`).MatchString(category)
}
