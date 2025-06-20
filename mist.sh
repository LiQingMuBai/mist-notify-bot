#!/bin/bash
# Dedicated monitor for /seven/agent/mist-agent

### Configuration ###
PROCESS_NAME="mist-agent"
PROCESS_CMD="/seven/agent/mist-agent"    # Full path to your binary
PROCESS_ARGS=""                               # Add any required arguments here
CHECK_INTERVAL=30                             # Check every 30 seconds
LOG_DIR="/seven/agent/logs"              # Dedicated log directory
LOG_FILE="${LOG_DIR}/monitor.log"             # Monitor log file
PROCESS_LOG="${LOG_DIR}/mist-agent.log"       # Process output log
MAX_RESTARTS=6                                # Max restarts per hour
LOCK_FILE="/tmp/mist-agent-monitor.lock"

### Initial Setup ###
mkdir -p "${LOG_DIR}"
touch "${LOG_FILE}" "${PROCESS_LOG}"
chmod 600 "${LOG_FILE}" "${PROCESS_LOG}"
exec > >(tee -a "${LOG_FILE}") 2>&1

# Lock file mechanism
if [ -e "${LOCK_FILE}" ] && kill -0 "$(cat "${LOCK_FILE}")" 2>/dev/null; then
    echo "[$(date '+%Y-%m-%d %T')] ERROR: Monitor already running (PID: $(cat "${LOCK_FILE}"))" | tee -a "${LOG_FILE}"
    exit 1
fi
echo $$ > "${LOCK_FILE}"
trap 'rm -f "${LOCK_FILE}"; exit' EXIT INT TERM

### Functions ###
log() {
    echo "[$(date '+%Y-%m-%d %T')] $1"
}

is_running() {
    pgrep -f "${PROCESS_CMD}" >/dev/null  # Using -f to match full command path
}

start_agent() {
    # Rotate logs if process log gets too large (>10MB)
    if [ -f "${PROCESS_LOG}" ] && [ $(stat -c%s "${PROCESS_LOG}") -gt 10485760 ]; then
        mv "${PROCESS_LOG}" "${PROCESS_LOG}.old"
    fi

    log "Launching: ${PROCESS_CMD} ${PROCESS_ARGS}"
    nohup "${PROCESS_CMD}" ${PROCESS_ARGS} >> "${PROCESS_LOG}" 2>&1 &
    sleep 2  # Give it time to start

    if is_running; then
        log "Started successfully (PID: $(pgrep -f "${PROCESS_CMD}"))"
        return 0
    else
        log "FAILED to start process"
        return 1
    fi
}

### Main Monitor ###
log "==== Starting mist-agent monitor ===="
log "Process: ${PROCESS_CMD}"
log "Args: ${PROCESS_ARGS:-None}"
log "Check interval: ${CHECK_INTERVAL}s"
log "Max restarts: ${MAX_RESTARTS}/hour"
log "Process output: ${PROCESS_LOG}"
log "Monitor log: ${LOG_FILE}"

RESTART_COUNT=0
LAST_RESTART=0

while true; do
    if is_running; then
        log "Process running (PID: $(pgrep -f "${PROCESS_CMD}"))"
        RESTART_COUNT=0  # Reset counter if running normally
    else
        CURRENT_TIME=$(date +%s)

        # Reset counter if last restart was >1 hour ago
        if [ $((CURRENT_TIME - LAST_RESTART)) -gt 3600 ]; then
            RESTART_COUNT=0
        fi

        if [ ${RESTART_COUNT} -ge ${MAX_RESTARTS} ]; then
            log "EMERGENCY: Max restarts reached (${MAX_RESTARTS}/hour) - waiting..."
        else
            log "Process not running - attempting restart..."
            if start_agent; then
                LAST_RESTART=${CURRENT_TIME}
                ((RESTART_COUNT++))
            fi
        fi
    fi
    sleep ${CHECK_INTERVAL}
done