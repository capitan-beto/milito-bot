package models

import "github.com/shopspring/decimal"

type WhatsappBody struct {
	MessagingProduct string       `json:"messaging_product"`
	RecipientType    string       `json:"recipient_type,omitempty"`
	Status           string       `json:"status,omitempty"`
	To               string       `json:"to,omitempty"`
	Type             string       `json:"type,omitempty"`
	Template         *Template    `json:"template,omitempty"`
	Text             *Text        `json:"text,omitempty"`
	Context          *Context     `json:"context,omitempty"`
	MessageID        string       `json:"message_id,omitempty"`
	Interactive      *Interactive `json:"interactive,omitempty"`
}

type Template struct {
	Name     string   `json:"name"`
	Language Language `json:"language"`
}

type Language struct {
	Code string `json:"code"`
}

type Text struct {
	Body string `json:"body,omitempty"`
}

type Context struct {
	MessageId string `json:"message_id"`
}

//lists types

type Interactive struct {
	Type   string `json:"type"`
	Header Header `json:"header"`
	Body   Body   `json:"body"`
	Footer Footer `json:"footer,omitempty"`
	Action Action `json:"action"`
}

type Header struct {
	Type string `json:"type"`
	Text string `json:"text,omitempty"`
}

type Body struct {
	Text string `json:"text,omitempty"`
}

type Footer struct {
	Text string `json:"text"`
}

type Action struct {
	Button   string    `json:"button"`
	Sections []Section `json:"sections"`
}

type Section struct {
	Title string `json:"title"`
	Rows  []Row  `json:"rows"`
}

type Row struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

//Post Webhhok types

type PostWebhook struct {
	Object string         `json:"object"`
	Entry  []WebhookEntry `json:"entry"`
}

type WebhookEntry struct {
	Id      string           `json:"id"`
	Changes []WebhookChanges `json:"changes"`
}

type WebhookChanges struct {
	Value WebhookValue `json:"value"`
	Field string       `json:"field"`
}

type WebhookValue struct {
	MessagingProduct string            `json:"messaging_product"`
	Metadata         WebhookMetadata   `json:"metadata"`
	Contacts         []WebhookContacts `json:"contacts"`
	Messages         []WebhookMessages `json:"messages"`
}

type WebhookMetadata struct {
	DisplayPhoneNumber string `json:"display_phone_number"`
	PhoneNumberId      string `json:"phone_number_id"`
}

type WebhookContacts struct {
	Profile WebhookProfile `json:"profile"`
	WaId    string         `json:"wa_id"`
}

type WebhookProfile struct {
	Name string `json:"name"`
}

type WebhookMessages struct {
	From        string             `json:"from"`
	Id          string             `json:"id"`
	Timestamp   string             `json:"timestamp"`
	Text        WebhookText        `json:"text,omitempty"`
	Type        string             `json:"type"`
	Context     WebhookContext     `json:"context,omitempty"`
	Interactive WebhookInteractive `json:"interactive,omitempty"`
}

type WebhookText struct {
	Body string `json:"body"`
}

type WebhookContext struct {
	From string `json:"from"`
	Id   string `json:"id"`
}

type WebhookInteractive struct {
	Type      string    `json:"type,omitempty"`
	ListReply ListReply `json:"list_reply"`
}

type ListReply struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

//Products list and MAC Services models DELETE THIS SHIT LATER!!!

type Product struct {
	ID    string          `json:"id"`
	Desc  string          `json:"item_desc"`
	Price decimal.Decimal `json:"price"`
	Date  string          `json:"date"`
}
