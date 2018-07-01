package helpers

type Message struct {
    ID      uint   `json:"id"`
    Title   string `json:"title"`
    Message string `json:"message"`
    Type string `json:"type"`
}