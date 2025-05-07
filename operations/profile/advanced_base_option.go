package profile

import "github.com/mark3labs/mcp-go/mcp"

var AccountInfoToolOption = mcp.WithObject("accountInfo",
	mcp.Description("Profile account information"),
	mcp.Properties(map[string]interface{}{
		"platformId": map[string]interface{}{
			"type":        "integer",
			"description": "Platform ID, 9999 for customized platforms, other platform IDs can be obtained through the 'get configurable platforms' interface",
			"required":    true,
		},
		"customerUrl": map[string]interface{}{
			"type":        "string",
			"description": "Customized platform URL, required when platform ID=9999, must be a legitimate URL address",
		},
		"username": map[string]interface{}{
			"type":        "string",
			"description": "User name, length limit 64 characters",
			"maxLength":   64,
		},
		"password": map[string]interface{}{
			"type":        "string",
			"description": "Password, length limit 50 characters",
			"maxLength":   50,
		},
		"otpSecret": map[string]interface{}{
			"type":        "string",
			"description": "2FA key, generates secondary verification codes for websites",
		},
		"siteId": map[string]interface{}{
			"type":        "integer",
			"description": "Site ID, available through the 'get configurable platform' interface",
		},
	}),
)

var AdvancedBaseOptionsToolOption = mcp.WithObject("advancedSetting",
	mcp.Description("Advanced setting"),
	mcp.Properties(map[string]interface{}{
		"ua": map[string]interface{}{
			"type":        "string",
			"description": "Customize the profile UA, the format should be uploaded according to the standard format",
		},
		"time_zone": map[string]interface{}{
			"type":        "object",
			"description": "The timezone, e.g., Asia/Shanghai ",
			"properties": map[string]interface{}{
				"switcher": map[string]interface{}{
					"type":        "integer",
					"description": "Timezone option, default 1, 1: Match IP, 2: Custom",
					"default":     1,
				},
				"value": map[string]interface{}{
					"type":        "string",
					"description": "When the switcher sends 2, it is mandatory to send the corresponding time zone ID",
				},
			},
		},
		"web_rtc": map[string]interface{}{
			"type":        "object",
			"description": "WebRTC",
			"properties": map[string]interface{}{
				"switcher": map[string]interface{}{
					"type":        "integer",
					"description": "WebRTC option, default 2, 1: Privacy, 2: Replacement, 3: Real, 4: Disabled, 5: Forwarding",
					"default":     2,
				},
			},
		},
		"geo_location": map[string]interface{}{
			"type":        "object",
			"description": "Geographic location",
			"properties": map[string]interface{}{
				"switcher": map[string]interface{}{
					"type":        "integer",
					"description": "Geographical location option, default 1, 1: Inquiry, 2: Disabled",
					"default":     1,
				},
				"base_on_ip": map[string]interface{}{
					"type":        "boolean",
					"description": "Whether to generate based on IP, default: true",
					"default":     true,
				},
				"latitude": map[string]interface{}{
					"type":        "number",
					"description": "Latitude, mandatory when generating corresponding geographic locations without IP",
				},
				"longitude": map[string]interface{}{
					"type":        "number",
					"description": "Longitude, mandatory when generating corresponding geographic locations without IP",
				},
				"accuracy": map[string]interface{}{
					"type":        "number",
					"description": "Accuracy (in meters), must be transmitted when generating corresponding geographic locations without IP",
				},
			},
		},
		"language": map[string]interface{}{
			"type":        "object",
			"description": "Language",
			"properties": map[string]interface{}{
				"switcher": map[string]interface{}{
					"type":        "integer",
					"description": "Language options, default 1, 1: Match IP, 2: Custom",
					"default":     1,
				},
				"value": map[string]interface{}{
					"type":        "string",
					"description": "When the switcher sends 2, it is mandatory to send the language ID",
				},
			},
		},
		"resolution": map[string]interface{}{
			"type":        "object",
			"description": "Resolution",
			"properties": map[string]interface{}{
				"switcher": map[string]interface{}{
					"type":        "integer",
					"description": "Resolution option, default 1: Real, 2: Custom",
				},
				"id": map[string]interface{}{
					"type":        "string",
					"description": "When the switcher sends 2, it is mandatory to send the resolution ID",
				},
			},
		},
		"font": map[string]interface{}{
			"type":        "object",
			"description": "Font",
			"properties": map[string]interface{}{
				"switcher": map[string]interface{}{
					"type":        "integer",
					"description": "Font options, default 1: Real, 2: Custom",
				},
				"value": map[string]interface{}{
					"type":        "string",
					"description": "Switcher must be filled in when selecting custom font, separated by commas",
				},
			},
		},
		"canvas": map[string]interface{}{
			"type":        "object",
			"description": "Canvas",
			"properties": map[string]interface{}{
				"switcher": map[string]interface{}{
					"type":        "integer",
					"description": "Canvas option, default 1: Noise, 2: Real",
				},
			},
		},
		"webgl_image": map[string]interface{}{
			"type":        "object",
			"description": "WebGL image",
			"properties": map[string]interface{}{
				"switcher": map[string]interface{}{
					"type":        "integer",
					"description": "WebGL image options, default 1, 1: Noise, 2: Real",
					"default":     1,
				},
			},
		},
		"webgl_metadata": map[string]interface{}{
			"type":        "object",
			"description": "WebGL metadata",
			"properties": map[string]interface{}{
				"switcher": map[string]interface{}{
					"type":        "integer",
					"description": "WebGL metadata options, default 3, 1: Real, 2: Turn off hardware acceleration, 3: Custom",
					"default":     3,
				},
			},
		},
		"audio_context": map[string]interface{}{
			"type":        "object",
			"description": "AudioContext",
			"properties": map[string]interface{}{
				"switcher": map[string]interface{}{
					"type":        "integer",
					"description": "AudioContext option, default 1: Noise, 2: Realistic",
				},
			},
		},
		"media_device": map[string]interface{}{
			"type":        "object",
			"description": "Media device",
			"properties": map[string]interface{}{
				"switcher": map[string]interface{}{
					"type":        "integer",
					"description": "Media device options, default 1, 1: Noise, 2: Real",
					"default":     1,
				},
			},
		},
		"client_rects": map[string]interface{}{
			"type":        "object",
			"description": "ClientRects",
			"properties": map[string]interface{}{
				"switcher": map[string]interface{}{
					"type":        "integer",
					"description": "ClientRects option, default 1, 1: Noise, 2: Real",
					"default":     1,
				},
			},
		},
		"speech_voise": map[string]interface{}{
			"type":        "object",
			"description": "SpeechVoices",
			"properties": map[string]interface{}{
				"switcher": map[string]interface{}{
					"type":        "integer",
					"description": "SpeechVoices option, default 1, 1: Privacy, 2: Real",
					"default":     1,
				},
			},
		},
		"hardware_concurrency": map[string]interface{}{
			"type":        "integer",
			"description": "Hardware concurrency, default 4 (real concurrency), 2, 3, 4, 6, 8, 10, 12",
			"default":     4,
		},
		"memery_device": map[string]interface{}{
			"type":        "integer",
			"description": "Device memory, default 8 (real memory), 2, 4, 6, 8",
			"default":     8,
		},
		"do_not_track": map[string]interface{}{
			"type":        "integer",
			"description": "Do Not Track, default 2, 1: on, 2: off",
			"default":     2,
		},
		"bluetooth": map[string]interface{}{
			"type":        "object",
			"description": "Bluetooth",
			"properties": map[string]interface{}{
				"switcher": map[string]interface{}{
					"type":        "integer",
					"description": "Bluetooth option, default 1: Privacy, 2: Authenticity",
					"default":     1,
				},
			},
		},
		"battery": map[string]interface{}{
			"type":        "object",
			"description": "Battery",
			"properties": map[string]interface{}{
				"switcher": map[string]interface{}{
					"type":        "integer",
					"description": "Battery option, default 1, 1: Privacy, 2: Authenticity",
					"default":     1,
				},
			},
		},
		"port_scan_protection": map[string]interface{}{
			"type":        "object",
			"description": "Port scan protection",
			"properties": map[string]interface{}{
				"switcher": map[string]interface{}{
					"type":        "integer",
					"description": "Port scan protection option, default 1, 1: On, 2: Off",
					"default":     1,
				},
				"value": map[string]interface{}{
					"type":        "string",
					"description": "Allowed local network ports to be connected",
				},
			},
		},
		"os_version": map[string]interface{}{
			"type":        "string",
			"description": "MacOS system version, e.g.: macOS 12, macOS 13, macOS 14",
		},
		"web_gpu": map[string]interface{}{
			"type":        "object",
			"description": "WebGPU",
			"properties": map[string]interface{}{
				"switcher": map[string]interface{}{
					"type":        "integer",
					"description": "WebGPU options, default 1, 1: WebGL-based matching, 2: true, 3: disabled",
					"default":     1,
				},
			},
		},
	}),
)
