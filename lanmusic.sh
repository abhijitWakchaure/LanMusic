#!/bin/bash

print_help() {
  echo -e $'\e[32;1m\tSyntax:\n\t\tlanmusic \e[0m[parameters] '
  echo -e $'\e[32;1m\n\tAvailable Parameters:'
  print_help_string "-mr | --music-root" "The root path for your music directory"
  print_help_string "-h | --help" "Prints this message"
}

print_help_string() {
  echo -e $'\e[32;1m\t\t'$1'\e[0m'
  echo -e '\t\t\t'$2
}
launch_app() {
  docker run -it -v $1:/Music -p 80:80 -p 8080:8080 -p 9000:9000 lanmusic
}
if [[ $# -gt 0 ]]; then
  key="$1"
  case $key in
    -mr|--music-root)
      launch_app $2
    ;;
    *)
    print_help
    shift # past argument
    ;;
  esac
else
  print_help
fi