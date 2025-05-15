package types

// Start Browser Environment
type StartProfile struct {
	EnvId      int    `json:"envId"`
	UniqueId   int32  `json:"uniqueId"`
	EncryptKey string `json:"encryptKey"`
}

// Stop Browser Environment
type StopProfile struct {
	EnvId    int   `json:"envId"`
	UniqueId int32 `json:"uniqueId"`
}

// Quickly Create Browser Environment
type QuickCreateProfile struct {
	BrowserTypeId    int   `json:"browserTypeId"`
	OperatorSystemId int   `json:"operatorSystemId"`
	Quantity         int   `json:"quantity"`
	BrowserCore      int   `json:"browserCore"`
	GroupId          int32 `json:"groupId"`
	IsEncrypt        int   `json:"isEncrypt"`
}

// AdvancedCreateProfile Advanced creation of browser environment request parameters
type AdvancedCreateProfile struct {
	BrowserTypeId      int32            `json:"browserTypeId"`                // Browser type: 1-Chrome, 2-Firefox
	OperatorSystemId   int32            `json:"operatorSystemId"`             // Operating system type: 1-Windows, 2-macOS, 3-Android, 4-iOS
	EnvName            string           `json:"envName,omitempty"`            // Environment name, limited to 100 characters
	AccountInfo        *AccountInfo     `json:"accountInfo,omitempty"`        // Environment account information
	AdvancedSetting    *AdvancedSetting `json:"advancedSetting,omitempty"`    // Advanced configuration
	AfterStartupConfig *StartupConfig   `json:"afterStartupConfig,omitempty"` // Configuration after environment startup
	BrowserCore        int32            `json:"browserCore,omitempty"`        // Core version number, default: 0-auto match
	Cookies            string           `json:"cookies,omitempty"`            // Cookie
	EnvRemark          string           `json:"envRemark,omitempty"`          // Environment remarks, limited to 1500 characters
	GroupId            int64            `json:"groupId,omitempty"`            // Environment group ID, default: Ungrouped-0
	IsEncrypt          int32            `json:"isEncrypt,omitempty"`          // Whether to enable end-to-end encryption 0-off, 1-on, default 0
	ProxyId            int64            `json:"proxyId,omitempty"`            // Proxy ID, default: 0
	TagIds             []int64          `json:"tagIds,omitempty"`             // Tag IDs, default: None
	UaVersion          int32            `json:"uaVersion,omitempty"`          // UA version, default: 0-all
	StartupParams      []string         `json:"startupParams,omitempty"`      // Startup parameters
	DisableAudio       int32            `json:"disableAudio,omitempty"`       // Disable audio playback: 0-off, 1-on
	DisableVideo       int32            `json:"disableVideo,omitempty"`       // Disable video loading: 0-off, 1-on
	DisableImg         int32            `json:"disableImg,omitempty"`         // Disable image loading: 0-off, 1-on
	ImgLimitSize       int32            `json:"imgLimitSize,omitempty"`       // Image limit size, default 10kb
}

// AccountInfo Account Information
type AccountInfo struct {
	PlatformId  string `json:"platformId"`            // Platform ID, 9999-Custom platform
	CustomerUrl string `json:"customerUrl,omitempty"` // Custom platform URL, required when platformId=9999
	Username    string `json:"username,omitempty"`    // Username, limited to 64 characters
	Password    string `json:"password,omitempty"`    // Password, limited to 50 characters
	OtpSecret   string `json:"otpSecret,omitempty"`   // 2FA Secret Key
	SiteId      string `json:"siteId,omitempty"`      // Site ID
}

