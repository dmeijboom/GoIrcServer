package main;

import (
    "fmt"
    "strings"
)

func Banlist(serverAddr string, channelName string, banId int) string {
    return fmt.Sprintf(":%v %v %v %v\n", serverAddr, "367", channelName, banId)
}

func EndOfBanlist(serverAddr string, channelName string) string {
    return fmt.Sprintf(":%v %v %v :End of channel ban list\n", serverAddr, "368", channelName)
}

func EndOfNames(serverAddr string, nickname string, channelName string) string {
    return fmt.Sprintf(":%v %v %v %v :End of /NAMES list\n", serverAddr, "366", nickname, channelName)
}

func EndOfWho(serverAddr string, query string) string {
    return fmt.Sprintf(":%v %v %v :End of WHO list\n", serverAddr, "315", query)
}

func InvalidCommand(serverAddr string, commandName string) string {
    return fmt.Sprintf(":%v %v %v :Unknown command\n", serverAddr, "421", commandName)
}

func IsOn(serverAddr string, nicknames []string) string {
    return fmt.Sprintf(":%v %v :%v\n", serverAddr, "303", strings.Join(nicknames, " "))
}

// NameReply sends a name reply command to the client, please note that
// the `=` character is being used to specify that it's a public channel
func NameReply(serverAddr string, channelName string, nickname string, nicknames []string) string {
    return fmt.Sprintf(":%v %v %v = %v :%v\n", serverAddr, "353", nickname, channelName, strings.Join(nicknames, " "))
}

func NoSuchChannel(serverAddr string, channelName string) string {
    return fmt.Sprintf(":%v %v %v :No such channel\n", serverAddr, "403", channelName)
}

func NoTopic(serverAddr string, channelName string) string {
    return fmt.Sprintf(":%v %v %v :No topic is set\n", serverAddr, "421", channelName)
}

func Topic(serverAddr string, channelName string, channelTopic string) string {
    return fmt.Sprintf(":%v %v %v :%v\n", serverAddr, "332", channelName, channelTopic)
}

func Welcome(serverAddr string, nickname string) string {
    return fmt.Sprintf(":%v %v %v Welcome to the Internet Relay Network\n", serverAddr, "001", nickname)
}

func WhoReply(serverAddr string, channelName string, username string, nickname string, realname string, addr string) string {
    return fmt.Sprintf(":%v %v %v ~%v %v %v %v H+ :0 %v\n", serverAddr, "352", channelName, username, addr, "localhost", nickname, realname)
}