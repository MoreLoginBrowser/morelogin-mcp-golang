package profile

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
	"morelogin.com/mcp/operations/types"
	"morelogin.com/mcp/utils"
)

const (
	GetResolutionToolName = "get_resolution_list"
)

var GetResolutionTool = func() mcp.Tool {
	options := utils.CombineOptions(
		[]mcp.ToolOption{
			mcp.WithDescription("Get available resolution based on UA. You need to update MoreLogin client to version 2.14.0 or above."),
		},
		[]mcp.ToolOption{
			mcp.WithNumber(
				"os",
				mcp.Description("Corresponding to different operating systems. 1: Windows, 2: macOS, 3: Android, 4: iOS"),
				mcp.Required(),
			),
			mcp.WithString(
				"ua",
				mcp.Description("user agent"),
			),
		},
	)
	return mcp.NewTool(GetResolutionToolName, options...)
}()

func GetResolutionToolHandleFunc(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	apiUrl := fmt.Sprintf("/api/env/base/resolution/list")
	moreLoginClient := utils.NewMoreLoginClient("POST", apiUrl, utils.WithPayload(request.Params.Arguments))
	data := &types.CommonResponse[[]types.Resolution]{}
	return moreLoginClient.HandleMCPResult(data)
}
