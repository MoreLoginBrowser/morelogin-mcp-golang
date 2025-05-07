package tag

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
	"morelogin.com/mcp/operations/types"
	"morelogin.com/mcp/utils"
)

const (
	DeleteTagToolName = "delete_tag"
)

var DeleteTagTool = func() mcp.Tool {
	options := utils.CombineOptions(
		[]mcp.ToolOption{
			mcp.WithDescription("Delete the specified tag. The MoreLogin application needs to be updated to version 2.14.0 and above.\n\n"),
		},
		[]mcp.ToolOption{
			mcp.WithArray(
				"ids",
				mcp.Description("tag ID"),
				mcp.Items(map[string]interface{}{
					"type": "integer",
				}),
				mcp.Required(),
			),
		},
	)
	return mcp.NewTool(DeleteTagToolName, options...)
}()

func DeleteTagHandleFunc(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	apiUrl := fmt.Sprintf("/api/envtag/delete")
	moreLoginClient := utils.NewMoreLoginClient("POST", apiUrl, utils.WithPayload(request.Params.Arguments))
	data := &types.CommonResponse[bool]{}
	return moreLoginClient.HandleMCPResult(data)
}
