package profile

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
	"morelogin.com/mcp/operations/types"
	"morelogin.com/mcp/utils"
)

const (
	RemoveProfileToolName = "profile_remove"
)

var RemoveProfileTool = func() mcp.Tool {
	options := utils.CombineOptions(
		[]mcp.ToolOption{
			mcp.WithDescription(" Delete unwanted profiles and retrieve them in the Trash within 7 days after deletion. You need to update MoreLogin client to version 2.14.0 or above."),
		},
		[]mcp.ToolOption{
			mcp.WithArray(
				"envIds",
				mcp.Description("Profile ID Send the profile ID or the profile order number. If both are sent, the profile ID takes precedence."),
				mcp.Items(map[string]interface{}{
					"type": "string",
				}),
			),
			mcp.WithBoolean(
				"removeEnvData",
				mcp.Description("Whether to delete profiles at the same time, supported by version 2.28.0 and above"),
			),
		},
	)
	return mcp.NewTool(RemoveProfileToolName, options...)
}()

func RemoveProfileToolHandleFunc(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	apiUrl := fmt.Sprintf("/api/env/removeToRecycleBin/batch")
	moreLoginClient := utils.NewMoreLoginClient("POST", apiUrl, utils.WithPayload(request.Params.Arguments))
	pull := &types.CommonResponse[bool]{}
	return moreLoginClient.HandleMCPResult(pull)
}
