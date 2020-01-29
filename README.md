# LanMusic

## Run using docker commands
Simply create a volume mount to your music directory and you are done!

```docker run -it -v path/to/your/music/directory/:/Music abhijitwakchaure/lanmusic```

## Run using shell script
Download the lanmusic shell script from [here](https://raw.githubusercontent.com/abhijitWakchaure/LanMusic/master/lanmusic.sh) and pass the required argument `-mr` or `--music-root` like this

```./lanmusic.sh -mr ~/Music```


## Contributors
* Abhijit Wakchaure