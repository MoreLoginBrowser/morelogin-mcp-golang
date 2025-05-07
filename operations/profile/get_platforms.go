package profile

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
	"morelogin.com/mcp/operations/types"
	"morelogin.com/mcp/utils"
)

const (
	GetPlatformToolName = "get_platform_list"
)

var GetPlatformTool = func() mcp.Tool {
	options := utils.CombineOptions(
		[]mcp.ToolOption{
			mcp.WithDescription(" Get available platform information. You need to update MoreLogin client to version 2.14.0 or above."),
		},
	)
	return mcp.NewTool(GetPlatformToolName, options...)
}()

func GetPlatformToolHandleFunc(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	apiUrl := fmt.Sprintf("/api/system/platform/list")
	moreLoginClient := utils.NewMoreLoginClient("GET", apiUrl)
	data := &types.CommonResponse[[]types.Platform]{}
	return moreLoginClient.HandleMCPResult(data)
}
