package tag

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
	"morelogin.com/mcp/operations/types"
	"morelogin.com/mcp/utils"
)

const (
	CreateTagToolName = "create_tag"
)

var CreateTagTool = func() mcp.Tool {
	options := utils.CombineOptions(
		[]mcp.ToolOption{
			mcp.WithDescription("Add tags, when added successfully, can be used to set tags for the profiles. The tag ID will be returned when created successfully. The MoreLogin application needs to be updated to version 2.14.0 and above."),
		},
		[]mcp.ToolOption{
			mcp.WithString(
				"tagName",
				mcp.Description("The name of the tag"),
				mcp.Required(),
				mcp.MaxLength(100),
				mcp.MinLength(1),
			),
		},
	)
	return mcp.NewTool(CreateTagToolName, options...)
}()

func CreateTagHandleFunc(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	apiUrl := fmt.Sprintf("/api/envtag/create")
	moreLoginClient := utils.NewMoreLoginClient("POST", apiUrl, utils.WithPayload(request.Params.Arguments))
	data := &types.CommonResponse[types.TagListRes]{}
	return moreLoginClient.HandleMCPResult(data)
}
