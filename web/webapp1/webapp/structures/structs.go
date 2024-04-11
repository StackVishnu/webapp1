package structures

type User struct {
	ID       int    `json:"uid"`
	Username string `json:"user"`
	Password string `json:"pass"`
	Role     string `json:"role"`
}

type Mobile struct {
	ID    int    `json:"mid"`
	Name  string `json:"mName"`
	Price int    `json:"mPrice"`
	Specs string `json:"mSpecs"`
	Ipath string `json:"path"`
}
