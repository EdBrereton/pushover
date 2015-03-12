package pushover

// Response is a container used to hold the response from the pushover service
type Response struct {
	Status  int      `json:"status"`
	Request string   `json:"request"`
	Error   []string `json:"errors"`

	Token     string `json:"token"`
	User      string `json:"user"`
	Message   string `json:"message"`
	Device    string `json:"device"`
	Title     string `json:"title"`
	Url       string `json:"url"`
	Url_title string `json:"url_title"`
	Priority  string `json:"priority"`
	Timestamp string `json:"timestamp"`
	Sound     string `json:"sound"`
}
