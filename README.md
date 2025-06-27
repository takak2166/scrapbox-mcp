# Scrapbox MCP

[English](#english) | [日本語](#japanese)

## English

### Overview

Scrapbox MCP is a Message Control Protocol (MCP) server implementation for Scrapbox. This project provides four different implementations for comparison:

1. Implementation using `mark3labs/mcp-go` framework
2. Implementation using `ktr0731/go-mcp`
3. Implementation using `metoro-io/mcp-golang`
4. Implementation using the official Model Context Protocol Go SDK (recommended)

All implementations provide a standardized interface for interacting with Scrapbox through various tools and commands.

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
# Build all implementations
make build

# Or build specific implementation
make build-mcp-go      # mark3labs/mcp-go implementation
make build-go-mcp      # ktr0731/go-mcp implementation
make build-mcp-golang  # metoro-io/mcp-golang implementation
make build-official-mcp # Official Go SDK implementation (recommended)
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
make run-mcp-go

# For ktr0731/go-mcp implementation
make run-go-mcp

# For metoro-io/mcp-golang implementation
make run-mcp-golang

# For official Go SDK implementation (recommended)
make run-official-mcp
```

Or run directly:

```bash
# For mark3labs/mcp-go implementation
./bin/scrapbox-mcp

# For ktr0731/go-mcp implementation
./bin/scrapbox-mcp-go

# For metoro-io/mcp-golang implementation
./bin/scrapbox-mcp-golang

# For official Go SDK implementation (recommended)
./bin/scrapbox-mcp-official
```

### Make Commands

```bash
make help           # Show available commands
make build          # Build all implementations
make build-mcp-go   # Build mark3labs/mcp-go implementation
make build-go-mcp   # Build ktr0731/go-mcp implementation
make build-mcp-golang # Build metoro-io/mcp-golang implementation
make build-official-mcp # Build official Go SDK implementation (recommended)
make run-mcp-go     # Build and run mark3labs/mcp-go implementation
make run-go-mcp     # Build and run ktr0731/go-mcp implementation
make run-mcp-golang # Build and run metoro-io/mcp-golang implementation
make run-official-mcp # Build and run official Go SDK implementation (recommended)
make clean          # Clean build artifacts
```

### Project Structure

```
.
├── cmd/              # Command-line applications
│   ├── mcp-go/      # mark3labs/mcp-go implementation
│   ├── go-mcp/      # ktr0731/go-mcp implementation
│   ├── mcp-golang/  # metoro-io/mcp-golang implementation
│   └── official-mcp/ # Official Go SDK implementation (recommended)
├── internal/         # Private application code
├── pkg/             # Public library code
└── bin/             # Compiled binaries
```

## Japanese

### 概要

Scrapbox MCP は、Scrapbox 用の Message Control Protocol（MCP）サーバー実装です。このプロジェクトは比較のために 4 つの異なる実装を提供しています：

1. `mark3labs/mcp-go`フレームワークを使用した実装
2. `ktr0731/go-mcp`を使用した実装
3. `metoro-io/mcp-golang`を使用した実装
4. 公式 Model Context Protocol Go SDK を使用した実装（推奨）

すべての実装は、様々なツールやコマンドを通じて Scrapbox と対話するための標準化されたインターフェースを提供します。

### 機能

- Scrapbox API との統合
- ツールベースのコマンドインターフェース
- 環境ベースの設定管理
- 以下の Scrapbox 操作をサポート：
  - ページの取得
  - ページの一覧表示
  - ページの検索
  - ページ作成 URL の生成

### 必要条件

- Go 1.23.5 以上
- Scrapbox API アクセストークン（SID）

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
# 全ての実装をビルド
make build

# または特定の実装をビルド
make build-mcp-go      # mark3labs/mcp-go実装
make build-go-mcp      # ktr0731/go-mcp実装
make build-mcp-golang  # metoro-io/mcp-golang実装
make build-official-mcp # 公式Go SDK実装（推奨）
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
make run-mcp-go

# ktr0731/go-mcp実装用
make run-go-mcp

# metoro-io/mcp-golang実装用
make run-mcp-golang

# 公式Go SDK実装用（推奨）
make run-official-mcp
```

または直接実行:

```bash
# mark3labs/mcp-go実装用
./bin/scrapbox-mcp

# ktr0731/go-mcp実装用
./bin/scrapbox-mcp-go

# metoro-io/mcp-golang実装用
./bin/scrapbox-mcp-golang

# 公式Go SDK実装用（推奨）
./bin/scrapbox-mcp-official
```

### Make コマンド

```bash
make help           # 利用可能なコマンドを表示
make build          # 全ての実装をビルド
make build-mcp-go   # mark3labs/mcp-go実装をビルド
make build-go-mcp   # ktr0731/go-mcp実装をビルド
make build-mcp-golang # metoro-io/mcp-golang実装をビルド
make build-official-mcp # 公式Go SDK実装をビルド（推奨）
make run-mcp-go     # mark3labs/mcp-go実装をビルドして実行
make run-go-mcp     # ktr0731/go-mcp実装をビルドして実行
make run-mcp-golang # metoro-io/mcp-golang実装をビルドして実行
make run-official-mcp # 公式Go SDK実装をビルドして実行（推奨）
make clean          # ビルド成果物を削除
```

### プロジェクト構造

```
.
├── cmd/              # コマンドラインアプリケーション
│   ├── mcp-go/      # mark3labs/mcp-go実装
│   ├── go-mcp/      # ktr0731/go-mcp実装
│   ├── mcp-golang/  # metoro-io/mcp-golang実装
│   └── official-mcp/ # 公式Go SDK実装（推奨）
├── internal/         # プライベートなアプリケーションコード
├── pkg/             # パブリックなライブラリコード
└── bin/             # コンパイル済みバイナリ
```
