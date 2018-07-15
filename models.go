package models

import "time"

//******************************************************************************

// Media type.
type Media struct {
	Title string
	Type  string // gif, png, mp4, ...
	Path  string
}

// Address type.
type Address struct {
	Value     string
	IsDefault bool
}

//******************************************************************************

// SMS model.
type SMS struct {
	ID          int64 `json:"-"`
	Code        int
	UUID        string
	PhoneNumber string
	RequestDate time.Time
	ExpiryDate  time.Time
}

//******************************************************************************

// AppInfo is used to keep the latest version of different components of the
// system in database such as app version or backend version.AppInfo.
type AppInfo struct {
	ID          int64 `json:"-"`
	Name        string
	Version     string
	IsSupported bool
}

//******************************************************************************

// UserType enum type.
type UserType string

const (
	// UserTypeBusinessOwner enum value.
	UserTypeBusinessOwner UserType = "business-owner"
	// UserTypeAdmin enum value.
	UserTypeAdmin UserType = "admin"
	// UserTypeShopOwner enum value.
	UserTypeShopOwner UserType = "shop-owner"
	// UserTypeAppUser enum value.
	UserTypeAppUser UserType = "app-user"
)

// User model.
type User struct {
	ID          int `json:"-"`
	Type        UserType
	PhoneNumber string
	RealName    string
	Email       string
	Address     []Address
	IsComplete  bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

//******************************************************************************

// MetaProduct model.
type MetaProduct struct {
	ID        int `json:"-"`
	ProductID int
	Revision  int
}

// Product model.
type Product struct {
	ID           int `json:"-"`
	Name         string
	Code         int // TODO: ???
	Price        float32
	Brand        string
	Model        string
	Thumbnail    Media
	MediaGallery []Media
	Attributes   map[string]interface{}
	Description  string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
}

//******************************************************************************

// MetaShop model.
type MetaShop struct {
	ID       int `json:"-"`
	ShopID   int
	Revision int
}

// Shop model.
type Shop struct {
	ID           int `json:"-"`
	Name         string
	Brand        string
	Owner        string
	Keeper       string
	Address      string  // TODO: ???
	Thumbnail    Media   // TODO: ???
	MediaGallery []Media // TODO: ???
	Description  string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
}

//******************************************************************************

// Contracts model
type Contracts struct {
	ID        int `json:"-"`
	ShopID    int
	Number    int
	Date      time.Time
	Duration  time.Duration
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}