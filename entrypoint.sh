#!/bin/sh
# Start the exporter in the background
/opt/gluetun-exporter &

# Capture the PID of the exporter
EXPORTER_PID=$!

# Run the gluetun entrypoint as a foreground process
/gluetun-entrypoint

# Wait for the exporter to exit and then return the exit status
wait $EXPORTER_PID
EXIT_CODE=$?

# If exporter fails, exit with the same code
if [ $EXIT_CODE -ne 0 ]; then
  echo "Exporter exited with error code $EXIT_CODE, stopping the container."
  exit $EXIT_CODE
fi
