# 🧠 Mnemosyne

> *“Memory is the mother of wisdom.”*  
> — Mnemosyne, Titaness of Memory

**Mnemosyne** is a command-line tool that syncs your ChatGPT conversations into a structured **Notion knowledge database**.  
It automatically organizes chats, adds tags, and groups related topics to help you build your personal second brain.

---

## ✨ Features

- 📤 **Sync ChatGPT conversations** (from exported data or API)
- 🏷️ **Auto-tagging** based on topic keywords (Security, AI, DevOps, etc.)
- 🗂️ **Group conversations** by domain or project
- 🧩 **Notion integration** — each conversation becomes a Notion page
- 🔁 **Incremental sync** — only uploads new or changed conversations
- 🕐 **Automated schedule support** via cron or GitHub Actions

---

## 🧰 Tech Stack

- **Language:** Go
- **Storage:** Notion Database
- **Config:** YAML/JSON-based
- **Tags:** Keyword-based, extendable via config or AI model

---

## 🚀 Getting Started

### 1. Prerequisites

- Go 1.22+
- Notion API integration key
- Exported ChatGPT data (`.json` or `.zip`)

### 2. Installation

```bash
git clone https://github.com/ximura/mnemosyne.git
cd mnemosyne
go build -o mnemo ./cmd/mnemosyne
./bin/mnemo sync --config .mnemo.yaml
```