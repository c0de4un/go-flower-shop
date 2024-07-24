package bot

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"

	"github.com/c0de4un/go-flower-shop/internal/logging"
)

type TelegramConfig struct {
	XMLName xml.Name `xml:"TelegramConfig"`

	Token  string `xml:"Token"`
	AppUrl string `xml:"AppUrl"`
}

func LoadTGConfig(filePath string) (*TelegramConfig, error) {
	logging.GetLogger().Debug(fmt.Sprintf("TelegramConfig::LoadTGConfig: reading %s", filePath))

	f, err := os.Open(filePath)
	if err != nil {
		logging.GetLogger().Error(fmt.Sprintf("TelegramConfig::LoadTGConfig: %v", err))

		return nil, err
	}
	defer f.Close()
	decoder := xml.NewDecoder(f)

	cfg := &TelegramConfig{
		Token:  "",
		AppUrl: "",
	}

	for {
		token, err := decoder.Token()
		if err == io.EOF {
			break
		}

		if err != nil {
			logging.GetLogger().Error(fmt.Sprintf("TelegramConfig::LoadTGConfig: %v", err))

			return nil, err
		}
		if token == nil {
			break
		}

		switch tokenType := token.(type) {
		case xml.StartElement:
			if tokenType.Name.Local == "TelegramConfig" {
				err = decoder.DecodeElement(&cfg, &tokenType)
				if err != nil {
					logging.GetLogger().Error(fmt.Sprintf("TelegramConfig::LoadTGConfig: %v", err))

					return nil, err
				}

				break
			}
		}
	}

	return cfg, nil
}
