package metrica

type EmojiStatus struct {
	DocumentID string `json:"document_id"`
}

type Profile struct {
	ID          int64        `json:"id"`
	FirstName   string       `json:"first_name"`
	LastName    string       `json:"last_name"`
	Title       string       `json:"title"`
	Username    string       `json:"username"`
	Phone       string       `json:"phone"`
	Bio         string       `json:"bio"`
	EmojiStatus *EmojiStatus `json:",inline"`
}

type StatusInfo struct {
	Status string `json:"status"`
	Date   string `json:"date"`
}

type PhotoInfo struct {
	CurrentPhotoURL string `json:"current_photo_url"`
	TotalPhotos     int    `json:"total_photos"`
}

type WeeklyOpen struct {
	StartMinute int `json:"start_minute"`
	EndMinute   int `json:"end_minute"`
}

type WorkHours struct {
	Timezone   string       `json:"timezone"`
	WeeklyOpen []WeeklyOpen `json:",inline"`
}

type GeoPoint struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"long"`
}

type Location struct {
	Address  string    `json:"address"`
	GeoPoint *GeoPoint `json:",inline"`
}

type BusinessInfo struct {
	WorkHours WorkHours `json:",inline"`
	Location  *Location `json:",inline"`
}

type CollectibleInfo struct {
	Username   string `json:"username"`
	IsActive   bool   `json:"is_active"`
	IsEditable bool   `json:"is_editable"`
}

type PremiumInfo struct {
	IsPremium bool `json:"is_premium"`
}

type UserFlags struct {
	IsVerified bool `json:"is_verified"`
	IsScam     bool `json:"is_scam"`
	IsSupport  bool `json:"is_support"`
	IsBot      bool `json:"is_bot"`
	IsSelf     bool `json:"is_self"`
}

type User struct {
	Profile          *Profile           `json:",inline"`
	StatusInfo       *StatusInfo        `json:",inline"`
	PhotoInfo        *PhotoInfo         `json:",inline"`
	BusinessInfo     *BusinessInfo      `json:",inline"`
	CollectiblesInfo *[]CollectibleInfo `json:",inline"`
	PremiumInfo      *PremiumInfo       `json:",inline"`
	UserFlags        *UserFlags         `json:",inline"`
}
