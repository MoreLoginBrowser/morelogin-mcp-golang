package profile

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
	"morelogin.com/mcp/operations/types"
	"morelogin.com/mcp/utils"
)

const (
	CloseProfileToolName = "profile_close"
)

var CloseProfileTool = func() mcp.Tool {
	options := utils.CombineOptions(
		[]mcp.ToolOption{
			mcp.WithDescription("Stopping the specified environment requires the specified profile ID. You need to update MoreLogin client to version 2.15.0 or above."),
		},
		[]mcp.ToolOption{
			mcp.WithString(
				"envId",
				mcp.Description("Profile ID Send the profile ID or the profile order number. If both are sent, the profile ID takes precedence."),
			),
			mcp.WithNumber(
				"uniqueId",
				mcp.Description("Profile order number Send the profile ID or the profile order number. If both are sent, the profile ID takes precedence."),
			),
		},
	)
	return mcp.NewTool(CloseProfileToolName, options...)
}()

func CloseProfileToolHandleFunc(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	apiUrl := fmt.Sprintf("/api/env/close")
	moreLoginClient := utils.NewMoreLoginClient("POST", apiUrl, utils.WithPayload(request.Params.Arguments))
	data := &types.CommonResponse[types.ProfileEnvIdRes]{}
	return moreLoginClient.HandleMCPResult(data)
}
