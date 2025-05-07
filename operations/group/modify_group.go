package group

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
	"morelogin.com/mcp/operations/types"
	"morelogin.com/mcp/utils"
)

const (
	ModifyGroupToolName = "modify_group"
)

var ModifyGroupTool = func() mcp.Tool {
	options := utils.CombineOptions(
		[]mcp.ToolOption{
			mcp.WithDescription("Interface description: The group name can be modified and the name cannot be duplicated. The MoreLogin application needs to be updated to version 2.14.0 and above."),
		},
		[]mcp.ToolOption{
			mcp.WithNumber(
				"id",
				mcp.Description("ID of groups"),
				mcp.Required(),
			),
			mcp.WithString(
				"groupName",
				mcp.Description("The name of the group"),
				mcp.Required(),
				mcp.MaxLength(100),
				mcp.MinLength(1),
			),
		},
	)
	return mcp.NewTool(ModifyGroupToolName, options...)
}()

func ModifyGroupHandleFunc(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	apiUrl := fmt.Sprintf("/api/envgroup/edit")
	moreLoginClient := utils.NewMoreLoginClient("POST", apiUrl, utils.WithPayload(request.Params.Arguments))
	data := &types.CommonResponse[bool]{}
	return moreLoginClient.HandleMCPResult(data)
}
