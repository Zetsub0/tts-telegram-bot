package tts

import (
	htgotts "github.com/hegedustibor/htgo-tts"
	"log"
	"os"
)

func CreateSpeech(userID string, lang string, text string) error {
	speech := htgotts.Speech{
		Folder:   userID,
		Language: lang,
	}
	err := speech.Speak(text)
	if err != nil {
		return err
	}

	return nil
}

func GetFileName(folderPath string) string {
	entries, err := os.ReadDir(folderPath)
	if err != nil {
		log.Fatal(err)
	}
	res := entries[0].Name()
	return res
}

func DelFile(folderPath string) {
	os.RemoveAll(folderPath)
}
