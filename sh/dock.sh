export DDE_DEBUG_LEVEL=debug
export DDE_DEBUG_MATCH=daemon/dock
pkill -ef dde-session-daemon;/usr/lib/deepin-daemon/dde-session-daemon
