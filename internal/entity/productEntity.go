package entity

type Product struct {
	Id       string   `bson:"_id" json:"_id"`
	Title    string   `bson:"title" json:"title"`
	Year     int      `bson:"year" json:"year"`
	Director string   `bson:"director" json:"string"`
	Cast     []string `bson:"cast" json:"cast"`
}
