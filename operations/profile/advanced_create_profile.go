package profile

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
	"morelogin.com/mcp/operations/types"
	"morelogin.com/mcp/utils"
)

const (
	AdvancedCreateToolName = "advanced_create_profile"
)

var AdvancedCreateProfileTool = func() mcp.Tool {
	options := utils.CombineOptions(
		[]mcp.ToolOption{
			mcp.WithDescription("Advanced create profile, support to set platform account password, cookie, fingerprint information and so on for the profile. Returns the profile ID after successful creation. You need to update MoreLogin client to version 2.14.0 or above."),
		},
		[]mcp.ToolOption{
			// 必填字段
			mcp.WithNumber(
				"browserTypeId",
				mcp.Description("Browser type: 1 - Chrome, 2 - Firefox"),
				mcp.Required(),
			),
			mcp.WithNumber(
				"operatorSystemId",
				mcp.Description("Operating system type: 1 - Windows, 2 - macOS, 3 - Android, 4 - IOS"),
				mcp.Required(),
				mcp.Min(1),
				mcp.Max(4),
			),
			mcp.WithString(
				"envName",
				mcp.Description("Profile name, length limit 100 characters"),
				mcp.MaxLength(100),
			),
			AccountInfoToolOption,
			AdvancedBaseOptionsToolOption,
			mcp.WithObject("afterStartupConfig",
				mcp.Description("Configuration after profile startup"),
				mcp.Properties(map[string]interface{}{
					"afterStartup": map[string]interface{}{
						"type":        "integer",
						"description": "Set after startup, default 1: 1: Continue browsing the last opened webpage, 2: Open the specified webpage, 3: Open the specified webpage and platform, 4: Continue browsing the last opened webpage and platform",
					},
					"autoOpenUrls": map[string]interface{}{
						"type":        "array",
						"description": "Open the specified webpage address, which must be a valid URL address",
						"items": map[string]interface{}{
							"type": "string",
						},
					},
				}),
			),
			mcp.WithNumber("browserCore",
				mcp.Description("Kernel version number, default 0 - Auto Match"),
			),
			mcp.WithString("cookies",
				mcp.Description("Cookie"),
			),
			mcp.WithString("envRemark",
				mcp.Description("Profile remarks, length limit 1500 characters"),
				mcp.MaxLength(1500),
			),
			mcp.WithNumber("groupId",
				mcp.Description("Profile group ID, default: not grouped -0, limit minimum value 0"),
				mcp.Min(0),
			),
			mcp.WithNumber("isEncrypt",
				mcp.Description("Whether to enable 'End-to-End Encryption', 0: off, 1: on, default 0"),
				mcp.Min(0),
				mcp.Max(1),
			),
			mcp.WithNumber("proxyId",
				mcp.Description("Proxy ID, default: 0, limit minimum value 0"),
				mcp.Min(0),
			),
			mcp.WithArray("tagIds",
				mcp.Description("Tag ID, default: null"),
				mcp.Items(map[string]interface{}{
					"type": "integer",
				}),
			),
			mcp.WithNumber("uaVersion",
				mcp.Description("UA, default: 0-all"),
			),
			mcp.WithArray("startupParams",
				mcp.Description("Profile startup parameters"),
				mcp.Items(map[string]interface{}{
					"type": "string",
				}),
			),
			mcp.WithNumber("disableAudio",
				mcp.Description("Disable audio playback: default 0, 0 off, 1 on"),
				mcp.Min(0),
				mcp.Max(1),
			),
			mcp.WithNumber("disableVideo",
				mcp.Description("Disable video loading: default 0, 0 off, 1 on"),
				mcp.Min(0),
				mcp.Max(1),
			),
			mcp.WithNumber("disableImg",
				mcp.Description("Disable image loading: default 0, 0 off, 1 on"),
				mcp.Min(0),
				mcp.Max(1),
			),
			mcp.WithNumber("imgLimitSize",
				mcp.Description("Image limit size; default 10kb"),
				mcp.Min(0),
			),
		},
	)

	return mcp.NewTool(AdvancedCreateToolName, options...)
}()

func AdvancedCreateProfileToolHandleFunc(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	apiUrl := fmt.Sprintf("/api/env/create/advanced")
	moreLoginClient := utils.NewMoreLoginClient("POST", apiUrl, utils.WithPayload(request.Params.Arguments))
	data := &types.CommonResponse[types.QuickCreateResponse]{}
	return moreLoginClient.HandleMCPResult(data)
}
