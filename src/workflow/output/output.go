package output

const (
	IconError   = "icon/error.png"
	IconDefault = "icon.png"
)

const (
	TypeDefault     = "default"
	TypeFile        = "file"
	TypeFileNoCheck = "file:skipcheck"
)

const (
	IconTypeFileIcon = "fileicon"
	IconTypeFileType = "filetype"
)

const (
	ModKeyAlt = "alt"
	ModKeyCmd = "cmd"
)

type Icon struct {
	Type string `json:"type,omitempty"`
	Path string `json:"path"`
}

type Mod struct {
	Valid     bool   `json:"valid"`
	Arguments string `json:"arg,omitempty"`
	SubTitle  string `json:"subtitle,omitempty"`
	Icon      *Icon  `json:"icon,omitempty"`
}

type Text struct {
	Copy      string `json:"copy,omitempty"`
	LargeType string `json:"largetype,omitempty"`
}

type Item struct {
	Uid          string          `json:"uid,omitempty"`
	Type         string          `json:"type,omitempty"`
	Valid        bool            `json:"valid"`
	Title        string          `json:"title"`
	SubTitle     string          `json:"subtitle,omitempty"`
	Arguments    string          `json:"arg,omitempty"`
	AutoComplete string          `json:"autocomplete,omitempty"`
	Icon         *Icon           `json:"icon"`
	QuicklookUrl string          `json:"quicklookurl,omitempty"`
	Match        string          `json:"match,omitempty"`
	Mods         map[string]*Mod `json:"mods,omitempty"`
	Text         *Text           `json:"text,omitempty"`
}

type Result struct {
	Items []*Item `json:"items"`
}

func NewNotice(message string, err error) (res *Result) {
	item := &Item{Valid: true, Title: message, Icon: &Icon{Path: IconDefault}}

	if err != nil {
		item.SubTitle = err.Error()
		item.Icon.Path = IconError
	}

	return &Result{Items: []*Item{item}}
}