// AdvancedSetting Advanced Configuration
type AdvancedSetting struct {
	UA                  string              `json:"ua,omitempty"`                   // Custom environment UA
	TimeZone            *TimeZoneSetting    `json:"time_zone,omitempty"`            // Time zone
	WebRTC              *WebRTCSetting      `json:"web_rtc,omitempty"`              // WebRTC
	GeoLocation         *GeoLocationSetting `json:"geo_location,omitempty"`         // Geolocation
	Language            *LanguageSetting    `json:"language,omitempty"`             // Language
	Resolution          *ResolutionSetting  `json:"resolution,omitempty"`           // Resolution
	Font                *FontSetting        `json:"font,omitempty"`                 // Font
	Canvas              *SwitcherSetting    `json:"canvas,omitempty"`               // Canvas
	WebGLImage          *SwitcherSetting    `json:"webgl_image,omitempty"`          // WebGL Image
	WebGLMetadata       *SwitcherSetting    `json:"webgl_metadata,omitempty"`       // WebGL Metadata
	AudioContext        *SwitcherSetting    `json:"audio_context,omitempty"`        // Audio Context
	MediaDevice         *SwitcherSetting    `json:"media_device,omitempty"`         // Media Device
	ClientRects         *SwitcherSetting    `json:"client_rects,omitempty"`         // Client Rects
	SpeechVoice         *SwitcherSetting    `json:"speech_voise,omitempty"`         // Speech Voices
	HardwareConcurrency int                 `json:"hardware_concurrency,omitempty"` // Hardware concurrency, default: 4
	MemoryDevice        int                 `json:"memery_device,omitempty"`        // Device memory, default: 8
	DoNotTrack          int                 `json:"do_not_track,omitempty"`         // Do Not Track, default: 2
	Bluetooth           *SwitcherSetting    `json:"bluetooth,omitempty"`            // Bluetooth
	Battery             *SwitcherSetting    `json:"battery,omitempty"`              // Battery
	PortScanProtection  *PortScanSetting    `json:"port_scan_protection,omitempty"` // Port Scan Protection
	OSVersion           string              `json:"os_version,omitempty"`           // macOS System Version
	WebGPU              *SwitcherSetting    `json:"web_gpu,omitempty"`              // WebGPU
}

// TimeZoneSetting Time Zone Setting
type TimeZoneSetting struct {
	Switcher int    `json:"switcher"`        // Time zone option, default 1, 1: Match IP, 2: Custom
	Value    string `json:"value,omitempty"` // Required if switcher is 2
}

// WebRTCSetting WebRTC Setting
type WebRTCSetting struct {
	Switcher int `json:"switcher"` // WebRTC option, default 2, 1: Privacy, 2: Replace, 3: Real, 4: Disable, 5: Forward
}

// GeoLocationSetting Geolocation Setting
type GeoLocationSetting struct {
	Switcher  int     `json:"switcher"`            // Geolocation option, default: 1, 1: Ask, 2: Disable
	BaseOnIP  bool    `json:"base_on_ip"`          // Whether to generate based on IP, default: true
	Latitude  float64 `json:"latitude,omitempty"`  // Latitude
	Longitude float64 `json:"longitude,omitempty"` // Longitude
	Accuracy  float64 `json:"accuracy,omitempty"`  // Accuracy (meters)
}

// LanguageSetting Language Setting
type LanguageSetting struct {
	Switcher int    `json:"switcher"`        // Language option, default: 1, 1: Match IP, 2: Custom
	Value    string `json:"value,omitempty"` // Required if switcher is 2
}

// ResolutionSetting Resolution Setting
type ResolutionSetting struct {
	Switcher int    `json:"switcher"`     // Resolution option, default: 1, 1: Real, 2: Custom
	ID       string `json:"id,omitempty"` // Required if switcher is 2
}

// FontSetting Font Setting
type FontSetting struct {
	Switcher int    `json:"switcher"`        // Font option, default: 1, 1: Real, 2: Custom
	Value    string `json:"value,omitempty"` // Required if switcher is 2
}

// SwitcherSetting General Switch Setting
type SwitcherSetting struct {
	Switcher int `json:"switcher"` // Switch option
}

// PortScanSetting Port Scan Protection Setting
type PortScanSetting struct {
	Switcher int    `json:"switcher"`        // Port scan protection option, default: 1, 1: Enable, 2: Disable
	Value    string `json:"value,omitempty"` // Local network ports that are allowed to connect
}

// StartupConfig Configuration After Startup
type StartupConfig struct {
	AfterStartup int32    `json:"afterStartup"`           // Settings after startup, default: 1
	AutoOpenUrls []string `json:"autoOpenUrls,omitempty"` // Open specified web addresses
	PlatformUrl  []string `json:"platformUrl,omitempty"`  // Open specified web addresses
}

