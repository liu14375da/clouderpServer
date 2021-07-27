package tbsfUser

type TbsfUser struct {
	UserId       string `db:"user_id"`
	UserName     string `db:"user_name"`
	StaffId      string `db:"staff_id"`
	UserPassword string `db:"user_password"`
	UserType     string `db:"user_type"`
	IsEnable     bool   `db:"IsEnable"`
	IsAdmin      bool   `db:"IsAdmin"`
	CompanyId    string `db:"company_id"`
	//LastModificationTime time.Time `gorm:"type:datetime"`
}

type UserId struct {
	User_ID       string `db:"User_ID"`
	User_Name     string `db:"User_Name"`
	Staff_ID      string `db:"Staff_ID"`
	User_Password string `db:"User_Password"`
	Company_ID    string `db:"Company_ID"`
}

