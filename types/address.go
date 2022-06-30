package types

type Address struct {
	Country  string `bson:"country" json:"country"`
	State    string `bson:"state" json:"state"`
	City     string `bson:"city" json:"city"`
	District string `bson:"district" json:"district"`
	ZipCode  string `bson:"zipcode" json:"zipcode"`
	Line1    string `bson:"line1" json:"line1"`
	Line2    string `bson:"line2" json:"line2"`
}
