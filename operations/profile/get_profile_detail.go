package profile

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
	"morelogin.com/mcp/operations/types"
	"morelogin.com/mcp/utils"
)

const (
	GetProfileDetailToolName = "get_profile_detail"
)

var GetProfileDetailTool = func() mcp.Tool {
	options := utils.CombineOptions(
		[]mcp.ToolOption{
			mcp.WithDescription("Query profile details. Users can only query profile information for which they have permission. You need to update MoreLogin client to version 2.14.0 or above."),
		},
		[]mcp.ToolOption{
			mcp.WithString(
				"envId",
				mcp.Description("Profile ID Send the profile ID or the profile order number. If both are sent, the profile ID takes precedence."),
			),
		},
	)
	return mcp.NewTool(GetProfileDetailToolName, options...)
}()

func GetProfileDetailToolHandleFunc(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	apiUrl := fmt.Sprintf("/api/env/detail")
	moreLoginClient := utils.NewMoreLoginClient("POST", apiUrl, utils.WithPayload(request.Params.Arguments))
	data := &types.CommonResponse[types.ProfileDetailInfo]{}
	return moreLoginClient.HandleMCPResult(data)
}
