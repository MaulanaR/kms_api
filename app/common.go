package app

import (
	"fmt"
	"strings"

	"github.com/sahilm/fuzzy"
	"golang.org/x/net/html"
	"gopkg.in/gomail.v2"
)

func FindSimilarStrings(searchKey string, data []string) map[int]string {
	result := make(map[int]string)
	lowercaseSlice := toLower(data)
	matches := fuzzy.Find(strings.ToLower(searchKey), lowercaseSlice)
	for _, match := range matches {
		result[match.Index] = data[match.Index]
	}
	return result
}

func toLower(slice []string) []string {
	for i := range slice {
		slice[i] = strings.ToLower(slice[i])
	}
	return slice
}

// Function to remove HTML tags from a string
func RemoveHTMLTags(input string) string {
	// Parsing the HTML string
	doc, err := html.Parse(strings.NewReader(input))
	if err != nil {
		fmt.Println("Error parsing HTML:", err)
		return input
	}

	// Function to traverse the HTML structure and extract text
	var result string
	var extractText func(*html.Node)
	extractText = func(n *html.Node) {
		if n.Type == html.TextNode {
			result += n.Data + " "
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			extractText(c)
		}
	}

	extractText(doc)
	return strings.TrimSpace(result)
}

func SendMail(to, subject, message string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", GMAIL_NAME+"<"+GMAIL_AUTH+">")
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", message)

	// Send the email to Bob
	d := gomail.NewDialer(GMAIL_HOST, GMAIL_PORT, GMAIL_AUTH, GMAIL_PASS)
	err := d.DialAndSend(m)
	if err != nil {
		return err
	}

	return nil
}
