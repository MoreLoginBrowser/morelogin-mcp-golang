# Morelogin MCP Server

The Morelogin MCP server is an implementation of the Model Context Protocol (MCP) server for Morelogin. It provides a series of tools that interact with the Morelogin API, enabling AI assistants to manage repositories, issues, pull requests, and more.

##Functional Features

-Interact with Morelogin environment, proxy, grouping, and tag APIs
-Configurable API base URL
-Command line flags for easy configuration
-Support individual, organizational, and corporate operations
>

##Installation

###Preconditions

-Go 1.23.0 or higher version
-Morelogin account supports API permissions

###Build from source code

1. Clone repository:
```bash
git clone project
cd mcp-morelogin
```

2. Build the project:
```bash
make build
```
Will /bin/mcp-morelogin

###Install using go install
```bash
go install morelogin.com/ mcp@latest
```

##Usage method

Check the MCP Morelogin version:

```bash
mcp-morelogin --version
```

that it is certified by MCPHub 
url:https://mcphub.com/mcp-servers/moreloginbrowser/morelogin-mcp-golang
