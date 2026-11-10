#!/bin/sh
# ponytail: freedesktop bell + sway workspace 1. Change the two commands if either moves.
setsid sh -c 'ffmpeg -v error -stream_loop 3 -i /usr/share/sounds/freedesktop/stereo/bell.oga -af volume=2.5 -f wav - | paplay --volume=65536' >/dev/null 2>&1 &
swaymsg workspace number 1 >/dev/null
pos=$(swaymsg -t get_outputs | python3 -c 'import json,sys
for o in json.load(sys.stdin):
    if o.get("focused"):
        r = o["rect"]
        print(r["x"] + r["width"] // 2, r["y"] + r["height"] // 2)
        break')
swaymsg seat seat0 cursor set $pos >/dev/null
