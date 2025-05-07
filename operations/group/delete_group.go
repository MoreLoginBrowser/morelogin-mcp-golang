package group

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
	"morelogin.com/mcp/operations/types"
	"morelogin.com/mcp/utils"
)

const (
	DeleteGroupToolName = "delete_group"
)

var DeleteGroupTool = func() mcp.Tool {
	options := utils.CombineOptions(
		[]mcp.ToolOption{
			mcp.WithDescription("Delete the specified group. The MoreLogin application needs to be updated to version 2.14.0 and above."),
		},
		[]mcp.ToolOption{
			mcp.WithArray(
				"ids",
				mcp.Description("group ID"),
				mcp.Items(map[string]interface{}{
					"type": "integer",
				}),
				mcp.Required(),
			),
			mcp.WithBoolean(
				"isDeleteAllEnv",
				mcp.Description("Whether to delete the profiles in the group at the same time, default: false false: do not delete, true: delete"),
				mcp.DefaultBool(false),
			),
		},
	)
	return mcp.NewTool(DeleteGroupToolName, options...)
}()

func DeleteGroupHandleFunc(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	apiUrl := fmt.Sprintf("/api/envgroup/delete")
	moreLoginClient := utils.NewMoreLoginClient("POST", apiUrl, utils.WithPayload(request.Params.Arguments))
	data := &types.CommonResponse[bool]{}
	return moreLoginClient.HandleMCPResult(data)
}
