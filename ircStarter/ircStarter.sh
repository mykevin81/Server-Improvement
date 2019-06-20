#!/bin/bash

session="irc"

# Check if the session exists, discarding output
# We can check $? for the exit status (zero for success, non-zero for failure)
tmux has-session -t $session 2>/dev/null

if [ $? != 0 ]; then
    echo "starting irc"
    tmux new -d -s irc 'weechat'
else
    echo "irc has already started"
fi
