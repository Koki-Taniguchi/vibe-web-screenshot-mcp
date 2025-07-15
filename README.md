# Screenshot MCP Server

A Model Context Protocol (MCP) server that captures full-page screenshots of web pages using [mcp-go](https://github.com/mark3labs/mcp-go) and chromedp, returning them as base64-encoded images.

[Japanese](README.ja.md)

## Features

- **Full-page screenshot capture**: Captures entire webpage screenshots
- **Base64 encoding**: Returns screenshots in base64 format
- **MCP protocol support**: Easy integration with LLM applications

## Requirements

- Go 1.24 or higher
- Google Chrome or Chromium browser

## Installation

1. Install dependencies:
```bash
go mod tidy
```

2. Build the program:
```bash
go build -o screenshot-server main.go
```

## Usage as MCP Server

### 1. Starting the Server

```bash
./screenshot-server
```

The server uses stdio transport to communicate via the MCP protocol.

### 2. Available Tools

#### `full_screenshot`

Captures a full-page screenshot of the specified URL.

**Parameters:**
- `url` (string, required): The URL to capture a screenshot of

**Example usage:**
```json
{
  "method": "tools/call",
  "params": {
    "name": "full_screenshot",
    "arguments": {
      "url": "https://example.com"
    }
  }
}
```

### 3. Integration with LLM Applications

This MCP server can be used with Claude Desktop, Cursor, and other MCP-compatible LLM applications.

#### Configuration Example for Cursor

Add the following to `.cursor/mcp.json`:

```json
{
  "mcpServers": {
    "screener": {
      "command": "/path/to/screenshot-server",
      "args": []
    }
  }
}
```

## Important Notes

- Screenshot capture may take time (maximum 30-second timeout)
- 2-second wait time is set to ensure page loading completion
- Runs in headless mode, so no browser window is displayed
- When running as MCP server, stdin/stdout is used for MCP protocol communication

## Troubleshooting

### Chrome/Chromium Not Found

Make sure Google Chrome or Chromium is installed on your system.

### Timeout Errors

Check your network connection and page loading speed. Timeout duration can be adjusted if needed.

### MCP Client Connection Issues

- Verify the server is running correctly
- Check MCP client-side configuration
- Review logs for error messages

## Developer Information

### Dependencies

- [github.com/mark3labs/mcp-go](https://github.com/mark3labs/mcp-go): MCP protocol implementation
- [github.com/chromedp/chromedp](https://github.com/chromedp/chromedp): Chrome DevTools Protocol implementation

### Build and Test

```bash
# Build
go build -o screenshot-server main.go

# Test tools
echo -e '{"jsonrpc":"2.0","id":1,"method":"initialize","params":{"protocolVersion":"2024-11-05","capabilities":{},"clientInfo":{"name":"cursor","version":"1.0.0"}}}\n{"jsonrpc":"2.0","method":"initialized"}\n{"jsonrpc":"2.0","id":2,"method":"tools/list"}' | screenshot-server
```
