package gossip

import (
    "bytes"
    "fmt"
)

type Method string

const (
    INVITE   Method = "INVITE"
    ACK      Method = "ACK"
    CANCEL   Method = "CANCEL"
    BYE      Method = "BYE"
    REGISTER Method = "REGISTER"
    OPTIONS  Method = "OPTIONS"
)

type SipMessage interface {
    String() (string)
}

type Request struct {
    Method Method
    Uri SipUri
    SipVersion string
    Headers []SipHeader
    Body *string
}
func (request *Request) String() (string) {
    var buffer bytes.Buffer

    // Every SIP request starts with a Request Line - RFC 2361 7.1.
    buffer.WriteString(fmt.Sprintf("%s %s %s\r\n",
        (string)(request.Method),
        request.Uri.String(),
        request.SipVersion))

    // Construct each header in turn and add it to the message.
    for idx, header := range(request.Headers) {
        buffer.WriteString(header.String())

        if (idx < len(request.Headers)) {
            buffer.WriteString("\r\n")
        }
    }

    // If the request has a message body, add it.
    if (request.Body != nil) {
        buffer.WriteString("\r\n" + *request.Body)
    }

    return buffer.String()
}

type Response struct {
    SipVersion string
    StatusCode uint8
    Reason string
    Headers []SipHeader
    Body *string
}
func (response *Response) String() (string) {
    var buffer bytes.Buffer

    // Every SIP response starts with a Status Line - RFC 2361 7.2.
    buffer.WriteString(fmt.Sprintf("%s %d %s\r\n",
        response.SipVersion,
        response.StatusCode,
        response.Reason))

    // Construct each header in turn and add it to the message.
    for idx, header := range(response.Headers) {
        buffer.WriteString(header.String())

        if (idx < len(response.Headers)) {
            buffer.WriteString("\r\n")
        }
    }

    // If the request has a message body, add it.
    if (response.Body != nil) {
        buffer.WriteString("\r\n" + *response.Body)
    }

    return buffer.String()
}
