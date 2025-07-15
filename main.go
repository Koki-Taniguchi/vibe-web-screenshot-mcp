package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	s := server.NewMCPServer(
		"Screenshot Server",
		"1.0.0",
		server.WithToolCapabilities(true),
	)

	fullScreenshotTool := mcp.NewTool("full_screenshot",
		mcp.WithDescription("Take a full page screenshot of the specified URL and return it as base64 encoded data"),
		mcp.WithString("url",
			mcp.Required(),
			mcp.Description("URL to take screenshot of"),
		),
	)

	s.AddTool(fullScreenshotTool, func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		url, err := req.RequireString("url")
		if err != nil {
			return mcp.NewToolResultError("URL parameter is required"), nil
		}

		base64Data, err := ScreenshotToBase64(url)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Screenshot capture error: %v", err)), nil
		}

		return mcp.NewToolResultImage(url, base64Data, "image/jpeg"), nil
	})

	if err := server.ServeStdio(s); err != nil {
		fmt.Fprintf(os.Stderr, "Server startup failed: %v\n", err)
		os.Exit(1)
	}
}

func ScreenshotToBase64(url string) (string, error) {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	var buf []byte
	if err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.WaitVisible("body", chromedp.ByQuery),
		chromedp.Sleep(2*time.Second),     // Wait for page to load
		chromedp.FullScreenshot(&buf, 50), // Quality 50 to reduce context window usage
	); err != nil {
		return "", fmt.Errorf("Screenshot capture failed: %v", err)
	}
	base64String := base64.StdEncoding.EncodeToString(buf)
	return base64String, nil
}
