package profile

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
	"morelogin.com/mcp/operations/types"
	"morelogin.com/mcp/utils"
)

const (
	RefreshProfileFingerprintToolName = "refresh_profile_fingerprint"
)

var RefreshProfileFingerprintTool = func() mcp.Tool {
	options := utils.CombineOptions(
		[]mcp.ToolOption{
			mcp.WithDescription("Refresh fingerprints. Need to update the MoreLogin app to version 2.28.0 and above."),
		},
		[]mcp.ToolOption{
			mcp.WithString(
				"envId",
				mcp.Description("There is and can only be one Profile ID and profile number passed."),
			),
			mcp.WithNumber(
				"uniqueId",
				mcp.Description("There is and can only be one profile ID and profie number passed."),
			),
			mcp.WithNumber(
				"uaVersion",
				mcp.Description("UA version"),
			),
			AdvancedBaseOptionsToolOption,
		},
	)
	return mcp.NewTool(RefreshProfileFingerprintToolName, options...)
}()

func RefreshProfileFingerprintHandleFunc(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	apiUrl := fmt.Sprintf("/api/env/fingerprint/refresh")
	moreLoginClient := utils.NewMoreLoginClient("POST", apiUrl, utils.WithPayload(request.Params.Arguments))
	pull := &types.CommonResponse[any]{}
	return moreLoginClient.HandleMCPResult(pull)
}
