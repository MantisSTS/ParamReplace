package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/url"
	"os"
	"strings"
	"sync"
)

func main() {

	// Define command-line flags
	paramToReplace := flag.String("param", "", "Parameter to replace")
	replacementValue := flag.String("value", "", "Replacement value")
	flag.Parse()

	// Check if the required flags are provided
	if *paramToReplace == "" || *replacementValue == "" {
		fmt.Println("Usage: cat urls.txt | paramreplace -param=foo -value=bar")
		os.Exit(1)
	}

	inputURLChannel := make(chan string, 100)

	var wg sync.WaitGroup

	// read from stdin
	wg.Add(1)
	go func() {
		defer wg.Done()
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			inputURLChannel <- strings.TrimSpace(scanner.Text())
		}
		close(inputURLChannel)
	}()

	// Iterate through the input URLs
	for inputURL := range inputURLChannel {

		// Parse the input URL
		parsedURL, err := url.Parse(inputURL)
		if err != nil {
			fmt.Println("Error parsing input URL:", err)
		}

		// Get the query parameters
		queryParameters, err := url.ParseQuery(parsedURL.RawQuery)
		if err != nil {
			fmt.Println("Error parsing query parameters:", err)
		}

		// Convert the flag value to lowercase for case-insensitive matching
		paramToReplaceLower := strings.ToLower(*paramToReplace)

		// Iterate through the parameter names and check if the flag value (case-insensitive) is a substring of the parameter name
		var matchingParamName string
		for paramName := range queryParameters {
			if strings.Contains(strings.ToLower(paramName), paramToReplaceLower) {
				matchingParamName = paramName
				break
			}
		}

		// If a matching parameter is found, replace it with the new value
		if matchingParamName != "" {
			queryParameters.Set(matchingParamName, *replacementValue)
			parsedURL.RawQuery = queryParameters.Encode()
			fmt.Println(parsedURL.String())
		} else {
			// fmt.Println("No matching parameter found.")
		}
	}

	wg.Wait()
}
