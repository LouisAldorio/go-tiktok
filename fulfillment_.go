package gotiktok

type Tracking struct {
	Description      string `json:"description"`
	UpdateTimeMillis int64  `json:"update_time_millis"`
}

type TrackingList struct {
	Trackings []Tracking `json:"tracking"`
}

type PackageSku struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
	Quantity int    `json:"quantity"`
}

type PackageOrder struct {
	Id   string       `json:"id"`
	Skus []PackageSku `json:"skus"`
}

type PickupSlot struct {
	StartTime int `json:"start_time"`
	EndTime   int `json:"end_time"`
}

type PackageAddress struct {
	FullAddress   string `json:"full_address"`
	PhoneNumber   string `json:"phone_number"`
	Name          string `json:"name"`
	PostalCode    string `json:"postal_code"`
	AddressDetail string `json:"address_detail"`
	RegionCode    string `json:"region_code"`
	AddressLine1  string `json:"address_line1"`
	AddressLine2  string `json:"address_line2"`
	AddressLine3  string `json:"address_line3"`
	AddressLine4  string `json:"address_line4"`
}

type PackageDimension struct {
	Length string `json:"length"`
	Width  string `json:"width"`
	Height string `json:"height"`
	Unit   string `json:"unit"`
}

type PackageInsurance struct {
	IsPurchased    bool   `json:"is_purchased"`
	CoverageAmount string `json:"coverage_amount"`
	IsClaimEligible bool  `json:"is_claim_eligible"`
	ClaimStatus    string `json:"claim_status"`
}

type PackageData struct {
	PackageId             string           `json:"package_id"`
	Orders                []PackageOrder   `json:"orders"`
	PackageStatus         string           `json:"package_status"`
	PackageSubStatus      string           `json:"package_sub_status"`
	ShipExceptionReason   string           `json:"ship_exception_reason"`
	SplitAndCombineTag    string           `json:"split_and_combine_tag"`
	HasMultiSkus          bool             `json:"has_multi_skus"`
	NoteTag               string           `json:"note_tag"`
	ShippingProviderName  string           `json:"shipping_provider_name"`
	ShippingProviderId    string           `json:"shipping_provider_id"`
	ShippingType          string           `json:"shipping_type"`
	DeliveryOptionName    string           `json:"delivery_option_name"`
	DeliveryOptionId      string           `json:"delivery_option_id"`
	TrackingNumber        string           `json:"tracking_number"`
	LastMileTrackingNumber string          `json:"last_mile_tracking_number"`
	PickupSlot            PickupSlot       `json:"pickup_slot"`
	CreateTime            int              `json:"create_time"`
	HandoverMethod        string           `json:"handover_method"`
	OrderLineItemIds      []string         `json:"order_line_item_ids"`
	RecipientAddress      PackageAddress   `json:"recipient_address"`
	SenderAddress         PackageAddress   `json:"sender_address"`
	Weight                PackageWeight    `json:"weight"`
	Dimension             PackageDimension `json:"dimension"`
	UpdateTime            int              `json:"update_time"`
	Insurance             PackageInsurance `json:"insurance"`
}
