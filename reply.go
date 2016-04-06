package main;

const (
    ReplyWelcome = "001"
    ReplyYourhost = "002"
    ReplyCreated = "003"
    ReplyMyinfo = "004"
    ReplyBounce = "005"
    ReplyAway = "301"
    ReplyUserhost = "302"
    ReplyIson = "303"
    ReplyNoaway = "306"
    ReplyWhoisUser = "311"
    ReplyWhoisServer = "312"
    ReplyWhoisOperator = "313"
    ReplyWhoWhatUser = "314"
    ReplyWhoisIdle = "317"
    ReplyEndOfWhois = "318"
    ReplyWhoisChannels = "319"
    ReplyList = "322"
    ReplyListEnd = "323"
    ReplyChannelModeIs = "324"
    ReplyUniqueOperatorIs = "325"
    ReplyNoTopic = "331"
    ReplyTopic = "332"
    ReplyInviting = "341"
    ReplySummoning = "342"
    ReplyInvitelist = "346"
    ReplyEndOfInviteList = "347"
    ReplyExceptList = "348"
    ReplyEndOfExceptList = "349"
    ReplyVersion = "351"
    ReplyEndOfWhoWhat = "369"
    ReplyWhoReply = "352"
    ReplyEndOfWho = "315"
    ReplyNameReply = "353"
    ReplyEndOfNames = "366"
    ReplyLinks = "364"
    ReplyEndOfLinks = "365"
    ReplyBanlist = "367"
    ReplyEndOfBanlist = "368"
    ReplyInfo = "371"
    ReplyEndOfInfo = "374"
    ReplyMotdStart = "375"
    ReplyMotd = "372"
    ReplyEndOfMotd = "376"
    ReplyYourOperator = "381"
    ReplyTime = "391"
    ReplyUsersStart = "392"
    ReplyUsers = "393"
    ReplyEndOfUsers = "394"
    ReplyNoUsers = "395"
    ErrorNoSuchNick = "401"
    ErrorNoSuchServer = "402"
    ErrorNoSuchChannel = "403"
    ErrorCannotSendToChannel = "404"
    ErrorTooManyChannels = "405"
    ErrorWasNoSuchNick = "406"
    ErrorTooManyTargets = "407"
    ErrorNoSuchService = "408"
    ErrorNoSuchOrigin = "409"
    ErrorNoRecipient = "411"
    ErrorNoTextToSend = "412"
    ErrorNoTopLevel = "413"
    ErrorWildTopLevel = "414"
    ErrorBadMask = "415"
    ErrorUnknownCommand = "421"
)

type Reply struct {
    Code string
    Prefix string
    Params []string
    Trailing string
}

func (reply *Reply) Raw() string {
    raw := ""
    
    if len(reply.Prefix) > 0 {
        raw += ":"
        raw += reply.Prefix
        raw += " "
    }
    
    if len(reply.Code) > 0 {
        raw += reply.Code
    }
    
    for i := range reply.Params {
        if i == 0 && len(reply.Code) == 0 {
            raw += reply.Params[i]
            continue
        }
        
        raw += " "
        raw += reply.Params[i]
    }
    
    if len(reply.Trailing) > 0 {
        raw += " "
        raw += ":"
        raw += reply.Trailing
    }
    
    return raw
}