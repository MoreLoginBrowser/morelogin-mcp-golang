package proxy

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
	"morelogin.com/mcp/operations/types"
	"morelogin.com/mcp/utils"
)

const (
	CreateProxyToolName = "create_proxy"
)

var CreateProxyTool = func() mcp.Tool {
	options := utils.CombineOptions(
		[]mcp.ToolOption{
			mcp.WithDescription("dd proxy, after adding successfully, can be used to proxy the profile. The name of proxys can not be duplicated. After the successful creation of the proxy, the proxy ID will be returned. The MoreLogin application needs to be updated to version 2.9.0 and above."),
		},
		[]mcp.ToolOption{
			mcp.WithString(
				"city",
				mcp.Description("City"),
				mcp.MaxLength(100),
			),
			mcp.WithString(
				"country",
				mcp.Description("Country (check appendix for specific country codes), proxyProvider is 16/17/18 required"),
				mcp.MaxLength(100),
			),
			mcp.WithNumber(
				"encryptionType",
				mcp.Description("Encryption method, not null when proxyProvider is 11. 1：aes-128-gcm，2：aes-192-gcm，3：aes-256-gcm，4：aes-128-cfb，5：aes-192-cfb，6：aes-256-cfb，7：aes-128-ctr，8：aes-192-ctr，9：aes-256-ctr，10：rc4-md5，11：chacha20-ietf，12：xchacha20，13：chacha20-ietf-poly1305，14：xchacha20-ietf-poly1305"),
				mcp.Max(14),
				mcp.Min(1),
			),
			mcp.WithBoolean(
				"ipChangeAction",
				mcp.Description("Whether to enable IP change monitoring true: on, false: off Default: false"),
				mcp.DefaultBool(false),
			),
			mcp.WithNumber(
				"ipMonitor",
				mcp.Description("IP change monitoring 0: No access, 1: Warning"),
				mcp.DefaultNumber(0),
				mcp.Max(1),
				mcp.Min(0),
			),
			mcp.WithString(
				"username",
				mcp.Description("User name (up to 200 characters)"),
				mcp.MaxLength(200),
			),
			mcp.WithString(
				"password",
				mcp.Description("Password (up to 100 characters)"),
				mcp.MaxLength(100),
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
				"proxyPort",
				mcp.Description("Proxy port (only numbers from 1-65535 are supported), proxyProvider 16/17/18 can be null, others cannot be null."),
				mcp.Max(65535),
				mcp.Min(1),
			),
			mcp.WithNumber(
				"proxyProvider",
				mcp.Description("proxyProvider 0：http，1：https，2：socks5，3：ssh，4：Oxylabs，5：Proxys.io，7：Luminati，8：Lumauto，9：Oxylabsauto，10：Trojan，11：Shadowsocks，13：ABCPROXY，14：LunaProxy，15：IPHTML，16：PiaProxy，17：922S5，18：360Proxy"),
				mcp.Max(18),
				mcp.Min(0),
			),
			mcp.WithNumber(
				"proxyType",
				mcp.Description("Proxy type, 0: http, 1: https, not null if proxyProvider is 7/8"),
				mcp.Max(1),
				mcp.Min(0),
			),
			mcp.WithString(
				"refreshUrl",
				mcp.Description("Refresh URL"),
				mcp.MaxLength(300),
			),
			mcp.WithString(
				"state",
				mcp.Description("State/Province"),
				mcp.MaxLength(50),
			),
		},
	)
	return mcp.NewTool(CreateProxyToolName, options...)
}()

func CreateProxyHandleFunc(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	apiUrl := fmt.Sprintf("/api/proxyInfo/add")
	moreLoginClient := utils.NewMoreLoginClient("POST", apiUrl, utils.WithPayload(request.Params.Arguments))
	data := &types.CommonResponse[string]{}
	return moreLoginClient.HandleMCPResult(data)
}
