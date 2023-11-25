package app

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

// Function to calculate Levenshtein distance between two strings
func LevenshteinDistance(s1, s2 string) int {
	m := len(s1)
	n := len(s2)

	// Create a matrix to store distances
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			// If one of the strings is empty, the distance is the length of the other string
			if i == 0 {
				dp[i][j] = j
			} else if j == 0 {
				dp[i][j] = i
			} else if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1])
			}
		}
	}

	return dp[m][n]
}

// Function to find similar strings in a slice based on a threshold distance and return their indices and titles
func FindSimilarWords(key string, data []string, threshold int) map[int]string {
	result := make(map[int]string)

	for i, str := range data {
		distance := LevenshteinDistance(strings.ToLower(key), strings.ToLower(str))
		if distance <= threshold {
			result[i] = str
		}
	}

	return result
}

func FindSimilarStrings(searchKey string, data []string, treshbold int) map[int]string {
	result := make(map[int]string)
	for index, title := range data {
		words := strings.Fields(title)

		// Finding similar strings for each word in the title
		similarTitles := FindSimilarWords(strings.ToLower(searchKey), words, 2)
		if len(similarTitles) > 0 {
			result[index] = title
		}
	}
	return result
}

func min(a, b, c int) int {
	if a <= b && a <= c {
		return a
	} else if b <= a && b <= c {
		return b
	}
	return c
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
