package messagefile

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func GetMSG(messageKey string) (string, error) {
	// Split the message key
	parts := regexp.MustCompile(`:`).Split(messageKey, 2)
	if len(parts) != 2 {
		return "", fmt.Errorf("invalid message key format. Expected 'section:message', got '%s'", messageKey)
	}

	section := parts[0]
	messageType := parts[1]

	// Read the XML file
	data, err := os.ReadFile("messagefile.xml")
	if err != nil {
		return "", fmt.Errorf("error reading messagefile.xml: %w", err)
	}

	// Convert to string and clean up newlines for consistency
	xmlContent := string(data)

	// Skip past standard elements
	standardHeader := `<?xml version="1.0" encoding="UTF-8"?>
<messages>`

	if !strings.HasPrefix(xmlContent, standardHeader) {
		return "", fmt.Errorf("invalid message file format: missing standard header")
	}

	// Get content after standard elements
	xmlContent = xmlContent[len(standardHeader):]

	// Pattern for section (handles multiline with (?s))
	sectionRegex := regexp.MustCompile(fmt.Sprintf(`(?s)<%s>(.*?)</%s>`, section, section))
	sectionMatch := sectionRegex.FindStringSubmatch(xmlContent)
	if len(sectionMatch) < 2 {
		return "", fmt.Errorf("section not found: %s", section)
	}
	sectionContent := sectionMatch[1]

	// Pattern for message type within section
	messageRegex := regexp.MustCompile(fmt.Sprintf(`(?s)<%s>(.*?)</%s>`, messageType, messageType))
	messageMatch := messageRegex.FindStringSubmatch(sectionContent)
	if len(messageMatch) < 2 {
		return "", fmt.Errorf("message not found: %s", messageType)
	}

	// Trim the message content and normalize newlines
	messageContent := strings.TrimSpace(messageMatch[1])

	return messageContent, nil
}

/*func main() {
    tests := []string{
        "utilmessages:query_rewrite",
    }

    for _, test := range tests {
        fmt.Printf("\nTesting: %s\n", test)
        fmt.Println(strings.Repeat("-", 50))

        msg, err := GetMSG(test)
        if err != nil {
            log.Printf("Error: %v\n", err)
            continue
        }

        fmt.Printf("Message:\n%s\n", msg)
    }
}
*/
