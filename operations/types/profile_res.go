package types

// QuickCreateResponse create Profile response
type QuickCreateResponse = CommonResponse[[]string]

// StartProfileRes
type StartProfileRes struct {
	EnvId     string `json:"envId"`     // Profile ID
	DebugPort string `json:"debugPort"` // debugPort
	Webdriver string `json:"webdriver"` // Webdriver Path
}

// ProfileEnvIdRes
type ProfileEnvIdRes struct {
	EnvId string `json:"envId"` // Profile ID
}

// StartProfileResponse start Profile response
type StartProfileResponse = CommonResponse[StartProfileRes]

// ProfilePageRes
type ProfilePageRes struct {
	Current  string            `json:"current"`  // current page number
	Pages    string            `json:"pages"`    // page number
	Total    string            `json:"total"`    // total number
	DataList []PageProfileInfo `json:"dataList"` // Webdriver Path
}

// PageProfileInfo
type PageProfileInfo struct {
	ID      string `json:"id"`      // Profile ID
	EnvName string `json:"envName"` // Profile name
	GroupID string `json:"groupId"` // Group ID
	ProxyID string `json:"proxyId"` // Proxy ID
}

type ProfileSecurityLockStatus struct {
	EnvId  string `json:"envId"`  // Profile ID
	Locked bool   `json:"locked"` // Profile name
}

type ProfileDetailInfo struct {
	AccountInfo        AccountInfo        `json:"accountInfo"`        // Account information
	AdvancedSetting    AdvancedSetting    `json:"advancedSetting"`    // Advanced setting (empty object)
	AfterStartupConfig AfterStartupConfig `json:"afterStartupConfig"` // Configuration after startup
	BrowserCore        int                `json:"browserCore"`        // Kernel version number
	BrowserTypeId      int                `json:"browserTypeId"`      // Browser type, 1: Chrome, 2: Firefox
	Cookies            string             `json:"cookies"`            // Cookie
	EnvName            string             `json:"envName"`            // Profile name
	EnvRemark          string             `json:"envRemark"`          // Profile remark
	GroupId            string             `json:"groupId"`            // Group ID
	Id                 string             `json:"id"`                 // Profile ID
	IsEncrypt          int                `json:"isEncrypt"`          // Whether “end-to-end encryption”, 0: No, 1: Yes
	OperatorSystemId   int                `json:"operatorSystemId"`   // Operating system type, 1: Windows, 2: macOS, 3: Android, 4: IOS
	ProxyId            string             `json:"proxyId"`            // Proxy ID
	TagIds             []string           `json:"tagIds"`             // Tag ID list
	UaVersion          int                `json:"uaVersion"`          // UA version
}

type AfterStartupConfig struct {
	AfterStartup int      `json:"afterStartup"` // Startup settings: 1~4
	AutoOpenUrls []string `json:"autoOpenUrls"` // Open URLs
	PlatformUrl  string   `json:"platformUrl"`  // Platform address
}

type BrowserKernel struct {
	BrowserType int   `json:"browserType"` //  Browser type, 1：Chrome，2： Firefox
	Versions    []int `json:"versions"`    // Versions
}

type UA struct {
	Ua         string     `json:"Ua"`         // user agent
	Resolution Resolution `json:"resolution"` // user agent
}

type Resolution struct {
	Id    string `json:"id"`    //  Resolution ID
	Value string `json:"value"` // Resolution value
}

// Platform
type Platform struct {
	CategoryID string `json:"categoryId"` // Category ID
	GroupName  string `json:"groupName"`  // Group name
	Groups     int    `json:"groups"`     // Group, e.g. 0: Amazon
	ID         int64  `json:"id"`         // Platform ID
	IsCustomer bool   `json:"isCustomer"` // Whether to customize the platform
	Logo       string `json:"logo"`       // Platform logo
	Name       string `json:"name"`       // Platform name
	OrderNo    int    `json:"orderNo"`    // Sort Number
	Sites      []Site `json:"sites"`      // Sites information
}

// Site
type Site struct {
	Country   string `json:"country"`   // Country
	Host      string `json:"host"`      // Site Domain Name
	ID        string `json:"id"`        // Site ID
	IsDefault bool   `json:"isDefault"` // Is default site
	Logo      string `json:"logo"`      // Logo
	Name      string `json:"name"`      // Size name
	NameBak   string `json:"nameBak"`   // Size name backup
	URL       string `json:"url"`       // Size address
}

type TimeZoneAndLanguage struct {
	LanguageList []Language `json:"language_list"`  // Language
	TimeZoneList []TimeZone `json:"time_zone_list"` // TimeZone
}
type TimeZone struct {
	Id    string `json:"id"`              // TimeZone Id
	Value string `json:"value,omitempty"` // TimeZone value
}

type Language struct {
	Id    string `json:"id"`              // Language Id
	Value string `json:"value,omitempty"` // Language value
}

// ProfileRunningStatusRes
type ProfileRunningStatusRes struct {
	EnvId       string `json:"envId"`       // Profile ID
	Status      string `json:"status"`      // Profile ID
	LocalStatus string `json:"localStatus"` // Profile ID
	DebugPort   string `json:"debugPort"`   // debugPort
	Webdriver   string `json:"webdriver"`   // Webdriver Path
}
