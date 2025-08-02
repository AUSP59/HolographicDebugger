
package main

import (
    "encoding/json"
    "io/ioutil"
    "log"
    "net/http"
    "os"
    "path/filepath"
    "runtime"
    "runtime/pprof"
    "strings"
    "time"

    "github.com/gorilla/websocket"
)

type Config struct {
    Port             int      `json:"port"`
    RefreshInterval  int      `json:"refresh_interval_ms"`
    AuthToken        string   `json:"auth_token"`
    AllowOrigins     []string `json:"allow_origins"`
    ExportDir        string   `json:"export_dir"`
    ExportFormat     string   `json:"export_format"`
    Language         string   `json:"language"`
}

type GoroutineInfo struct {
    ID    int    `json:"id"`
    State string `json:"state"`
}

type Snapshot struct {
    Time       string          `json:"time"`
    Count      int             `json:"count"`
    Goroutines []GoroutineInfo `json:"goroutines"`
}

var (
    upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}
    config   Config
    snapshots []Snapshot
)

func classify(stack string) string {
    if strings.Contains(stack, "select") {
        return "waiting"
    } else if strings.Contains(stack, "syscall") {
        return "syscall"
    } else if strings.Contains(stack, "runtime.") {
        return "internal"
    }
    return "running"
}

func traceGoroutines() []GoroutineInfo {
    buf := make([]byte, 1<<20)
    n := runtime.Stack(buf, true)
    lines := strings.Split(string(buf[:n]), "\n")
    var result []GoroutineInfo
    id := 0
    for i := 0; i < len(lines)-1; i++ {
        if strings.HasPrefix(lines[i], "goroutine ") {
            id++
            state := classify(lines[i+1])
            result = append(result, GoroutineInfo{ID: id, State: state})
        }
    }
    return result
}

func saveSnapshot(snapshot Snapshot) {
    snapshots = append(snapshots, snapshot)
}

func exportJSON(w http.ResponseWriter, r *http.Request) {
    ts := time.Now().Format("20060102_150405")
    filename := filepath.Join(config.ExportDir, "session_"+ts+".json")
    f, err := os.Create(filename)
    if err != nil {
        http.Error(w, "Failed to export", 500)
        return
    }
    defer f.Close()
    json.NewEncoder(f).Encode(snapshots)
    w.Write([]byte("Exported to " + filename))
}

func serveWS(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println("WebSocket upgrade error:", err)
        return
    }
    defer conn.Close()

    for {
        snapshot := Snapshot{
            Time:       time.Now().Format(time.RFC3339),
            Count:      runtime.NumGoroutine(),
            Goroutines: traceGoroutines(),
        }
        saveSnapshot(snapshot)
        msg, _ := json.Marshal(snapshot)
        if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
            break
        }
        time.Sleep(time.Duration(config.RefreshInterval) * time.Millisecond)
    }
}

func main() {
    raw, err := ioutil.ReadFile("config.json")
    if err != nil {
        log.Fatal("Cannot read config:", err)
    }
    if err := json.Unmarshal(raw, &config); err != nil {
        log.Fatal("Invalid config format:", err)
    }

    http.HandleFunc("/ws", serveWS)
    http.HandleFunc("/export", exportJSON)
    http.Handle("/", http.FileServer(http.Dir("web/static")))
    go func() {
        log.Println("pprof available at /debug/pprof/")
        log.Println(http.ListenAndServe("localhost:6060", nil)) // expose pprof
    }()
    addr := ":" + os.Getenv("PORT")
    if addr == ":" {
        addr = ":" + string(rune(config.Port))
    }
    log.Println("HoloDebug SUPRA running on http://localhost" + addr)
    log.Fatal(http.ListenAndServe(addr, nil))
}
