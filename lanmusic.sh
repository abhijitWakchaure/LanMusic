#!/bin/bash

IMAGE_NAME="lanmusic"
TAG=":latest"
CONTAINER_NAME="lanmusic"
DOCKER_REPO="abhijitwakchaure/"
MUSIC_ROOT=""
RUN_MODE=""
print_help() {
  echo -e $'\e[32;1m\tSyntax:\n\t\tlanmusic \e[0m[-local | -remote] [-mr | -music-root] <yourMusicRootDir> '
  echo -e $'\e[32;1m\n\tAvailable Parameters:'
  print_help_string "-mr | -music-root | --music-root" "The root path for your music directory"
  print_help_string "-l  | -local      | --local" "Local Mode: Run app from your local docker repo"
  print_help_string "-r  | -remote     | --remote" "Remote Mode: Run app via pulling the latest docker image"
  print_help_string "-h  | -help       | --help" "Prints this message"
  echo -e $'\e[32;1m\n\tExample:'
  echo -e $'\e[0m\t\t./lanmusic.sh -local -music-root ~/Music\n'
}

print_help_string() {
  echo -e $'\e[32;1m\t\t'$1'\e[0m'
  echo -e '\t\t\t'$2
}
set_music_root() {
  echo "Setting music root to: $1"
  MUSIC_ROOT=$1
}
launch_app_remote() {
  RUN_MODE="remote"
  echo "Setting launch mode to: $RUN_MODE"
  docker pull $DOCKER_REPO$IMAGE_NAME$TAG
}
launch_app_local() {
  RUN_MODE="local"
  echo "Setting launch mode to: $RUN_MODE"
  DOCKER_REPO=""
}
if [[ $# -eq 0 ]]; then
  print_help
  exit 1
else
  while [ "$1" != "" ]; do
    case $1 in
    -mr | -music-root | --music-root)
      set_music_root $2
      shift
      shift
      ;;
    -l | -local | --local)
      launch_app_local
      shift
      ;;
    -r | -remote | --remote)
      launch_app_remote
      shift
      ;;
    *)
      echo "Unrecognized option!"
      print_help
      exit 1
      ;;
    esac
  done
fi

if [ "$MUSIC_ROOT" == "" ] || [ "$RUN_MODE" == "" ]; then
  echo "Please provide Music root and mode. Please check help for more details"
  exit 1
else
  docker rm $CONTAINER_NAME >/dev/null 2>&1
  docker run -it \
    -v $MUSIC_ROOT:/Music \
    -p 80:80 \
    -p 9000:9000 \
    --name $CONTAINER_NAME \
    $DOCKER_REPO$IMAGE_NAME$TAG
fi