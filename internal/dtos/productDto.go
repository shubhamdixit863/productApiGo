package dtos

type ProductRequest struct {
	Title    string   `json:"title"`
	Year     int      ` json:"year"`
	Director string   ` json:"director"`
	Cast     []string `bson:"cast" json:"cast"`
}
