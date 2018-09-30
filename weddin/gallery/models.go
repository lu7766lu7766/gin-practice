package gallery

type GalleryModel struct {
	ID       uint   `gorm:"primary_key"`
	Filename string `gorm:"column:filename"`
	Uploder  string `gorm:"column:uploader"`

	// Time
	// Image        *string `gorm:"column:image"`

}

// Migrate the schema of database if needed
// func AutoMigrate() {
// 	db := common.GetDB()

// 	db.AutoMigrate(&UserModel{})
// 	db.AutoMigrate(&FollowModel{})
// }
