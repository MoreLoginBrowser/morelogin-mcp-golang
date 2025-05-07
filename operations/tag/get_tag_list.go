package tag

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
	"morelogin.com/mcp/operations/types"
	"morelogin.com/mcp/utils"
)

const (
	GetTagListToolName = "get_tag_list"
)

var GetTagListTool = func() mcp.Tool {
	options := utils.CombineOptions(
		[]mcp.ToolOption{
			mcp.WithDescription("Query tag information, tag information includes tag ID and tag name, where tag ID is used to set tags for the profile. The MoreLogin application needs to be updated to version 2.14.0 and above."),
		},
	)
	return mcp.NewTool(GetTagListToolName, options...)
}()

func GetTagListHandleFunc(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	apiUrl := fmt.Sprintf("/api/envtag/all")
	moreLoginClient := utils.NewMoreLoginClient("GET", apiUrl, utils.WithPayload(request.Params.Arguments))
	data := &types.CommonResponse[[]types.TagListRes]{}
	return moreLoginClient.HandleMCPResult(data)
}
