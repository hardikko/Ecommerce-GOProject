package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User
type User struct {
	ID            primitive.ObjectID `json:"_id" bson:"_id"`
	FirstName     *string            `json:"firstname"		validate:"required,min=2,max=30"`
	LastName      *string            `json:"lastname"		validate:"required,min=2,max=30"`
	Email         *string            `json:"email"			validate:"email,required"`
	Password      *string            `json:"password"		validate:"required,min=6`
	Phone         *string            `json:"phone"			validate:"required"`
	Token         *string            `json:"token"`
	Refresh_Token *string            `json:"refresh_token"`
	Created_At    time.Time          `json:"created_at"`
	Updated_At    time.Time          `json:"updated_at"`
	User_ID       string             `json:"user_id"`
	UserCart      []ProductUser      `json:"usercart" bson:"usercart"`
	Address       []Address          `json:"address" bson:"address"`
	Order_Status  []Order            `json:"order" bson:"order"`
}

// Product
type Product struct {
	Product_ID   primitive.ObjectID `bson:"_id"`
	Product_Name *string            `json:"product_name"`
	Price        *uint64            `json:"price"`
	Rating       *uint8             `json:"rating"`
	Image        *string            `json:"image"`
}

// ProductUser
type ProductUser struct {
	Product_ID   primitive.ObjectID `bson:"_id"`
	Product_Name *string            `json:"product_name" bson:"product_name"`
	Price        int                `json:"price" bson:"price"`
	Rating       *uint              `json:"rating" bson:"rating"`
	Image        *string            `json:"image" bson:"image"`
}

// Address
type Address struct {
	Address_ID primitive.ObjectID `bson:"_id"`
	House      *string            `json:"house" bson:"house"`
	Street     *string            `json:"street" bson:"street"`
	City       *string            `json:"city" bson:"city"`
	Pincode    *string            `json:"pincode" bson:"pincode"`
}

// Order
type Order struct {
	Order_ID       primitive.ObjectID `bson:"_id"`
	Order_Cart     []ProductUser      `json:"order_list" bson:"order_list"`
	Ordered_At     time.Time          `json:"ordered_at" bson:"ordered_at"`
	Price          int                `json:"total_price" bson:"total_price"`
	Discount       *int               `json:"discount" bson:"discount"`
	Payment_Method Payment            `json:"payment" bson:"payment"`
}

// Payment
type Payment struct {
	Digital bool
	COD     bool
}
