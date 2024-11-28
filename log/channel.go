package log

import (
	"bytes"
	"errors"
	"fmt"
)

var errUnmarshalNilChannel = errors.New("can't unmarshal a nil *Channel")

// Channel 是一个表示不同日志通道的类型。
type Channel int8

const (
	ConsoleChannel Channel = iota
	FileChannel
	LarkChannel     // 飞书
	DingTalkChannel // 钉钉
	WeComChannel    // 企业微信
	Server3Channel  // server3酱
)

// ParseChannel 解析一个通道的字符串表示。如果提供的字符串无效，则返回错误。
func ParseChannel(text string) (Channel, error) {
	var channel Channel
	err := channel.UnmarshalText([]byte(text))
	return channel, err
}

// String 返回通道的字符串表示。
func (c Channel) String() string {
	switch c {
	case ConsoleChannel:
		return "console"
	case FileChannel:
		return "file"
	case LarkChannel:
		return "lark"
	case DingTalkChannel:
		return "dingtalk"
	case WeComChannel:
		return "wecom"
	case Server3Channel:
		return "server3"
	default:
		return fmt.Sprintf("Channel(%d)", c)
	}
}

// MarshalText 将通道序列化为文本。
func (c Channel) MarshalText() ([]byte, error) {
	return []byte(c.String()), nil
}

// UnmarshalText 将文本反序列化为通道。
func (c *Channel) UnmarshalText(text []byte) error {
	if c == nil {
		return errUnmarshalNilChannel
	}
	if !c.unmarshalText(text) && !c.unmarshalText(bytes.ToLower(text)) {
		return fmt.Errorf("unrecognized channel: %q", text)
	}
	return nil
}

// unmarshalText 是一个辅助方法，用于将文本反序列化为通道。
func (c *Channel) unmarshalText(text []byte) bool {
	switch string(text) {
	case "console":
		*c = ConsoleChannel
	case "file":
		*c = FileChannel
	case "lark":
		*c = LarkChannel
	case "dingtalk":
		*c = DingTalkChannel
	case "wecom":
		*c = WeComChannel
	case "server3":
		*c = Server3Channel
	default:
		return false
	}
	return true
}

// Set 设置通道的值。
func (c *Channel) Set(s string) error {
	return c.UnmarshalText([]byte(s))
}

// Get 获取通道的值。
func (c *Channel) Get() interface{} {
	return *c
}

// Enabled 返回给定通道是否有效。
func (c Channel) Enabled(ch Channel) bool {
	return ch >= c
}

// ChannelEnabler 决定给定通道是否启用。
type ChannelEnabler interface {
	Enabled(Channel) bool
}
