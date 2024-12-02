package sound

import (
	"bytes"
	"fmt"
	"os"
	"time"

	"hello-go/zoopark/animal"

	"github.com/ebitengine/oto/v3"
	"github.com/hajimehoshi/go-mp3"
)

func PlaySound(animal animal.Animal) error {
	fileName := animal.SoundFile()

	// Проверка существования файла
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return fmt.Errorf("файл %s не найден", fileName)
	}

	fileBytes, err := os.ReadFile(fileName)
	if err != nil {
		return fmt.Errorf("ошибка чтения файла: %w", err)
	}

	fileBytesReader := bytes.NewReader(fileBytes)
	decodedMp3, err := mp3.NewDecoder(fileBytesReader)
	if err != nil {
		return fmt.Errorf("ошибка mp3 декодера: %w", err)
	}

	op := &oto.NewContextOptions{
		SampleRate:   44100,
		ChannelCount: 2,
		Format:       oto.FormatSignedInt16LE,
	}

	otoCtx, readyChan, err := oto.NewContext(op)
	if err != nil {
		return fmt.Errorf("ошибка создания контекста oto.NewContext: %w", err)
	}

	<-readyChan
	player := otoCtx.NewPlayer(decodedMp3)

	defer func() {
		if err := player.Close(); err != nil {
			fmt.Printf("ошибка закрытия проигрывателя: %v\n", err)
		}
	}()

	player.Play()

	for player.IsPlaying() {
		time.Sleep(time.Millisecond)
	}

	return nil
}
