package proxy

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
	"morelogin.com/mcp/operations/types"
	"morelogin.com/mcp/utils"
)

const (
	GetProxyListToolName = "get_proxy_list"
)

var GetProxyListTool = func() mcp.Tool {
	options := utils.CombineOptions(
		[]mcp.ToolOption{
			mcp.WithDescription("Queries information about added proxies. The MoreLogin application needs to be updated to version 2.9.0 and above."),
		},
		[]mcp.ToolOption{
			mcp.WithNumber(
				"pageNo",
				mcp.Description("Current page, default 1"),
				mcp.DefaultNumber(1),
				mcp.Required(),
			),
			mcp.WithNumber(
				"pageSize",
				mcp.Description("Number of articles per page, default 10"),
				mcp.DefaultNumber(10),
				mcp.Min(1),
				mcp.Max(200),
				mcp.Required(),
			),
			mcp.WithNumber(
				"id",
				mcp.Description("Search by profile name"),
			),
			mcp.WithBoolean(
				"isCloudPhoneProxy",
				mcp.Description("Whether it is a proxy that can be used by the cloud phone profile true: yes; false: no"),
			),
			mcp.WithNumber(
				"proxyCategoryType",
				mcp.Description("Query by proxy type 1: Platform proxy; 2: Self-owned proxy"),
			),
			mcp.WithNumber(
				"proxyCheckStatus",
				mcp.Description("Query by detection status 0: pending detection, 1: successful detection, 2: failed detection, 3: unknown error"),
			),
			mcp.WithString(
				"proxyIp",
				mcp.Description("Proxy IP, proxyProvider for 16/17/18 can be empty, the other is required"),
				mcp.MaxLength(50),
			),
			mcp.WithString(
				"proxyName",
				mcp.Description("Proxy name (up to 600 characters)"),
				mcp.MaxLength(50),
			),
			mcp.WithNumber(
				"proxyProviders",
				mcp.Description("Query by proxy provider 0: None, 4: Oxylabs, 5: Proxys.io, 7: Luminati, 8: Lumauto, 9: Oxylabsauto, 10: Trojan, 11: Shadowsocks, 13: ABCPROXY, 14: LunaProxy. 15: IPHTML, 16: PiaProxy, 17: 922S5 18：360Proxy Default: 0"),
				mcp.Items(map[string]interface{}{
					"type": "integer",
				}),
			),
			mcp.WithNumber(
				"proxyStatus",
				mcp.Description("Query by proxy status 0: Normal , 1: Pending assignment , 2: Upgrading , 3: Expired , 4: Expiring soon"),
				mcp.Max(1),
				mcp.Min(0),
			),
			mcp.WithArray(
				"proxyTypes",
				mcp.Description("Query by proxy type 0: http, 1: https, 2: socks5, 3: ssh"),
				mcp.Items(map[string]interface{}{
					"type": "integer",
				}),
			),
		},
	)
	return mcp.NewTool(GetProxyListToolName, options...)
}()

func GetProxyListHandleFunc(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	apiUrl := fmt.Sprintf("/api/proxyInfo/page")
	moreLoginClient := utils.NewMoreLoginClient("GET", apiUrl, utils.WithPayload(request.Params.Arguments))
	data := &types.CommonResponse[types.CommonPageRes[types.ProxyInfo]]{}
	return moreLoginClient.HandleMCPResult(data)
}
