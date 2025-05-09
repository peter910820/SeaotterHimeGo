package cmds

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
	"github.com/sirupsen/logrus"
)

func TextMessageEntryPoint(bot *messaging_api.MessagingApiAPI, e webhook.MessageEvent, message webhook.TextMessageContent) {
	var messages []messaging_api.MessageInterface

	message.Text = strings.TrimSpace(message.Text)

	if message.Text == "/test" {
		messages = append(messages, messaging_api.TextMessage{
			Text: "✅messaging_api 測試成功",
		})
	}

	if strings.Contains(message.Text, "查") {
		messages = append(messages, messaging_api.TextMessage{
			Text: "以下是查分器連結 請妥善使用～\nhttps://redive.estertion.win/arcaea/probe/",
		})
	}

	if strings.Contains(strings.ToLower(message.Text), "vc") {
		messages = append(messages, messaging_api.TextMessage{
			Text: "仙草快跟他結婚#",
		})
	}

	if strings.Contains(message.Text, "天堂門") {
		messages = append(messages, messaging_api.TextMessage{
			Text: "Snowth快去P!!!",
		})
	}

	if strings.Contains(strings.ToLower(message.Text), "運勢") {
		randomfortune := fortunate()
		messages = append(messages, messaging_api.TextMessage{
			Text: fmt.Sprintf("您今天的運勢: %s", randomfortune),
		})
	}

	reN := regexp.MustCompile(`(?i)^n\d{1,6}$`)
	if reN.MatchString(message.Text) {
		messages = append(messages, messaging_api.TextMessage{
			Text: fmt.Sprintf("https://nhentai.net/g/" + message.Text[1:]),
		})
	}

	_, err := bot.ReplyMessage(
		&messaging_api.ReplyMessageRequest{
			ReplyToken: e.ReplyToken,
			Messages:   messages,
		},
	)
	if err != nil {
		logrus.Error(err)
	} else {
		logrus.Info(fmt.Sprintf("使用者說: %s", message.Text))
	}
}
