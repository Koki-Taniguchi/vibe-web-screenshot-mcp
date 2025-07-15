# Screenshot MCP Server - ウェブページスクリーンショット取得MCPサーバー

[mcp-go](https://github.com/mark3labs/mcp-go)とchromedpを使用してウェブページのスクリーンショットを取得し、base64エンコードして返すMCP（Model Context Protocol）サーバーです。

[English](README.md)

## 機能

- **フルスクリーンショット取得**: ページ全体のスクリーンショットを取得
- **base64エンコード**: 取得したスクリーンショットをbase64形式で出力
- **MCPプロトコル対応**: LLMアプリケーションから簡単に利用可能

## 必要な環境

- Go 1.24以上
- Google Chrome または Chromium ブラウザ

## インストール

1. 依存関係をインストール:
```bash
go mod tidy
```

2. プログラムをビルド:
```bash
go build -o screenshot-server main.go
```

## MCPサーバーとしての使用方法

### 1. サーバーの起動

```bash
./screenshot-server
```

サーバーはstdio transportを使用してMCPプロトコルで通信します。

### 2. 利用可能なツール

#### `full_screenshot`

指定したURLのページ全体のスクリーンショットを取得します。

**パラメータ:**
- `url` (string, 必須): スクリーンショットを取得するURL

**使用例:**
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

### 3. LLMアプリケーションとの統合

このMCPサーバーは、Claude Desktop、Cursor、その他のMCP対応LLMアプリケーションから利用できます。

#### Cursorでの設定例

`.cursor/mcp.json`に以下を追加:

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

## 注意事項

- スクリーンショット取得には時間がかかる場合があります（最大30秒のタイムアウト）
- ページの読み込み完了を待つため、2秒の待機時間が設定されています
- ヘッドレスモードで動作するため、ブラウザウィンドウは表示されません
- MCPサーバーとして動作中は、標準入出力がMCPプロトコル通信に使用されます

## トラブルシューティング

### Chrome/Chromiumが見つからない場合

システムにGoogle ChromeまたはChromiumがインストールされていることを確認してください。

### タイムアウトエラーが発生する場合

ネットワーク接続やページの読み込み速度を確認してください。必要に応じてタイムアウト時間を調整できます。

### MCPクライアントとの接続問題

- サーバーが正しく起動していることを確認
- MCPクライアント側の設定が正しいことを確認
- ログを確認してエラーメッセージを調査

## 開発者向け情報

### 依存関係

- [github.com/mark3labs/mcp-go](https://github.com/mark3labs/mcp-go): MCPプロトコル実装
- [github.com/chromedp/chromedp](https://github.com/chromedp/chromedp): Chrome DevTools Protocol実装

### ビルドとテスト

```bash
# ビルド
go build -o screenshot-server main.go

# toolsの確認
echo -e '{"jsonrpc":"2.0","id":1,"method":"initialize","params":{"protocolVersion":"2024-11-05","capabilities":{},"clientInfo":{"name":"cursor","version":"1.0.0"}}}\n{"jsonrpc":"2.0","method":"initialized"}\n{"jsonrpc":"2.0","id":2,"method":"tools/list"}' | screenshot-server
```
