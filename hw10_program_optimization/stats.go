package hw10programoptimization

import (
	"bufio"
	"io"
	"strings"

	"github.com/mailru/easyjson"
)

type User struct {
	Email string
}

type DomainStat map[string]int

func GetDomainStat(r io.Reader, domain string) (DomainStat, error) {
	domainStat := make(DomainStat)
	domain = "." + domain
	scanner := bufio.NewScanner(r)
	var user User
	for scanner.Scan() {
		if err := easyjson.Unmarshal(scanner.Bytes(), &user); err != nil {
			return nil, err
		}

		if strings.HasSuffix(user.Email, domain) {
			emailDomain := strings.ToLower(strings.SplitN(user.Email, "@", 2)[1])
			domainStat[emailDomain]++
		}
	}

	return domainStat, nil
}
