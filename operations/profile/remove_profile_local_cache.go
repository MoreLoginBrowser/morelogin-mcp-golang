package profile

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
	"morelogin.com/mcp/operations/types"
	"morelogin.com/mcp/utils"
)

const (
	RemoveProfileLocalCacheToolName = "remove_profile_local_cache"
)

var RemoveProfileLocalCacheTool = func() mcp.Tool {
	options := utils.CombineOptions(
		[]mcp.ToolOption{
			mcp.WithDescription("Queries the added profile information. Users can query only the profile information to which they have access. You need to update MoreLogin client to version 2.14.0 or above."),
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
			mcp.WithBoolean(
				"localStorage",
				mcp.Description("Whether to clear LocalStorage, default false"),
				mcp.DefaultBool(false),
			),
			mcp.WithBoolean(
				"indexedDB",
				mcp.Description("Whether to clear cookies, default false"),
				mcp.DefaultBool(false),
			),
			mcp.WithBoolean(
				"cookie",
				mcp.Description("Whether to clear cookies, default false"),
				mcp.DefaultBool(false),
			),
			mcp.WithBoolean(
				"extension",
				mcp.Description("Whether to clear the extension, the default false true: clear the extension and extension data, if extensionFile is not passed false: do not clear"),
				mcp.DefaultBool(false),
			),
			mcp.WithBoolean(
				"extensionFile",
				mcp.Description("Whether to clear the extension, the default false, do not clear Note: Requires version V2.36.0 or above."),
				mcp.DefaultBool(false),
			),
		},
	)
	return mcp.NewTool(RemoveProfileLocalCacheToolName, options...)
}()

func RemoveProfileLocalToolHandleFunc(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	apiUrl := fmt.Sprintf("/api/env/removeLocalCache")
	moreLoginClient := utils.NewMoreLoginClient("POST", apiUrl, utils.WithPayload(request.Params.Arguments))
	data := &types.CommonResponse[types.ProfileEnvIdRes]{}
	return moreLoginClient.HandleMCPResult(data)
}
