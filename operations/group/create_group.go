package group

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
	"morelogin.com/mcp/operations/types"
	"morelogin.com/mcp/utils"
)

const (
	CreateGroupToolName = "create_group"
)

var CreateGroupTool = func() mcp.Tool {
	options := utils.CombineOptions(
		[]mcp.ToolOption{
			mcp.WithDescription("dd group, after adding successfully, can be used to group the profile. The name of groups can not be duplicated. After the successful creation of the group, the group ID will be returned. The MoreLogin application needs to be updated to version 2.9.0 and above."),
		},
		[]mcp.ToolOption{
			mcp.WithString(
				"groupName",
				mcp.Description("The name of the group"),
				mcp.Required(),
				mcp.MaxLength(100),
				mcp.MinLength(1),
			),
		},
	)
	return mcp.NewTool(CreateGroupToolName, options...)
}()

func CreateGroupHandleFunc(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	apiUrl := fmt.Sprintf("/api/envgroup/create")
	moreLoginClient := utils.NewMoreLoginClient("POST", apiUrl, utils.WithPayload(request.Params.Arguments))
	data := &types.CommonResponse[int64]{}
	return moreLoginClient.HandleMCPResult(data)
}
