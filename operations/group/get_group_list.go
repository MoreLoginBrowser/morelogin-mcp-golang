package group

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
	"morelogin.com/mcp/operations/types"
	"morelogin.com/mcp/utils"
)

const (
	GetGroupListToolName = "get_group_list"
)

var GetGroupListTool = func() mcp.Tool {
	options := utils.CombineOptions(
		[]mcp.ToolOption{
			mcp.WithDescription("Query group information, group information includes group ID and group name, where group ID is used to set groups for the profile. The MoreLogin application needs to be updated to version 2.14.0 and above."),
		},
		[]mcp.ToolOption{
			mcp.WithNumber(
				"pageNo",
				mcp.Description("Current page, default 1"),
				mcp.DefaultNumber(1),
				mcp.Required(),
			),
			mcp.WithNumber(
				"pageSize",
				mcp.Description("Number of articles per page, default 10"),
				mcp.DefaultNumber(20),
				mcp.Min(1),
				mcp.Max(200),
				mcp.Required(),
			),
			mcp.WithString(
				"groupName",
				mcp.Description("Search by Group Name"),
			),
		},
	)
	return mcp.NewTool(GetGroupListToolName, options...)
}()

func GetGroupListHandleFunc(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	apiUrl := fmt.Sprintf("/api/envgroup/page")
	moreLoginClient := utils.NewMoreLoginClient("POST", apiUrl, utils.WithPayload(request.Params.Arguments))
	data := &types.CommonResponse[types.CommonPageRes[types.GroupListRes]]{}
	return moreLoginClient.HandleMCPResult(data)
}
