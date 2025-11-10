package users

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID                  primitive.ObjectID   `bson:"_id,omitemtpy" json:"id"`
	Name                string               `bson:"name" json:"name"`
	Email               string               `bson:"email" json:"email"`
	Password            string               `bson:"password" json:"-"`
	Phone               string               `bson:"phone,omitempty" json:"phone"`
	Role                string               `bson:"role" json:"role"`
	IsVerified          bool                 `bson:"isVerified" json:"isVerified"`
	Addresses           []primitive.ObjectID `bson:"addresses,omitempty" json:"addresses"`
	ProfileImage        string               `bson:"profileImage,omitempty" json:"profileImage"`
	ResetPasswordToken  string               `bson:"resetPasswordToken,omitempty" json:"resetPasswordToken"`
	ResetPasswordExpiry time.Time            `bson:"resetPasswordExpiry,omitempty" json:"resetPasswordExpiry"`
	LastLoginAt         *time.Time           `bson:"lastLoginAt, omitempty" json:"lastLoginAt"`
	IsBlocked           bool                 `bson:"isBlocked,omitempty" json:"isBlocked"`
	CreatedAt           time.Time            `bson:"createdAt,omitempty" json:"createdAt"`
	UpdatedAt           time.Time            `bson:"updatedAt,omitempty" json:"updatedAt"`
	DeletedAt           time.Time            `bson:"deletedAt,omitempty" json:"deletedAt"`
}

type Address struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID    primitive.ObjectID `bson:"userId,omitempty" json:"userId"`
	Label     string             `bson:"label,omitempty" json:"label"`
	Name      string             `bson:"name,omitempty" json:"name"`
	Phone     string             `bson:"phone,omitempty" json:"phone"`
	Line1     string             `bson:"line1,omitempty" json:"line1"`
	City      string             `bson:"city,omitempty" json:"city"`
	State     string             `bson:"state,omitempty" json:"state"`
	Pincode   string             `bson:"pincode,omitempty" json:"pincode"`
	Country   string             `bson:"country,omitempty" json:"country"`
	IsDefault bool               `bson:"isDefault,omitempty" json:"isDefault"`
	CreatedAt time.Time          `bson:"createdAt,omitempty" json:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt,omitempty" json:"updatedAt"`
}
