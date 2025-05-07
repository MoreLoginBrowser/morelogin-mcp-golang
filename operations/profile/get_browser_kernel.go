package profile

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
	"morelogin.com/mcp/operations/types"
	"morelogin.com/mcp/utils"
)

const (
	GetBrowserKernelToolName = "browser_kernel_list"
)

var GetBrowserKernelTool = func() mcp.Tool {
	options := utils.CombineOptions(
		[]mcp.ToolOption{
			mcp.WithDescription("Queries the available browser kernel versions. You need to update MoreLogin client to version 2.14.0 or above."),
		},
	)
	return mcp.NewTool(GetBrowserKernelToolName, options...)
}()

func GetBrowserKernelToolHandleFunc(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	apiUrl := fmt.Sprintf("/api/env/advanced/ua/versions")
	moreLoginClient := utils.NewMoreLoginClient("GET", apiUrl)
	data := &types.CommonResponse[types.BrowserKernel]{}
	return moreLoginClient.HandleMCPResult(data)
}
