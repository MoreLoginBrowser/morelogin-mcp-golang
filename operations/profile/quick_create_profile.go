package profile

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
	"morelogin.com/mcp/operations/types"
	"morelogin.com/mcp/utils"
)

const (
	// ProfileQuickCreateToolName is the name of the tool
	QuickCreateToolName = "profile_quick_create"
)

var QuickCreateTool = func() mcp.Tool {
	options := utils.CombineOptions(
		[]mcp.ToolOption{
			mcp.WithDescription("Quickly create an environment"),
		},
		[]mcp.ToolOption{
			mcp.WithNumber(
				"browserTypeId",
				mcp.Description("The browserTypeId of the create request,1：Chrome，2： Firefox"),
				mcp.DefaultNumber(1),
				mcp.Min(1),
				mcp.Max(2),
				mcp.Required(),
			),
			mcp.WithNumber(
				"operatorSystemId",
				mcp.Description("The operatorSystemId of the create request,1：Windows，2：macOS，3：Android，4：IOS"),
				mcp.DefaultNumber(1),
				mcp.Min(1),
				mcp.Max(4),
				mcp.Required(),
			),
			mcp.WithNumber(
				"quantity",
				mcp.Description("The operatorSystemId of the create request,1-50"),
				mcp.DefaultNumber(1),
				mcp.Min(1),
				mcp.Max(50),
				mcp.Required(),
			),
			mcp.WithNumber(
				"browserCore",
				mcp.Description("The browserCore of the create request,default:0,"),
				mcp.DefaultNumber(0),
				mcp.Required(),
			),
			mcp.WithNumber(
				"groupId",
				mcp.Description("The groupId of the create request"),
				mcp.DefaultNumber(0),
			),
			mcp.WithNumber(
				"isEncrypt",
				mcp.Description("The isEncrypt of the create request,default:0,0:close,1:open"),
				mcp.DefaultNumber(0),
			),
		},
	)
	return mcp.NewTool(QuickCreateToolName, options...)
}()

func QuickCreateToolHandleFunc(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	apiUrl := fmt.Sprintf("/api/env/create/quick")
	moreLoginClient := utils.NewMoreLoginClient("POST", apiUrl, utils.WithPayload(request.Params.Arguments))
	pull := &types.QuickCreateResponse{}
	return moreLoginClient.HandleMCPResult(pull)
}
