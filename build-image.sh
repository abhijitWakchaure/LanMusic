#!/bin/bash
if [ "$1" == "ui" ]; then
  cd ./UI && npm run build && cd ..
fi
docker build -t lanmusic .