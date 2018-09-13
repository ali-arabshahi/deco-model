package models

import "time"

//******************************************************************************

// Media type.
type Media struct {
	Title string
	Type  string // gif, png, mp4, ...
	Path  string
}

// Model3D type.
type Model3D struct {
	Title        string
	ObjectPath   string
	TexturesPath []struct {
		Path      string
		Thumbnail string
	}
}

// Address type.
type Address struct {
	Country     string
	City        string
	Value       string
	PhoneNumber string
	PostCode    string
	IsDefault   bool
}

//BankAccount type
type BankAccount struct {
	Bank          string
	Ownername     string
	AccountNumber string
}

//Discount model //TODO: ????
type Discount struct {
	ID         int `sql:",AUTO_INCREMENT"`
	Percentage int
	ValidTime  time.Time
}

//******************************************************************************

//SMS model.
type SMS struct {
	ID          int64 `json:"-" sql:",AUTO_INCREMENT"`
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
	ID          int64 `json:"-" sql:",AUTO_INCREMENT"`
	Name        string
	Version     string
	IsSupported bool
}

//******************************************************************************

// UserType enum type.TODO: sepehr will check if is this needed or not
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

//AppUser model
type AppUser struct {
	ID                int `json:"-" sql:",AUTO_INCREMENT"`
	PhoneNumber       string
	FirstName         string
	LastName          string
	Gender            string
	Birthday          string
	Email             string
	Address           []Address
	FavouriteProducts []int `pg:",array"`
	IsComplete        bool
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         time.Time
}

//ShopUser model
type ShopUser struct {
	ID          int `json:"-" sql:",AUTO_INCREMENT"`
	PhoneNumber string
	FirstName   string
	LastName    string
	Gender      string
	Birthday    string
	Email       string
	Address     []Address
	IsComplete  bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

//SuperUser model
type SuperUser struct {
	ID          int `json:"-" sql:",AUTO_INCREMENT"`
	PhoneNumber string
	FirstName   string
	LastName    string
	Gender      string
	Birthday    string
	Email       string
	Address     []Address
	IsComplete  bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

//DebitcardUser model.
type DebitcardUser struct {
	ID          int `sql:",AUTO_INCREMENT"`
	Number      int
	ShopID      int
	UserID      int
	Credit      int
	Active      bool
	AssighnDate time.Time
}

// AdminPanelUser model.
type AdminPanelUser struct {
	ID       int `sql:",AUTO_INCREMENT"`
	Username string
	Password string
	RealName string
}

//Mall model
type Mall struct {
	ID         int `sql:",AUTO_INCREMENT"`
	Name       string
	FloorCount int
	UnitCount  int
}

// //ShopsDetails model--dont need it
// type ShopsDetails struct {
// 	ID              int
// 	MallID          int
// 	AppAddress      int
// 	ReadableAddress string
// }

// MetaProduct model.
type MetaProduct struct {
	ID        int `json:"-" sql:",AUTO_INCREMENT"`
	ProductID int
	Revision  int
}

// Product model.
type Product struct {
	ID                int `sql:",AUTO_INCREMENT"`
	ShopID            int
	Name              string
	Brand             string
	Code              int
	Price             int
	DeliveryTime      int
	ManufactureSerial string
	CategoryID        int
	Discount          int
	Thumbnail         string
	Gallery           []Media
	Gallery3D         Model3D
	DeliveryPrice     int
	Guarantee         string
	Description       string
	Attributes        map[string]interface{}
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         time.Time
}

//******************************************************************************

// MetaShop model.
type MetaShop struct {
	ID       int `json:"-" sql:",AUTO_INCREMENT"`
	ShopID   int
	Revision int
}

//Shop model
type Shop struct {
	ID           int `json:"-" sql:",AUTO_INCREMENT"`
	MallID       int
	ContractID   int
	ShopName     string
	Brand        string
	Owner        string
	Keeper       string
	VirAddress   int // TODO: format 311 floor 3 number 11
	Adress       Address
	PhoneNumbers []string `pg:",array"`
	Thumbnail    Media    // TODO: create from on file chossen by user
	MediaGallery []Media  // TODO: hard disk format (by name or by directory hierarchical)
	Description  string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
}

//Category model
type Category struct {
	ID        int `sql:",AUTO_INCREMENT"`
	Name      string
	Thumbnail Media
}

//******************************************************************************

//FIXME: add 4 field for %
//Contracts model (contracts are immutible)
type Contracts struct {
	ID             int `json:"-" sql:",AUTO_INCREMENT"`
	Number         int
	ContractScan   []Media
	ContracrtType  string //pysical or virtual
	CompaneyName   string
	Date           time.Time
	BankAccount    []BankAccount
	Duration       time.Duration
	MontlytDueDate time.Time
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      time.Time
}
