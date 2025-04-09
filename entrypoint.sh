#!/bin/sh

# Start the first process
/opt/gluetun-exporter &
EXPORTER_PID=$!

# Start the second process
/gluetun-entrypoint &
GLUETUN_PID=$!

# Function to clean up and exit
cleanup_and_exit() {
    echo "Exiting..."
    kill $EXPORTER_PID $GLUETUN_PID 2>/dev/null
    exit 1
}

# Monitor the processes
while true; do
    if ! kill -0 $EXPORTER_PID 2>/dev/null; then
        echo "gluetun-exporter stopped!"
        cleanup_and_exit
    fi

    if ! kill -0 $GLUETUN_PID 2>/dev/null; then
        echo "gluetun stopped!"
        cleanup_and_exit
    fi

    sleep 1
done
