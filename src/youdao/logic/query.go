package logic

import (
	"alfred/workflow/output"
	"alfred/workflow/util"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/url"
	"strconv"
	"strings"
)

type Req struct {
	From   string `json:"from"`
	To     string `json:"to"`
	AppKey string `json:"appKey"`
	Word   string `json:"q"`
	Salt   string `json:"salt"`
	Sign   string `json:"sign"`
}

type Resp struct {
	Query       string        `json:"query"`
	Code        string        `json:"errorCode"`
	Convert     string        `json:"l"`
	Translation []string      `json:"translation"`
	Phrases     []*RespPhrase `json:"web"`
	Basic       *RespBasic    `json:"basic"`
}

type RespPhrase struct {
	Values []string `json:"value"`
	Key    string   `json:"key"`
}

type RespBasic struct {
	Phonetic   string   `json:"phonetic"`
	UsPhonetic string   `json:"us-phonetic"`
	UkPhonetic string   `json:"uk-phonetic"`
	Explains   []string `json:"explains"`
}

func (r *Resp) wordItem(word, explain, icon string) (item *output.Item) {
	var readWord, readExplain string
	soundWord, soundExplain := "🔊"+word, "🔊"+explain

	if r.Convert == ConvertCn2En {
		readWord = "🔊🇨🇳" + word
		readExplain = "🔊🇺🇸" + explain
	} else {
		readWord = "🔊🇺🇸" + word
		readExplain = "🔊🇨🇳" + explain
	}

	copyText := word + "\n" + explain

	return &output.Item{
		Valid:        true,
		Title:        explain,
		SubTitle:     word,
		Arguments:    copyText,
		QuicklookUrl: QuicklookUrl + url.QueryEscape(r.Query),
		Icon:         output.NewIcon(icon),
		Mods: map[string]*output.Mod{
			output.ModKeyCmd: {Valid: true, Arguments: readWord, SubTitle: soundWord},
			output.ModKeyAlt: {Valid: true, Arguments: readExplain, SubTitle: soundExplain},
		},
		Text: &output.Text{Copy: copyText, LargeType: copyText},
	}
}

func parseData(content []byte) (items []*output.Item, err error) {
	resp := &Resp{}

	if err = json.Unmarshal(content, resp); err != nil {
		return
	}

	if resp.Code != CodeSuccess {
		err = fmt.Errorf("invalid response")
		return
	}

	items = make([]*output.Item, 0, 8)
	existExplains := map[string]bool{}

	if resp.Basic != nil {
		var phonetics []string

		if resp.Basic.UsPhonetic != "" || resp.Basic.UkPhonetic != "" {
			if resp.Basic.UsPhonetic != "" {
				phonetics = append(phonetics, fmt.Sprintf("[美: %s]", resp.Basic.UsPhonetic))
			}

			if resp.Basic.UkPhonetic != "" {
				phonetics = append(phonetics, fmt.Sprintf("[英: %s]", resp.Basic.UkPhonetic))
			}
		} else if resp.Basic.Phonetic != "" {
			phonetics = append(phonetics, fmt.Sprintf("[%s]", resp.Basic.Phonetic))
		}

		if len(phonetics) > 0 {
			explain := strings.Join(phonetics, " ")
			items = append(items, resp.wordItem(resp.Query, explain, IconTranslateSay))
			existExplains[explain] = true
		}
	}

	if len(resp.Translation) > 0 {
		for _, explain := range resp.Translation {
			if existExplains[explain] {
				continue
			}

			items = append(items, resp.wordItem(resp.Query, explain, IconTranslate))
			existExplains[explain] = true
		}
	}

	if resp.Basic != nil && len(resp.Basic.Explains) > 0 {
		for _, explain := range resp.Basic.Explains {
			if existExplains[explain] {
				continue
			}

			items = append(items, resp.wordItem(resp.Query, explain, IconTranslate))
			existExplains[explain] = true
		}
	}

	if len(resp.Phrases) > 0 {
		for _, phrase := range resp.Phrases {
			explain := strings.Join(phrase.Values, ", ")

			if existExplains[explain] {
				continue
			}

			items = append(items, resp.wordItem(phrase.Key, explain, IconTranslate))
			existExplains[explain] = true
		}
	}

	return
}

func queryWord(appKey, secret, word string) (items []*output.Item, err error) {
	salt := strconv.Itoa(rand.Intn(100000))
	sign := fmt.Sprintf("%x", md5.Sum([]byte(appKey+word+salt+secret)))

	params := url.Values{}
	params.Add("from", LangAuto)
	params.Add("to", LangAuto)
	params.Add("appKey", appKey)
	params.Add("q", word)
	params.Add("salt", salt)
	params.Add("sign", sign)

	uri := ApiUri + "?" + params.Encode()
	body, err := util.HttpRequest(uri, nil, nil)

	if err != nil {
		return
	}

	return parseData(body)
}
