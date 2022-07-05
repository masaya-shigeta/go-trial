package validation

import (
    "fmt"

	"regexp"
	"strconv"
)

// check main
func CheckTextValidation(cardList []string) []string {
	messages := []string{}
	// count
	if (len(cardList) != 5) {
		messages = append(messages, `5つのカード指定文字を半角スペース区切りで入力してください。（例："S1 H3 D9 C13 S11"）`)
	}

	// duplicate
	if (len(messages) == 0 && isExistDuplicateCard(cardList)) {
		messages = append(messages, "カードが重複しています。")
	}

	// card type
	if (len(messages) == 0 ) {
		cardTypeMessages := checkCardType(cardList)
		if (len(cardTypeMessages) > 0) {
			messages = append(messages, cardTypeMessages...)
		}	
	}

	return messages
}

// check duplicate
func isExistDuplicateCard(cardList []string) bool {
	isDup := false
	for i, v := range cardList {
		for i2, v2 := range cardList {
			if (v == v2 && i != i2) {
				isDup = true
				break
			}
		}
		if (isDup) {
			break
		}
	}
    return isDup
}

// card type
func checkCardType(cardList []string) []string {
	messages := []string{}
	for i, v := range cardList {
		// suit and number
		suitReg := regexp.MustCompile(`[SHDC]`)
		number, _ := strconv.Atoi(v[1:])

		if (!suitReg.MatchString(v[0:1]) || (!(number >= 1 && number <= 13))) {
			messages = append(messages, fmt.Sprintf("%d番目のカード指定文字が不正です。", i))
		}
	}

	return messages
}
