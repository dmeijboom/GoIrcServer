package main;

import (
    "strings"
    "errors"
)

// Message defines an IRC RAW message containing the following input:
// --
//
// (optional) prefix
// command
// parameters (maximum of 15)
// trailing
type Message struct {
    Prefix string
    HasPrefix bool
    Command string
    Params []string
    Trailing string
    HasTrailing bool
}

// MessageCutset strips newlines from the string, somehow it's twice
// as fast than using `strings.Trim()``
func MessageCutset(str rune) bool {
    return str == '\n'
}

// Raw returns the RAW representation of the IRC message
func (message *Message) Raw() string {
    raw := ""
    
    if message.HasPrefix {
        raw += ":"
        raw += message.Prefix
        raw += " "
    }
    
    raw += message.Command
    
    if len(message.Params) > 0 {
        for n := range message.Params {
            raw += " "
            raw += message.Params[n]
        }
    }
    
    if message.HasTrailing {
        raw += " :"
        raw += message.Trailing
    }
    
    return raw
}

// ParseMessage parses an IRC message using the specified format
func ParseMessage(data string) (*Message, error) {
    msg := Message{}
    
    // Data starts with a prefix
    if data[0] == ':' {
        index := strings.IndexByte(data, ' ')
        
        msg.HasPrefix = true
        msg.Prefix = data[1:index]
        
        data = data[index + 1:]
    }
    
    // Parse the command-part
    index := strings.IndexByte(data, ' ')
    
    if index == -1 {
        msg.Command = strings.TrimFunc(data, MessageCutset)
        return &msg, nil
    }
    
    msg.Command = strings.TrimFunc(data[0:index], MessageCutset)
    data = data[index + 1:]
    
    index = strings.LastIndexByte(data, ':')
    
    // Parse the trailing part
    if index > -1 {
        endIndex := strings.IndexByte(data, '\n')
        
        if endIndex == -1 {
            return &msg, errors.New("Expecting a newline at the end of a message, instead got: `" + data + "`")
        }
        
        msg.Trailing = data[index + 1:endIndex]
        msg.HasTrailing = true
        
        data = data[0:index]
    }
    
    // Parse the parameters
    for {
        index = strings.IndexByte(data, ' ')

        if index == -1 {
            // If the last one only contains a newline or is empty it should not be recognized
            // as a valid parameter
            if data != "" && data != "\n" {
                msg.Params = append(msg.Params, strings.TrimFunc(data, MessageCutset))
            }
            break
        }
        
        msg.Params = append(msg.Params, strings.TrimFunc(data[0:index], MessageCutset))
        
        if len(data) > index {
            data = data[index + 1:]
        } else {
            data = data[index:]
        }
    }
    
    return &msg, nil
}