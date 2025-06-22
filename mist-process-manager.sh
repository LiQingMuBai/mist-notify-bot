#!/bin/bash

# Configuration ====================================
PROCESSES=(
    "mist-master:/root/seven/master/mist-master"
    "mist-agent:/root/seven/agent/mist-agent"
)
PID_DIR="/tmp/mist-pids"
LOG_FILE="/var/log/mist-process-manager.log"
MONITOR_INTERVAL=60  # Check every 60 seconds
# =================================================

# Initialize PID directory
mkdir -p "$PID_DIR"

# Logging function
log() {
    echo "[$(date '+%Y-%m-%d %H:%M:%S')] $1" | tee -a "$LOG_FILE"
}

# Check if a process is running
check_process() {
    local pid_file="$PID_DIR/$1.pid"
    if [ -f "$pid_file" ]; then
        local pid=$(cat "$pid_file")
        if ps -p "$pid" >/dev/null 2>&1; then
            return 0  # Process is alive
        else
            rm -f "$pid_file"  # Clean up stale PID file
        fi
    fi
    return 1  # Process is not running
}

# Start a process
start_process() {
    local name="$1"
    local path="$2"
    local pid_file="$PID_DIR/$name.pid"

    if check_process "$name"; then
        log "$name is already running (PID: $(cat "$pid_file"))"
        return 0
    fi

    log "Starting $name from $path..."
    nohup "$path" > "/tmp/$name.log" 2>&1 &
    local pid=$!
    echo "$pid" > "$pid_file"
    sleep 1

    if check_process "$name"; then
        log "$name started successfully (PID: $pid)"
    else
        log "Failed to start $name"
        return 1
    fi
}

# Stop a process
stop_process() {
    local name="$1"
    local pid_file="$PID_DIR/$name.pid"

    if check_process "$name"; then
        local pid=$(cat "$pid_file")
        log "Stopping $name (PID: $pid)..."
        kill "$pid" 2>/dev/null && sleep 2
        if check_process "$name"; then
            kill -9 "$pid" 2>/dev/null && sleep 1
        fi
        rm -f "$pid_file"
        log "$name stopped"
    else
        log "$name is not running"
    fi
}

# Restart a process
restart_process() {
    local name="$1"
    local path="$2"
    stop_process "$name"
    start_process "$name" "$path"
}

# Monitor all processes (auto-restart if crashed)
monitor() {
    log "Starting process monitor (interval: ${MONITOR_INTERVAL}s)..."
    while true; do
        for entry in "${PROCESSES[@]}"; do
            IFS=':' read -r name path <<< "$entry"
            if ! check_process "$name"; then
                log "Detected $name is down, restarting..."
                start_process "$name" "$path"
            fi
        done
        sleep "$MONITOR_INTERVAL"
    done
}

# Main logic
case "$1" in
    start)
        for entry in "${PROCESSES[@]}"; do
            IFS=':' read -r name path <<< "$entry"
            start_process "$name" "$path"
        done
        ;;
    stop)
        for entry in "${PROCESSES[@]}"; do
            IFS=':' read -r name path <<< "$entry"
            stop_process "$name"
        done
        ;;
    restart)
        for entry in "${PROCESSES[@]}"; do
            IFS=':' read -r name path <<< "$entry"
            restart_process "$name" "$path"
        done
        ;;
    status)
        for entry in "${PROCESSES[@]}"; do
            IFS=':' read -r name path <<< "$entry"
            if check_process "$name"; then
                log "$name is running (PID: $(cat "$PID_DIR/$name.pid"))"
            else
                log "$name is not running"
            fi
        done
        ;;
    monitor)
        monitor
        ;;
    *)
        echo "Usage: $0 {start|stop|restart|status|monitor}"
        exit 1
        ;;
esac

exit 0