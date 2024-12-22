# Message File Parser

A simple Go package for parsing structured XML message files and retrieving messages using a section:message key format.

## Installation

```bash
go get github.com/yourusername/messagefile
```

## Usage

The package provides a simple API to retrieve messages from an XML message file. The message file should be named `messagefile.xml` and be in the following format:

```xml
<?xml version="1.0" encoding="UTF-8"?>
<messages>
    <section_name>
        <message_type>
            Message content goes here
        </message_type>
    </section_name>
</messages>
```

### API Reference

#### GetMSG(messageKey string) (string, error)

Retrieves a message using a colon-separated section and message type key.

- `messageKey`: Format is "section:message"
- Returns: The message content and any error encountered

### Example

```go
package main

import (
    "fmt"
    "github.com/yourusername/messagefile"
)

func main() {
    // Retrieve a message
    msg, err := messagefile.GetMSG("utilmessages:query_rewrite")
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }
    
    fmt.Printf("Message: %s\n", msg)
}
```

### Message File Example

```xml
<?xml version="1.0" encoding="UTF-8"?>
<messages>
    <utilmessages>
        <query_rewrite>
            You are tasked with enhancing a user's query.
            Here is the user's original query:
            <user_query>
            %s
            </user_query>
        </query_rewrite>
    </utilmessages>
</messages>
```

## Error Handling

The package returns errors for:
- Invalid message key format
- Missing message file
- Invalid XML structure
- Section not found
- Message not found
