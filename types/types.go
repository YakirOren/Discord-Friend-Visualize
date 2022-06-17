package types

type Friend struct {
	Id   string `json:"id"`
	User struct {
		Username string `json:"username"`
	} `json:"user"`
}

type Self struct {
	Id               string      `json:"id"`
	Username         string      `json:"username"`
	Avatar           string      `json:"avatar"`
	AvatarDecoration interface{} `json:"avatar_decoration"`
	Discriminator    string      `json:"discriminator"`
	PublicFlags      int         `json:"public_flags"`
	Flags            int         `json:"flags"`
	PurchasedFlags   int         `json:"purchased_flags"`
	Banner           interface{} `json:"banner"`
	BannerColor      string      `json:"banner_color"`
	AccentColor      int         `json:"accent_color"`
	Bio              string      `json:"bio"`
	Pronouns         string      `json:"pronouns"`
	Locale           string      `json:"locale"`
	NsfwAllowed      bool        `json:"nsfw_allowed"`
	MfaEnabled       bool        `json:"mfa_enabled"`
	Email            string      `json:"email"`
	Verified         bool        `json:"verified"`
	Phone            string      `json:"phone"`
}
