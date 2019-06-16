#!/bin/sh

SESSION="gogame"

tmux kill-session -t $SESSION || true

tmux new-session -d -s $SESSION
tmux split-window -d -t 0 -h

tmux send-keys -t 0 "go run main.go --config config.yaml start" enter
tmux send-keys -t 1 "cd ../gogame-frontend && live-server" enter

tmux attach -t $SESSION
