#!/bin/sh
setsid sh -c 'ffmpeg -v error -stream_loop 3 -i /usr/share/sounds/freedesktop/stereo/bell.oga -af volume=2.5 -f wav - | paplay --volume=65536' >/dev/null 2>&1 &
