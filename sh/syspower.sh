export DDE_DEBUG_LEVEL=debug
export DDE_DEBUG_MATCH=daemon/system/power
sudo pkill -ef dde-system-daemon;sudo /usr/lib/deepin-daemon/dde-system-daemon
