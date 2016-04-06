package main;

import (
    "testing"
    "irc/assert"
)

// paramTest tests if the parameters are parsed correctly
func paramTest(test *testing.T, message *Message, params []string) {
    if lenTest := assert.EqualsInt(len(params), len(message.Params)); len(lenTest) > 0 {
        test.Errorf("Checking parameter length: %v", lenTest)
        
        // We need moar parameters to test this
        if len(message.Params) < len(params) {
            return
        }
    }
    
    for i := range params {
        if paramTest := assert.EqualsString(params[i], message.Params[i]); len(paramTest) > 0 {
            test.Errorf("Checking parameter %v: %v", params[i], paramTest)
        }
    }
}

// TestQuit tests if the message command is parsed when no other data
// is passed
func TestQuit(test *testing.T) {
    message, err := ParseMessage("QUIT")
    
    if err != nil {
        test.Errorf("Parsing message error: %v", err.Error())
    }
    
    if cmdTest := assert.EqualsString("QUIT", message.Command); len(cmdTest) > 0 {
        test.Errorf("Parsing command part: %v", cmdTest)
    }
    
    paramTest(test, message, []string{})
}

// TestWithPrefix tests if a message can be parsed with a prefix (both 
// with and without parameters)
func TestWithPrefix(test *testing.T) {
    message, err := ParseMessage(":prefixed QUIT")
    
    if err != nil {
        test.Errorf("Parsing message error: %v", err.Error())
    }
    
    if cmdTest := assert.EqualsString("QUIT", message.Command); len(cmdTest) > 0 {
        test.Errorf("Parsing command part: %v", cmdTest)
    }
    
    paramTest(test, message, []string{})
    
    if prefixTest := assert.EqualsString("prefixed", message.Prefix); len(prefixTest) > 0 {
        test.Errorf("Parsing prefix part: %v", prefixTest)
    }
    
    if hasPrefixTest := assert.EqualsBool(true, message.HasPrefix); len(hasPrefixTest) > 0 {
        test.Errorf("Parsing the prefix flag: %v", hasPrefixTest)
    }
}

// TestParams tests if the parameters are parsed correctly
func TestParams(test *testing.T) {
    message, err := ParseMessage("MODE #test b")
    
    if err != nil {
        test.Errorf("Parsing message error: %v", err.Error())
    }
    
    if cmdTest := assert.EqualsString("MODE", message.Command); len(cmdTest) > 0 {
        test.Errorf("Parsing command part: %v", cmdTest)
    }
    
    paramTest(test, message, []string{ "#test", "b" })
}

// TestTrailing tests if the trailing part is parsed correctly
func TestTrailing(test *testing.T) {
    message, err := ParseMessage("USER example 0 * :John von Appleseed")
    
    if err != nil {
        test.Errorf("Parsing message error: %v", err.Error())
    }
    
    if cmdTest := assert.EqualsString("USER", message.Command); len(cmdTest) > 0 {
        test.Errorf("Parsing command part: %v", cmdTest)
    }
    
    paramTest(test, message, []string{ "example", "0", "*" })
    
    if trailingTest := assert.EqualsString("John von Appleseed", message.Trailing); len(trailingTest) > 0 {
        test.Errorf("Parsing trailing part: %v", trailingTest)
    }
    
    if hasTrailingTest := assert.EqualsBool(true, message.HasTrailing); len(hasTrailingTest) > 0 {
        test.Errorf("Checking trailing flag: %v", hasTrailingTest)
    }
}

// TestPrefixAndTrailing tests if the trailing part is parsed correctly when a
// prefix is also defined
func TestPrefixAndTrailing(test *testing.T) {
    message, err := ParseMessage(":exam.pl USER example 0 * :John von Appleseed")
    
    if err != nil {
        test.Errorf("Parsing message error: %v", err.Error())
    }
    
    if cmdTest := assert.EqualsString("USER", message.Command); len(cmdTest) > 0 {
        test.Errorf("Parsing command part: %v", cmdTest)
    }
    
    paramTest(test, message, []string{ "example", "0", "*" })
    
    if prefixTest := assert.EqualsString("exam.pl", message.Prefix); len(prefixTest) > 0 {
        test.Errorf("Parsing prefix: %v", prefixTest)
    }
    
    if hasPrefixTest := assert.EqualsBool(true, message.HasPrefix); len(hasPrefixTest) > 0 {
        test.Errorf("Checking prefix flag: %v", hasPrefixTest)
    }
    
    if trailingTest := assert.EqualsString("John von Appleseed", message.Trailing); len(trailingTest) > 0 {
        test.Errorf("Parsing trailing part: %v", trailingTest)
    }
    
    if hasTrailingTest := assert.EqualsBool(true, message.HasTrailing); len(hasTrailingTest) > 0 {
        test.Errorf("Checking trailing flag: %v", hasTrailingTest)
    }
}

// TestRaw tests if the RAW representation of a message is generated correctly
func TestRaw(test *testing.T) {
    message, err := ParseMessage(":exam.pl USER example 0 * :John von Appleseed")
    if err != nil {
        test.Errorf("Parsing message error: %v", err.Error())
    }
    
    paramTest(test, message, []string{ "example", "0", "*" })
    
    if genTest := assert.EqualsString(":exam.pl USER example 0 * :John von Appleseed", message.Raw()); len(genTest) > 0 {
        test.Errorf("Generating RAW message: %v", genTest)
    }
}

// TestRawWithoutPrefix tests if the RAW representation of a message is generated correctly
func TestRawWithoutPrefix(test *testing.T) {
    message, err := ParseMessage("USER example 0 * :John von Appleseed")
    if err != nil {
        test.Errorf("Parsing message error: %v", err.Error())
    }
    
    paramTest(test, message, []string{ "example", "0", "*" })
    
    if genTest := assert.EqualsString("USER example 0 * :John von Appleseed", message.Raw()); len(genTest) > 0 {
        test.Errorf("Generating RAW message: %v", genTest)
    }
}

// BenchmarkParsing benchmarks how long it takes before a simple IRC message is parsed
func BenchmarkParsing(test *testing.B) {
    for i := 0; i < test.N; i++ {
        ParseMessage(":servername USER example 0 * :John von Appleseed")
    }
}

// BenchmarkParsingNoTrailing benchmarks how long it takes before a simple IRC message
// without a trailing part is parsed
func BenchmarkParsingNoTrailing(test *testing.B) {
    for i := 0; i < test.N; i++ {
        ParseMessage(":servername USER example 0 *")
    }
}

// BenchmarkParsingNoParams benchmarks how long it takes before a simple IRC message
// without a params part is parsed
func BenchmarkParsingNoParams(test *testing.B) {
    for i := 0; i < test.N; i++ {
        ParseMessage(":servername QUIT :John von Appleseed")
    }
}

// BenchmarkRaw benchmarks how long it takes before we have generated a RAW message
func BenchmarkRaw(test *testing.B) {
    for i := 0; i < test.N; i++ {
        msg, _ := ParseMessage(":servername QUIT :John von Appleseed")
        msg.Raw()
    }
}