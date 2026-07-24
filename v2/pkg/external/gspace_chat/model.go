package notify_error

type Request struct {
	Error  interface{} `json:"error"`
	Id     string      `json:"id"`
	Secret string      `json:"secret"`
	Token  string      `json:"token"`
}

type Response struct {
	Message string `json:"message"`
}

type NotifyRequest struct {
	Text string `json:"text"`
	Card
	Annotation
}

type Annotation struct {
	Annotations []Annotations `json:"annotations"`
}

type Annotations struct {
	Type        string      `json:"type"`
	StartIndex  int         `json:"startIndex"`
	Length      int         `json:"length"`
	UserMention UserMention `json:"userMention"`
}

type UserMention struct {
	User User   `json:"user"`
	Type string `json:"type"`
}

type User struct {
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
	AvatarUrl   string `json:"avatarUrl"`
	Type        string `json:"type"`
}

type Card struct {
	CardsV2 []CardHeader `json:"cardsV2"`
}

type CardHeader struct {
	CardId string     `json:"cardId"`
	Card   CardDetail `json:"card"`
}

type CardDetail struct {
	Header   Header    `json:"header"`
	Sections []Section `json:"sections"`
}

type Header struct {
	Title        string `json:"title"`
	Subtitle     string `json:"subtitle"`
	ImageUrl     string `json:"imageUrl"`
	ImageType    string `json:"imageType"`
	ImageAltText string `json:"imageAltText"`
}

type Section struct {
	Header                    string          `json:"header"`
	Collapsible               bool            `json:"collapsible"`
	UncollapsibleWidgetsCount int             `json:"uncollapsibleWidgetsCount"`
	Widgets                   []MessageWidget `json:"widgets"`
}

type MessageWidget struct {
	TextParagraph *Message    `json:"textParagraph,omitempty"`
	ButtonList    *ButtonList `json:"buttonList,omitempty"`
}

type Message struct {
	Text string `json:"text"`
}

type ButtonList struct {
	Buttons []Button `json:"buttons"`
}

type Button struct {
	Text    string  `json:"text"`
	OnClick OnClick `json:"onClick"`
}

type OnClick struct {
	OpenLink OpenLink `json:"openLink"`
}

type OpenLink struct {
	URL string `json:"url"`
}

func BuildButtons(buttons []Button) *ButtonList {
	var validButtons []Button

	for _, b := range buttons {
		if b.Text != "" && b.OnClick.OpenLink.URL != "" {
			validButtons = append(validButtons, b)
		}
	}

	// return nil if no valid buttons → JSON won't include buttonList
	if len(validButtons) == 0 {
		return nil
	}

	return &ButtonList{Buttons: validButtons}
}

func BuildMessage(text string) *Message {

	if text != "" {
		message := Message{
			Text: text,
		}
		return &message
	}
	return nil
}
