package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"morelogin.com/mcp/operations/group"
	"morelogin.com/mcp/operations/profile"
	"morelogin.com/mcp/operations/proxy"
	"morelogin.com/mcp/operations/tag"
	"morelogin.com/mcp/utils"
	"os"

	"github.com/mark3labs/mcp-go/server"
)

var (
	Version = utils.Version
)

func newMCPServer() *server.MCPServer {
	return server.NewMCPServer(
		"MoreLogin-mcp",
		Version,
		server.WithToolCapabilities(true),
		server.WithLogging(),
	)
}

func addTools(s *server.MCPServer) {
	// Profile Tools
	s.AddTool(profile.QuickCreateTool, profile.QuickCreateToolHandleFunc)
	s.AddTool(profile.StartProfileTool, profile.StartProfileToolHandleFunc)
	s.AddTool(profile.CloseProfileTool, profile.CloseProfileToolHandleFunc)
	s.AddTool(profile.AdvancedCreateProfileTool, profile.AdvancedCreateProfileToolHandleFunc)
	s.AddTool(profile.ModifyProfileTool, profile.ModifyProfileToolHandleFunc)
	s.AddTool(profile.RemoveProfileTool, profile.RemoveProfileToolHandleFunc)
	s.AddTool(profile.GetProfilePageTool, profile.GetProfilePageToolHandleFunc)
	s.AddTool(profile.GetProfileDetailTool, profile.GetProfileDetailToolHandleFunc)
	s.AddTool(profile.GetBrowserKernelTool, profile.GetBrowserKernelToolHandleFunc)
	s.AddTool(profile.GetBrowserUATool, profile.GetBrowserUAToolHandleFunc)
	s.AddTool(profile.GetResolutionTool, profile.GetResolutionToolHandleFunc)
	s.AddTool(profile.GetPlatformTool, profile.GetPlatformToolHandleFunc)
	s.AddTool(profile.GetProfileSecurityLockStatusTool, profile.GetProfileSecurityLockStatusToolHandleFunc)
	s.AddTool(profile.GetTimeZoneAndLanguageStatusTool, profile.GetTimeZoneAndLanguageToolHandleFunc)
	s.AddTool(profile.RemoveProfileLocalCacheTool, profile.RemoveProfileLocalToolHandleFunc)
	s.AddTool(profile.RefreshProfileFingerprintTool, profile.RefreshProfileFingerprintHandleFunc)
	s.AddTool(profile.GetProfileRunningStatusTool, profile.GetProfileRunningStatusToolHandleFunc)

	// Tag Tools
	s.AddTool(tag.CreateTagTool, tag.CreateTagHandleFunc)
	s.AddTool(tag.GetTagListTool, tag.GetTagListHandleFunc)
	s.AddTool(tag.ModifyTagTool, tag.ModifyTagHandleFunc)
	s.AddTool(tag.DeleteTagTool, tag.DeleteTagHandleFunc)

	// Group Tools
	s.AddTool(group.CreateGroupTool, group.CreateGroupHandleFunc)
	s.AddTool(group.GetGroupListTool, group.GetGroupListHandleFunc)
	s.AddTool(group.ModifyGroupTool, group.ModifyGroupHandleFunc)
	s.AddTool(group.DeleteGroupTool, group.DeleteGroupHandleFunc)

	// Proxy Tools
	s.AddTool(proxy.CreateProxyTool, proxy.CreateProxyHandleFunc)
	s.AddTool(proxy.GetProxyListTool, proxy.GetProxyListHandleFunc)
	s.AddTool(proxy.ModifyProxyTool, proxy.ModifyProxyHandleFunc)
	s.AddTool(proxy.DeleteProxyTool, proxy.DeleteProxyHandleFunc)
}

func run(transport, addr string) error {
	s := newMCPServer()
	addTools(s)

	switch transport {
	case "stdio":
		if err := server.ServeStdio(s); err != nil {
			if errors.Is(err, context.Canceled) {
				return nil
			}
			return err
		}
	case "sse":
		srv := server.NewSSEServer(s, server.WithBaseURL(fmt.Sprintf("http://%s", addr)),
			server.WithUseFullURLForMessageEndpoint(true))
		log.Printf("SSE server listening on %s", addr)
		if err := srv.Start(addr); err != nil {
			if errors.Is(err, context.Canceled) {
				return nil
			}
			return fmt.Errorf("server error: %v", err)
		}
	default:
		return fmt.Errorf(
			"invalid transport type: %s. Must be 'stdio' or 'sse'",
			transport,
		)
	}
	return nil
}

func main() {
	apiBase := flag.String("api-base", "", "MoreLogin API base URL (default: http://127.0.0.1:40000)")
	showVersion := flag.Bool("version", false, "Show version information")
	transport := flag.String("transport", "sse", "Transport type (stdio or sse)")
	addr := flag.String("sse-address", "localhost:8000", "The host and port to start the sse server on")
	flag.Parse()

	if *showVersion {
		fmt.Printf("MoreLogin MCP Server\n")
		fmt.Printf("Version: %s\n", Version)
		os.Exit(0)
	}
	if *apiBase != "" {
		utils.SetApiBase(*apiBase)
	}

	if err := run(*transport, *addr); err != nil {
		panic(err)
	}
}
