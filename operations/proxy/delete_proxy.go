package proxy

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
	"morelogin.com/mcp/operations/types"
	"morelogin.com/mcp/utils"
)

const (
	DeleteProxyToolName = "delete_proxy"
)

var DeleteProxyTool = func() mcp.Tool {
	options := utils.CombineOptions(
		[]mcp.ToolOption{
			mcp.WithDescription("Batch delete proxies. The MoreLogin application needs to be updated to version 2.9.0 and above."),
		},
		[]mcp.ToolOption{
			mcp.WithArray(
				"ids",
				mcp.Description("proxy ID"),
				mcp.Items(map[string]interface{}{
					"type": "integer",
				}),
				mcp.Required(),
			),
		},
	)
	return mcp.NewTool(DeleteProxyToolName, options...)
}()

func DeleteProxyHandleFunc(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	apiUrl := fmt.Sprintf("/api/proxyInfo/delete")
	moreLoginClient := utils.NewMoreLoginClient("POST", apiUrl, utils.WithPayload(request.Params.Arguments))
	data := &types.CommonResponse[bool]{}
	return moreLoginClient.HandleMCPResult(data)
}
