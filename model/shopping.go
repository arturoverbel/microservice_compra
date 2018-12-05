import "gopkg.in/mgo.v2/bson"

type Shopping struct {
	ID         bson.ObjectId `bson:"_id" json:"id"`
	User       int           `bson:"user" json:"user"`
	Products   []int         `bson:"products" json:"products"`
	Payment    string        `bson:"payment" json:"payment"`
	PriceTotal int           `bson:"price_total" json:"price_total"`
}
