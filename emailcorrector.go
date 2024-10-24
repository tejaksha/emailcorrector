package emailcorrector

import (
	"errors"
	"regexp"
	"strings"
)

// Email format validation regex
var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

// Common email domains
var commonDomains = []string{
	"gmail.com",
	"yahoo.com",
	"outlook.com",
	"hotmail.com",
	"icloud.com",
	"aol.com",
	"gmx.com",
	"zoho.com",
	"mail.com",
	"yandex.com",
	"protonmail.com",
	"fastmail.com",
	"tutanota.com",
	"lycos.com",
	"inbox.lv",
	"rediffmail.com",
	"qq.com",
	"163.com",
	"mail.ru",
	"seznam.cz",
	"tiscali.it",
	"bol.com.br",       // Brazil
	"o2.co.uk",         // UK
	"cox.net",          // USA
	"bell.net",         // Canada
	"sbcglobal.net",    // USA
	"telstra.com.au",   // Australia
	"optusnet.com.au",  // Australia
	"libero.it",        // Italy
	"wanadoo.fr",       // France
	"nordnet.fr",       // France
	"mailfence.com",    // Secure email
	"hushmail.com",     // Secure email
	"zohomail.com",     // India
	"gmx.de",           // Germany
	"fastmail.com",     // Australia
	"tutanota.com",     // Germany
	"hushmail.com",     // Secure email service
	"yahoo.co.uk",      // UK Yahoo
	"yahoo.fr",         // French Yahoo
	"hotmail.co.uk",    // UK Hotmail
	"outlook.co.uk",    // UK Outlook
	"me.com",           // Apple mail service
	"live.com",         // Microsoft Live
	"frontier.com",     // USA
	"t-online.de",      // Germany
	"blueyonder.co.uk", // UK
	"talktalk.net",     // UK
	"comcast.net",      // USA
	"virginmedia.com",  // UK
	"freenet.de",       // Germany
	"ntlworld.com",     // UK
	"skynet.be",        // Belgium
}

// ValidateEmail checks if the email format is valid
func ValidateEmail(email string) error {
	if !emailRegex.MatchString(email) {
		return errors.New("invalid email format")
	}
	return nil
}

// LevenshteinDistance computes the distance between two strings
func LevenshteinDistance(a, b string) int {
	var cost int
	d := make([][]int, len(a)+1)
	for i := range d {
		d[i] = make([]int, len(b)+1)
	}
	for i := 0; i <= len(a); i++ {
		d[i][0] = i
	}
	for j := 0; j <= len(b); j++ {
		d[0][j] = j
	}
	for i := 1; i <= len(a); i++ {
		for j := 1; j <= len(b); j++ {
			if a[i-1] == b[j-1] {
				cost = 0
			} else {
				cost = 1
			}
			d[i][j] = min(d[i-1][j]+1, min(d[i][j-1]+1, d[i-1][j-1]+cost))
		}
	}
	return d[len(a)][len(b)]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// SuggestDomainCorrection suggests a correction if the domain is a common typo
func SuggestDomainCorrection(domain string) string {
	minDistance := len(domain)
	suggestedDomain := domain

	for _, commonDomain := range commonDomains {
		distance := LevenshteinDistance(domain, commonDomain)
		if distance < minDistance {
			minDistance = distance
			suggestedDomain = commonDomain
		}
	}
	return suggestedDomain
}

// CorrectEmail checks the email and suggests corrections
func CorrectEmail(email string) (string, error) {
	if err := ValidateEmail(email); err != nil {
		return "", err
	}

	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return "", errors.New("invalid email format")
	}

	localPart := parts[0]
	domainPart := SuggestDomainCorrection(parts[1])

	correctedEmail := localPart + "@" + domainPart
	return correctedEmail, nil
}
