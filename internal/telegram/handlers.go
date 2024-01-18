package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/zetsub0/tts-telegram-bot/internal/tts"
	"github.com/zetsub0/tts-telegram-bot/internal/whatlang"
	"log"
	"os"
	"strconv"
	"time"
)

const commandStart = "start"

func (b *Bot) handleCommand(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "I don't know that command")
	switch message.Command() {
	case commandStart:
		msg.Text = "Just send me text"
		tm := time.Unix(int64(message.Date), 0)
		log.Printf("[%s] %s \t[%s]", message.From.UserName, message.Text, tm)
		b.bot.Send(msg)
	default:
		tm := time.Unix(int64(message.Date), 0)
		log.Printf("[%s] %s \t[%s]", message.From.UserName, message.Text, tm)
		b.bot.Send(msg)
	}

}

func (b *Bot) handleMessage(message *tgbotapi.Message) error {

	userID := strconv.FormatInt(message.From.ID, 10)
	audioFolder := "./audio/" + userID + "/"
	tm := time.Unix(int64(message.Date), 0)
	log.Printf("[%s] %s \t[%s]", message.From.UserName, message.Text, tm)

	lang, err := whatlang.Analyze(message.Text)
	if err != nil {
		b.bot.Send(tgbotapi.NewMessage(message.Chat.ID, "Unable to determine the language"))
		return err
	}
	err = tts.CreateSpeech(audioFolder, lang, message.Text)
	if err != nil {
		//return err
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, "Creating audio. Please wait a moment.")
	b.bot.Send(msg)
	time.Sleep(time.Second * 10)
	audioName := audioFolder + tts.GetFileName(audioFolder)
	b.sendFile(audioName, message.From.ID)
	time.Sleep(time.Second * 10)

	tts.DelFile(audioFolder)

	return nil
}

func (b *Bot) sendFile(fileName string, chatID int64) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		log.Panic(err)
	}

	fileBytes := make([]byte, fileInfo.Size())
	_, err = file.Read(fileBytes)
	if err != nil {
		log.Panic(err)
	}

	fileData := tgbotapi.FileBytes{
		Name:  fileInfo.Name(),
		Bytes: fileBytes,
	}

	audioConfig := tgbotapi.NewVoice(chatID, fileData)

	_, err = b.bot.Send(audioConfig)
	if err != nil {
		log.Panic(err)
	}
}
