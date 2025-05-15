package profile

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
	"morelogin.com/mcp/operations/types"
	"morelogin.com/mcp/utils"
)

const (
	GetProfileSecurityLockStatusToolName = "get_profile_security_lock_status"
)

var GetProfileSecurityLockStatusTool = func() mcp.Tool {
	options := utils.CombineOptions(
		[]mcp.ToolOption{
			mcp.WithDescription("Get the locking status of the profile security lock. You need to update MoreLogin client to version 2.14.0 or above."),
		},
		[]mcp.ToolOption{
			mcp.WithString(
				"envId",
				mcp.Description("Profile ID"),
			),
		},
	)
	return mcp.NewTool(GetProfileSecurityLockStatusToolName, options...)
}()

func GetProfileSecurityLockStatusToolHandleFunc(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	apiUrl := fmt.Sprintf("/api/env/lock/query")
	moreLoginClient := utils.NewMoreLoginClient("POST", apiUrl, utils.WithPayload(request.Params.Arguments))
	data := &types.CommonResponse[types.ProfileSecurityLockStatus]{}
	return moreLoginClient.HandleMCPResult(data)
}
