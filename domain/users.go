package domain

import (
	"time"

	"gorm.io/gorm"
)

type UserUsecase interface {
	CreateUser(User) (UserResponse, error)
	GetUserById(userID string) (User, error)
}

type UserRepository interface {
	CreateUser(User) (User, error)
	GetUserById(userID string) (User, error)
}

type User struct {
	UserID      string `gorm:"primary_key;AUTO_INCREMENT" json:"user_id"`
	Name        string `gorm:"NOT NULL;TYPE:varchar(100)" json:"name"`
	Address     string `gorm:"TYPE:varchar(100)" json:"address"`
	Numberphone string `gorm:"TYPE:varchar(10)" json:"numberphone"`
	Email       string `gorm:"unique;NOT NULL;TYPE:varchar(100)" json:"email"`
	Status      int
	CreatedAt   time.Time  `gorm:"DEFAULT:now()" json:"created_at"`
	UpdatedAt   *time.Time `gorm:"DEFAULT:now()" json:"updated_at"`
	DeletedAt   *time.Time `gorm:"DEFAULT:NULL" json:"deleted_at"`
}

// UserWithoutPassword ..
type UserResponse struct {
	ID       uint32 `json:"user_id" form:"user_id"`
	Name     string `json:"name" form:"name" valid:"required"`
	Username string `json:"username" form:"username" `
	Tel      string `json:"tel" form:"tel"`
	// Type      UserType       `json:"type" form:"type" valid:"required"`
	CreatedAt *time.Time     `json:"created_at" gorm:"default:now();"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	UpdatedAt *time.Time     `json:"updated_at" gorm:"default:now();"`
}

// type Shop struct {
// 	ShopID      uint32     `gorm:"primary_key;AUTO_INCREMENT" json:"shop_id"`
// 	ShopName    string     `gorm:"TYPE:varchar(100)" json:"shopname"`
// 	Address     string     `gorm:"TYPE:varchar(100)" json:"address"`
// 	NumberPhone string     `gorm:"TYPE:varchar(10)" json:"numberphone"`
// 	CreatedAt   time.Time  `gorm:"DEFAULT:now()" json:"created_at"`
// 	UpdatedAt   *time.Time `gorm:"DEFAULT:now()" json:"updated_at"`
// 	DeletedAt   *time.Time `gorm:"DEFAULT:NULL" json:"deleted_at"`
// }

// type Owner struct {
// 	OwnerID   uint       `gorm:"primary_key;AUTO_INCREMENT" json:"owner_id"`
// 	NameOwner string     `gorm:"TYPE:varchar(100)" json:"name_owner"`
// 	CreatedAt time.Time  `gorm:"DEFAULT:now()" json:"created_at"`
// 	UpdatedAt *time.Time `gorm:"DEFAULT:now()" json:"updated_at"`
// 	DeletedAt *time.Time `gorm:"DEFAULT:NULL" json:"deleted_at"`
// }

// type Baber struct {
// 	BaberID     uint32     `gorm:"primary_key;AUTO_INCREMENT" json:"baber_id"`
// 	Namebaber   string     `gorm:"NOT NULL;TYPE:varchar(100)" json:"name_baber"`
// 	PhoneNumber string     `gorm:"NOT NULL;TYPE:varchar(10)" json:"phone_number"`
// 	WorkStatus  bool       `json:"work_status"`
// 	CreatedAt   time.Time  `gorm:"DEFAULT:now()" json:"created_at"`
// 	UpdatedAt   *time.Time `gorm:"DEFAULT:now()" json:"updated_at"`
// 	DeletedAt   *time.Time `gorm:"DEFAULT:NULL" json:"deleted_at"`
// }

// type Booking struct {
// 	BookingID       uint32     `gorm:"primary_key;AUTO_INCREMENT" json:"booking_id"`
// 	UserID          string     `gorm:"NOT NULL;TYPE:varchar(100)" json:"user_id"`
// 	ShopID          string     `gorm:"NOT NULL;TYPE:varchar(100)" json:"shop_id"`
// 	PhoneNumberUser string     `gorm:"NOT NULL;TYPE:varchar(10)" json:"phone_number_user"`
// 	UserName        string     `gorm:"NOT NULL;TYPE:varchar(100)" json:"user_name"`
// 	ShopName        string     `gorm:"NOT NULL;TYPE:varchar(100)" json:"shopname"`
// 	TimeBooking     string     `gorm:"NOT NULL;TYPE:varchar(1)" json:"time_booking"`
// 	StatusCut       bool       `gorm:"DEFAULT:false" json:"status_cut"`
// 	CreatedAt       time.Time  `gorm:"DEFAULT:now()" json:"created_at"`
// 	UpdatedAt       *time.Time `gorm:"DEFAULT:now()" json:"updated_at"`
// 	DeletedAt       *time.Time `gorm:"DEFAULT:NULL" json:"deleted_at"`
// }
