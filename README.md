# 🧠 HoloDebug ∞ OMEGA — Real-Time Holographic Debugger for Golang

![License](https://img.shields.io/badge/license-Apache%202.0-blue)
![Build](https://img.shields.io/github/actions/workflow/status/AUSP59/HolographicDebugger/ci.yml?branch=main)
![Go Version](https://img.shields.io/badge/go-1.21+-00ADD8)
![Status](https://img.shields.io/badge/status-Production--Ready-brightgreen)

> **Visualize and interact with Go concurrency in real-time using a holographic 3D web interface.**  
> Debug goroutines, channels, states and metrics live — without modifying your codebase.

---

## 🔭 Features

- ⚙️ **Real-time WebSocket debugging engine** for Go apps
- 🔒 **Signed JSON exports** with SHA-256 integrity verification
- 🧱 **Modular plugin system** for extending logic at runtime
- 🌐 **3D Dashboard** (Three.js) with interactive controls and HUD
- 📊 **Live metrics**: goroutines, memory, trace stats
- 🧑‍💻 **Multi-user support** with token-based auth
- 🧪 **CI-ready**, Dockerized and cross-platform install scripts
- 📚 **Full documentation**, architecture, whitepaper, API reference
- 🌈 **Accessible** interface with theme and language switching

---

## 🚀 Quick Start

### 🔧 Build and Run

```bash
make run
# or manually:
go run cmd/server/main.go
Then open: http://localhost:8080

You’ll see a live 3D view of your app's goroutines and events.

🔐 JSON Export & Signing
bash
Copiar
Editar
curl http://localhost:8080/export > data.json
go run cmd/server/hashutil.go data.json
You’ll get a .sig file with the SHA-256 hash.

📦 Install
Linux/macOS:
bash
Copiar
Editar
chmod +x install.sh
./install.sh
Windows (PowerShell):
powershell
Copiar
Editar
Set-ExecutionPolicy Bypass -Scope Process -Force
.\install.ps1
🧩 Plugins
Build custom plugins using Go’s plugin system:

go
Copiar
Editar
package main
import "fmt"
func Name() string { return "MyPlugin" }
func Init()        { fmt.Println("Initialized!") }
Then compile and drop into the /plugins folder:

bash
Copiar
Editar
go build -buildmode=plugin -o myplugin.so myplugin.go
See docs/plugins.md for full details.

📁 Project Structure
bash
Copiar
Editar
├── cmd/server/         # Go backend (WebSocket, export, hash)
├── web/static/         # Web frontend (3D UI)
├── docs/               # Docs, whitepaper, API, architecture
├── plugins/            # Plugin system
├── release/            # Binary packages (future)
├── install.sh/ps1      # OS installers
📚 Documentation
🧾 Whitepaper

📐 Architecture

🔌 Plugin Guide

🔐 Security Policy

🧭 Roadmap

Navigate all docs via docs/index.html or host with GitHub Pages.

🛡 License & Policy
📜 LICENSE — Apache 2.0

🤝 ETHICS.md

🚨 SECURITY.md

👥 Contributing
We welcome contributions!
See CONTRIBUTING.md and open issues or submit pull requests.

🌐 Links
🔗 GitHub: github.com/AUSP59/HolographicDebugger

📦 Releases: Coming soon

🌍 Docs: GitHub Pages support ready

Built with ❤️ to rethink Go debugging through immersive visual systems.
