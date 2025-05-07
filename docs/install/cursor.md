stdio mode
```json
{
  "mcpServers": {
    "morelogin": {
      "command": "mcp-morelogin",
      "env": {
        "API_BASE": "http://127.0.0.1:40000"
      }
    }
  }
}
```

sse mode

start mcp server through sse
```bash
mcp-morelogin -transport sse
```
```json
{
  "mcpServers": {
    "morelogin": {
      "url": "http://127.0.0.1:40000",
    }
  }
}
```
![sse](./images/cursor.png)