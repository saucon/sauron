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
	TextParagraph Message `json:"textParagraph"`
}

type Message struct {
	Text string `json:"text"`
}
