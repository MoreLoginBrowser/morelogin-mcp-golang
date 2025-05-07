package profile

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
	"morelogin.com/mcp/operations/types"
	"morelogin.com/mcp/utils"
)

const (
	GetProfilePageToolName = "get_profile_page"
)

var GetProfilePageTool = func() mcp.Tool {
	options := utils.CombineOptions(
		[]mcp.ToolOption{
			mcp.WithDescription("Queries the added profile information. Users can query only the profile information to which they have access. You need to update MoreLogin client to version 2.14.0 or above."),
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
				mcp.DefaultNumber(10),
				mcp.Min(1),
				mcp.Max(200),
				mcp.Required(),
			),
			mcp.WithString(
				"envName",
				mcp.Description("Search by profile name"),
			),
			mcp.WithNumber(
				"envId",
				mcp.Description("Search by profile id"),
			),
			mcp.WithNumber(
				"groupId",
				mcp.Description("Query by group ID, 0: not grouped"),
			),
		},
	)
	return mcp.NewTool(GetProfilePageToolName, options...)
}()

func GetProfilePageToolHandleFunc(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	apiUrl := fmt.Sprintf("/api/env/page")
	moreLoginClient := utils.NewMoreLoginClient("POST", apiUrl, utils.WithPayload(request.Params.Arguments))
	data := &types.CommonResponse[types.CommonPageRes[types.PageProfileInfo]]{}
	return moreLoginClient.HandleMCPResult(data)
}
