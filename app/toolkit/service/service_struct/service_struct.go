package service_struct

type SignUpByPhone_PreviewForm struct {
	PhoneNumber string `json:"phone_number" binding:"required,number,len=11"`
}
type SignUpByPhone_Form struct {
	PhoneNumber string `json:"phone_number" binding:"required,number,len=11"`
	MessageCode string `json:"message_code" binding:"required,len=6"`
}
type SignUpByMail_PreviewForm struct {
	Mail string `form:"mail" json:"mail" binding:"required,email"`
}
type SignUpByMail_Form struct {
	Mail        string `form:"mail" json:"mail"  binding:"required,email"`
	Username    string `json:"username" form:"username"  binding:"required,gte=5,lte=10"`
	Password    string `json:"password" form:"password" binding:"required,gte=5,lte=18"`
	MessageCode string `json:"message_code" form:"message_code" binding:"required,len=6"`
}
type LoginByPass_Form struct {
	Username string `json:"username" form:"username" binding:"required,gte=5,lte=10"`
	Password string `json:"password" form:"password" binding:"required,gte=5,lte=18"`
}
type ChangePass_Form struct {
	Username string `json:"username" form:"username" binding:"required,gte=5,lte=10"`
	Password string `json:"password" form:"password" binding:"required,gte=5,lte=18"`
	New_Password string `json:"new_password" form:"new_password" binding:"required,gte=5,lte=18,nefield=Password"`
}

type SQL_Table_user struct {
	Id       int `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email string `json:"email"`
	Unique_salt string `json:"unique_salt"`
	Sign_time string `json:"sign_time"`
}
