package dto

type LoginResponse struct {
	Session   struct {
			  Name  string  `json:"name"`
			  Value string  `json:"value"`
		  } `json:"session"`

	LoginInfo struct {
			  FailedLoginCount    int16 `json:"failedLoginCount"`
			  LoginCount          int16 `json:"loginCount"`
			  LastFailedLoginTime string `json:"lastFailedLoginTime"`
			  PreviousLoginTime   string `json:"previousLoginTime"`
		  }`json:"loginInfo"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
