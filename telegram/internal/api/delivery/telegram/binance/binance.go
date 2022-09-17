/*
	copyright © 2022 by @https://github.com/Chuiko-GIT
*/

package binance

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"io"
	"net/http"
	"telegram/internal/config"
)

// handler - define telegram handler.
type handler struct {
	bot    *tgbotapi.BotAPI
	config config.Binance
}

// NewHandler - constructor telegram handler.
func NewHandler(bot *tgbotapi.BotAPI, config config.Binance) *handler {
	return &handler{bot: bot, config: config}
}

// CryptoCourseToUSDT - handle crypto course to USDT command.
func (h handler) CryptoCourseToUSDT(message *tgbotapi.Message) error {
	resp, err := http.Get(toBinanceCourseURL(h.config.BinanceCourseURL, message.Text, CryptoTypeUSDT))
	if err != nil {
		return err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	var result binanceResponseCourseToUSDT
	if err = json.Unmarshal(body, &result); err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf(
		"На данный момент 1 %s равен %s %s", message.Text, result.Price, CryptoTypeUSDT),
	)

	if _, err = h.bot.Send(msg); err != nil {
		return err
	}

	return nil
}
