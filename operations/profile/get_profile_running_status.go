package profile

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
	"morelogin.com/mcp/operations/types"
	"morelogin.com/mcp/utils"
)

const (
	GetProfileRunningStatusToolName = "get_profile_running_status"
)

var GetProfileRunningStatusTool = func() mcp.Tool {
	options := utils.CombineOptions(
		[]mcp.ToolOption{
			mcp.WithDescription("This interface is used to query the current status (on or off) of the specified browser profile. The user needs to provide the profile ID, and the interface will return the profile status information. The MoreLogin application needs to be updated to version 2.34.0 or above."),
		},
		[]mcp.ToolOption{
			mcp.WithString(
				"envId",
				mcp.Description("There is and can only be one Profile ID and profile number passed."),
			),
		},
	)
	return mcp.NewTool(GetProfileRunningStatusToolName, options...)
}()

func GetProfileRunningStatusToolHandleFunc(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	apiUrl := fmt.Sprintf("/api/env/status")
	moreLoginClient := utils.NewMoreLoginClient("POST", apiUrl, utils.WithPayload(request.Params.Arguments))
	data := &types.CommonResponse[types.ProfileRunningStatusRes]{}
	return moreLoginClient.HandleMCPResult(data)
}
