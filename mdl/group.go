package mdl

// Group user group,用户组
type Group struct {
	tableName   struct{} `sql:"group"`
	ID          string   `json:"id" validate:"required,max=10" sql:",pk"` //used for FK
	Name        string   `json:"name" validate:"required,max=10"`
	Permissions []string `json:"permissions" validate:"required,min=1,dive,min=3,max=10,required" sql:",array"` // Permissions the group own, lang:zh_CN 组具备的权限
	Desc        string   `json:"desc"`                                                                          //description
	Users       []User   `json:"user" pg:"many2many:user_group"`
}
