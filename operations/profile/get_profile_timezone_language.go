package profile

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
	"morelogin.com/mcp/operations/types"
	"morelogin.com/mcp/utils"
)

const (
	GetTimeZoneAndLanguageToolName = "get_time_zone_and_language"
)

var GetTimeZoneAndLanguageStatusTool = func() mcp.Tool {
	options := utils.CombineOptions(
		[]mcp.ToolOption{
			mcp.WithDescription("Get available timezones and languages. You need to update MoreLogin client to version 2.14.0 or above."),
		},
		[]mcp.ToolOption{
			mcp.WithNumber(
				"os",
				mcp.Description("Operating system type 1：Windows，2：macOS，3：Android，4：IOS"),
			),
		},
	)
	return mcp.NewTool(GetTimeZoneAndLanguageToolName, options...)
}()

func GetTimeZoneAndLanguageToolHandleFunc(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	apiUrl := fmt.Sprintf("/api/env/base/list")
	moreLoginClient := utils.NewMoreLoginClient("POST", apiUrl, utils.WithPayload(request.Params.Arguments))
	data := &types.CommonResponse[types.TimeZoneAndLanguage]{}
	return moreLoginClient.HandleMCPResult(data)
}
