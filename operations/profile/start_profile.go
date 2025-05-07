package profile

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
	"morelogin.com/mcp/operations/types"
	"morelogin.com/mcp/utils"
)

const (
	StartProfileToolName = "profile_start"
)

var StartProfileTool = func() mcp.Tool {
	options := utils.CombineOptions(
		[]mcp.ToolOption{
			mcp.WithDescription("Used to start the profile, you need to specify the profile ID, after successful startup," +
				" you can get the profile debug interface for the execution of selenium and puppeteer automation. " +
				"Selenium needs to use the Webdriver that matches the corresponding kernel version, " +
				"you can get the path of the corresponding Webdriver in the return value after starting the profile. " +
				"You need to update MoreLogin client to version 2.15.0 or above."),
		},
		[]mcp.ToolOption{
			mcp.WithString(
				"envId",
				mcp.Description("Profile I Send the profile ID or the profile order number. If both are sent, the profile ID takes precedence."),
				mcp.Required(),
			),
			mcp.WithBoolean(
				"isHeadless",
				mcp.Description("Whether to start the browser profile in headless mode.Note: Need to upgrade to V2.36.0 and above"),
			),
			mcp.WithNumber(
				"uniqueId",
				mcp.Description("Profile order number Send the profile ID or the profile order number. If both are sent, the profile ID takes precedence."),
			),
			mcp.WithString(
				"encryptKey",
				mcp.Description("Private key, mandatory when enabling end-to-end encryption in the profile"),
				mcp.DefaultString(""),
			),
			mcp.WithBoolean(
				"cdpEvasion",
				mcp.Description("Enable CDP Feature Evasion:When enabled, this can reduce the risk of detection by platforms.Default: false Note: Requires version V2.36.0 or above."),
				mcp.DefaultBool(false),
			),
		},
	)
	return mcp.NewTool(StartProfileToolName, options...)
}()

func StartProfileToolHandleFunc(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	apiUrl := fmt.Sprintf("/api/env/start")
	moreLoginClient := utils.NewMoreLoginClient("POST", apiUrl, utils.WithPayload(request.Params.Arguments))
	data := &types.CommonResponse[types.StartProfileRes]{}
	return moreLoginClient.HandleMCPResult(data)
}
