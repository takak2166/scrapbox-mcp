# Scrapbox MCP

[English](#english) | [日本語](#japanese)

## English

### Overview

Scrapbox MCP is a Message Control Protocol (MCP) server implementation for Scrapbox. This project provides three different implementations for comparison:

1. Implementation using `mark3labs/mcp-go` framework
2. Implementation using `ktr0731/go-mcp`
3. Implementation using `metoro-io/mcp-golang`

Both implementations provide a standardized interface for interacting with Scrapbox through various tools and commands.

### Features

- Scrapbox API integration
- Tool-based command interface
- Environment-based configuration
- Support for various Scrapbox operations:
  - Page retrieval
  - Page listing
  - Page search
  - Page creation for URL generation

### Prerequisites

- Go 1.23.5 or later
- Scrapbox API access token (SID)

### Installation

1. Clone the repository:
```bash
git clone https://github.com/takak2166/scrapbox-mcp.git
cd scrapbox-mcp
```

2. Install dependencies:
```bash
go mod download
```

3. Build the project:
```bash
# For mark3labs/mcp-go implementation
go build -o bin/scrapbox-mcp cmd/mcp-go/main.go

# For ktr0731/go-mcp implementation
go build -o bin/scrapbox-mcp-go cmd/go-mcp/main.go

# For metoro-io/mcp-golang implementation
go build -o bin/scrapbox-mcp-golang cmd/mcp-golang/main.go
```

### Configuration

Create a `.env` file in the project root with the following variables:

```env
SCRAPBOX_PROJECT=your-project-name
SCRAPBOX_SID=your-SID
```

### Usage

Run the server:

```bash
# For mark3labs/mcp-go implementation
./bin/scrapbox-mcp

# For ktr0731/go-mcp implementation
./bin/scrapbox-mcp-go

# For metoro-io/mcp-golang implementation
./bin/scrapbox-mcp-golang
```

### Project Structure

```
.
├── cmd/              # Command-line applications
│   ├── mcp-go/      # mark3labs/mcp-go implementation
│   ├── go-mcp/      # ktr0731/go-mcp implementation
│   └── mcp-golang/  # metoro-io/mcp-golang implementation
├── internal/         # Private application code
├── pkg/             # Public library code
└── bin/             # Compiled binaries
```


## Japanese

### 概要

Scrapbox MCPは、Scrapbox用のMessage Control Protocol（MCP）サーバー実装です。このプロジェクトは比較のために3つの異なる実装を提供しています：

1. `mark3labs/mcp-go`フレームワークを使用した実装
2. `ktr0731/go-mcp`を使用した実装
3. `metoro-io/mcp-golang`を使用した実装

両方の実装は、様々なツールやコマンドを通じてScrapboxと対話するための標準化されたインターフェースを提供します。

### 機能

- Scrapbox APIとの統合
- ツールベースのコマンドインターフェース
- 環境ベースの設定管理
- 以下のScrapbox操作をサポート：
  - ページの取得
  - ページの一覧表示
  - ページの検索
  - ページ作成URLの生成

### 必要条件

- Go 1.23.5以上
- Scrapbox APIアクセストークン（SID）

### インストール

1. リポジトリのクローン:
```bash
git clone https://github.com/takak2166/scrapbox-mcp.git
cd scrapbox-mcp
```

2. 依存関係のインストール:
```bash
go mod download
```

3. プロジェクトのビルド:
```bash
# mark3labs/mcp-go実装用
go build -o bin/scrapbox-mcp cmd/mcp-go/main.go

# ktr0731/go-mcp実装用
go build -o bin/scrapbox-mcp-go cmd/go-mcp/main.go

# metoro-io/mcp-golang実装用
go build -o bin/scrapbox-mcp-golang cmd/mcp-golang/main.go
```

### 設定

プロジェクトのルートディレクトリに`.env`ファイルを作成し、以下の変数を設定してください：

```env
SCRAPBOX_PROJECT=プロジェクト名
SCRAPBOX_SID=SID
```

### 使用方法

サーバーの起動:

```bash
# mark3labs/mcp-go実装用
./bin/scrapbox-mcp

# ktr0731/go-mcp実装用
./bin/scrapbox-mcp-go

# metoro-io/mcp-golang実装用
./bin/scrapbox-mcp-golang
```

### プロジェクト構造

```
.
├── cmd/              # コマンドラインアプリケーション
│   ├── mcp-go/      # mark3labs/mcp-go実装
│   ├── go-mcp/      # ktr0731/go-mcp実装
│   └── mcp-golang/  # metoro-io/mcp-golang実装
├── internal/         # プライベートなアプリケーションコード
├── pkg/             # パブリックなライブラリコード
└── bin/             # コンパイル済みバイナリ
```