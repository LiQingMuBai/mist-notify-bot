# mist-bot

## Telegram Notifications

The system sends real-time Telegram notifications for user deposits and energy exchanges.

### Features
- **Deposit Alerts**: Users receive confirmation when deposits are successful.
- **Energy Exchange Updates**: Notifications for energy-to-currency conversions.
- **Transaction Details**: Includes:
    - Amount deposited/exchanged
    - Timestamp
    - Updated balance
- **Error Handling**: Failed notifications are logged for debugging.

### Configuration
1. Add your Telegram Bot token to `.env`:
   ```env
   TELEGRAM_BOT_TOKEN=your_bot_token_here



# Make script executable
chmod +x mist-process-manager.sh

# Start all processes
./mist-process-manager.sh start

# Stop all processes
./mist-process-manager.sh stop

# Restart all processes
./mist-process-manager.sh restart

# Check status
./mist-process-manager.sh status

# Run in monitoring mode (background)
./mist-process-manager.sh monitor &