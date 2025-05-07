package tag

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
	"morelogin.com/mcp/operations/types"
	"morelogin.com/mcp/utils"
)

const (
	ModifyTagToolName = "modify_tag"
)

var ModifyTagTool = func() mcp.Tool {
	options := utils.CombineOptions(
		[]mcp.ToolOption{
			mcp.WithDescription(" Allow users to modify the name of the tags. The MoreLogin application needs to be updated to version 2.14.0 and above."),
		},
		[]mcp.ToolOption{
			mcp.WithNumber(
				"id",
				mcp.Description("ID of tags"),
				mcp.Required(),
			),
			mcp.WithString(
				"tagName",
				mcp.Description("The name of the tag"),
				mcp.Required(),
				mcp.MaxLength(100),
				mcp.MinLength(1),
			),
		},
	)
	return mcp.NewTool(ModifyTagToolName, options...)
}()

func ModifyTagHandleFunc(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	apiUrl := fmt.Sprintf("/api/envtag/edit")
	moreLoginClient := utils.NewMoreLoginClient("POST", apiUrl, utils.WithPayload(request.Params.Arguments))
	data := &types.CommonResponse[bool]{}
	return moreLoginClient.HandleMCPResult(data)
}