// UpdateEnvProfile Update Environment Request Parameters
type UpdateEnvProfile struct {
	EnvId              int64            `json:"envId"`                        // Environment ID
	EnvName            string           `json:"envName,omitempty"`            // Environment name, limited to 100 characters
	AccountInfo        *AccountInfo     `json:"accountInfo,omitempty"`        // Environment account information
	AdvancedSetting    *AdvancedSetting `json:"advancedSetting,omitempty"`    // Advanced settings
	AfterStartupConfig *StartupConfig   `json:"afterStartupConfig,omitempty"` // Configuration after environment startup
	BrowserCore        int32            `json:"browserCore,omitempty"`        // Core version number, default: 0-auto match
	Cookies            string           `json:"cookies,omitempty"`            // Cookie
	EnvRemark          string           `json:"envRemark,omitempty"`          // Environment remarks, limited to 1500 characters
	GroupId            int64            `json:"groupId,omitempty"`            // Environment group ID, default: Ungrouped-0
	IsEncrypt          int32            `json:"isEncrypt,omitempty"`          // Whether to enable end-to-end encryption 0-off, 1-on, default 0
	ProxyId            int64            `json:"proxyId,omitempty"`            // Proxy ID, default: 0
	TagIds             []int64          `json:"tagIds,omitempty"`             // Tag IDs, default: None
	UaVersion          int32            `json:"uaVersion,omitempty"`          // UA, default: 0-all
	StartupParams      []string         `json:"startupParams,omitempty"`      // Startup parameters
	DisableAudio       int32            `json:"disableAudio,omitempty"`       // Disable audio playback: 0-off, 1-on
	DisableVideo       int32            `json:"disableVideo,omitempty"`       // Disable video loading: 0-off, 1-on
	DisableImg         int32            `json:"disableImg,omitempty"`         // Disable image loading: 0-off, 1-on
	ImgLimitSize       int32            `json:"imgLimitSize,omitempty"`       // Image limit size, default 10kb
}

// Delete Browser Environment
type ProfileBatchDelete struct {
	EnvId         []int64 `json:"envIds,omitempty"`        // Environment IDs
	removeEnvData bool    `json:"removeEnvData,omitempty"` // Whether to delete configuration files simultaneously, supported in version 2.28.0 and above
}

// Get Environment List
type ProfilePage struct {
	PageNo   int64  `json:"pageNo"`            // Current page, default 1
	PageSize int64  `json:"pageSize"`          // Number of items per page, default 10
	EnvName  string `json:"envName,omitempty"` // Query by environment name
	GroupId  int64  `json:"groupId,omitempty"` // Query by group ID, 0: Ungrouped
	EnvId    int64  `json:"envId,omitempty"`   // Query by environment ID
}

// Get Browser Environment Details
type ProfileDetail struct {
	EnvID int64 `json:"envId"` // Environment ID
}

// Get Browser Core Version List
type ProfileUAVersions struct {
}

// Get Browser Core Version List
type ProfileUAInfo struct {
	OS        int `json:"os"`        // Corresponding different operating systems 1:Windows 2:macOS 3:Android 4:iOS
	OsVersion int `json:"osVersion"` // System version including Windows 7-11, macOS 12-14
	Vendor    int `json:"vendor"`    // Corresponding different browser types 1:Chrome, 2:Firefox
}

// Get Resolution
type ProfileResolutionList struct {
	OS int `json:"os"` // Corresponding different operating systems 1:Windows 2:macOS 3:Android 4:iOS
	UA int `json:"ua"` // ua
}

// Get Configurable Platforms
type ProfilePlatformList struct {
}

// Get Environment Security Lock Status
type ProfileLock struct {
	EnvId int64 `json:"envId"`
}

// ProfileLanguageList
type ProfileLanguageList struct {
	OS int `json:"os"` // Corresponding different operating systems 1:Windows 2:macOS 3:Android 4:iOS
}

// ClearCacheProfile
type ClearCacheProfile struct {
	EnvId        string `json:"envId,omitempty"`    // Environment ID, only one of EnvId or UniqueId can be provided
	UniqueId     int32  `json:"uniqueId,omitempty"` // Environment number, only one of EnvId or UniqueId can be provided
	LocalStorage bool   `json:"localStorage"`       // Whether to clear LocalStorage, default: false
	IndexedDB    bool   `json:"indexedDB"`          // Whether to clear IndexedDB, default: false
	Cookie       bool   `json:"cookie"`             // Whether to clear cookies, default: false
	Extension    bool   `json:"extension"`          // Whether to clear extensions, default: false
}

// RefreshFingerprintProfile
type RefreshFingerprintProfile struct {
	EnvId           string           `json:"envId,omitempty"`    // Environment ID, only one of EnvId or UniqueId can be provided
	UniqueId        int32            `json:"uniqueId,omitempty"` // Environment number, only one of EnvId or UniqueId can be provided
	UaVersion       int32            `json:"uaVersion"`          // UA version
	AdvancedSetting *AdvancedSetting `json:"advancedSetting"`    // Advanced configuration, see advanced creation interface for details
}

// ProfileStatus Get Browser Environment Running Status
type ProfileStatus struct {
	EnvID string `json:"envId"` // Environment ID
}
