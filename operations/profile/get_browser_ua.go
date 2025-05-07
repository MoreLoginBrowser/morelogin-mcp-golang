package profile

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
	"morelogin.com/mcp/operations/types"
	"morelogin.com/mcp/utils"
)

const (
	GetBrowserUAToolName = "get_profile_ua"
)

var GetBrowserUATool = func() mcp.Tool {
	options := utils.CombineOptions(
		[]mcp.ToolOption{
			mcp.WithDescription("Get the available browser profile UA. You need to update MoreLogin client to version 2.14.0 or above."),
		},
		[]mcp.ToolOption{
			mcp.WithNumber(
				"os",
				mcp.Description("Corresponding to different operating systems. 1: Windows, 2: macOS, 3: Android, 4: iOS"),
				mcp.Required(),
			),
			mcp.WithNumber(
				"browserTypeId",
				mcp.Description("Browser Type ID: 1: Chrome, 2: Firefox"),
				mcp.Required(),
			),
			mcp.WithNumber(
				"operatorSystemId",
				mcp.Description("Operator System type: 1: Windows, 2: macOS, 3: Android, 4: iOS"),
				mcp.Required(),
			),
			mcp.WithString(
				"osVersion",
				mcp.Description("System version (e.g. Windows 7-11, macOS 12-14)"),
			),
			mcp.WithNumber(
				"vendor",
				mcp.Description("Corresponding to different browser vendors: 1: Chrome, 2: Firefox"),
				mcp.Required(),
			),
		},
	)
	return mcp.NewTool(GetBrowserUAToolName, options...)
}()

func GetBrowserUAToolHandleFunc(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	apiUrl := fmt.Sprintf("/api/env/advanced/ua/get")
	moreLoginClient := utils.NewMoreLoginClient("POST", apiUrl, utils.WithPayload(request.Params.Arguments))
	data := &types.CommonResponse[types.UA]{}
	return moreLoginClient.HandleMCPResult(data)
}
