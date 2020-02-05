import { Injectable } from "@angular/core";
import { ISongMetadata } from "../models/iSongMetadata";
import { BehaviorSubject } from "rxjs";

@Injectable({
  providedIn: "root"
})
export class PlayerService {
  public nowPlaying: ISongMetadata;
  public observableNowPlaying;
  public songAudio: HTMLAudioElement;
  public songAudioProgress: number;
  public songAudioCurrentTimeDate: Date;
  public songAudioDurationDate: Date;

  constructor() {
    this.observableNowPlaying = new BehaviorSubject<ISongMetadata>(
      this.nowPlaying
    );
    this.observableNowPlaying
      .asObservable()
      .subscribe((md: ISongMetadata) => this.playSong(md));
  }

  playSong(md: ISongMetadata) {
    if (md) {
      if (this.songAudio) {
        this.songAudio.pause();
      }
      let url = "http://localhost:9000/stream/" + md.id;
      this.songAudio = new Audio(url);
      this.songAudio.play();
      if (!isNaN(this.songAudio.duration) && isFinite(this.songAudio.duration))
        this.songAudioDurationDate = this.toDateTime(this.songAudio.duration);
      else this.songAudioDurationDate = this.toDateTime(0);
      this.songAudioCurrentTimeDate = this.toDateTime(0);
      this.startProgressBar();
    }
  }
  startProgressBar() {
    this.songAudioProgress = 0;
    let updateProgressBar = setInterval(() => {
      if (this.songAudio.ended) {
        clearInterval(updateProgressBar);
      }
      this.songAudioCurrentTimeDate = this.toDateTime(
        Math.floor(this.songAudio.currentTime)
      );
      this.songAudioProgress = Math.floor(
        (Math.floor(this.songAudio.currentTime) /
          Math.floor(this.songAudio.duration)) *
          100
      );
    }, 1000);
  }

  togglePlayPause() {
    if (this.songAudio.paused) this.songAudio.play();
    else this.songAudio.pause();
  }

  toDateTime(secs) {
    var t = new Date(1970, 0, 1, 0, 0, 0, 0); // Epoch
    t.setSeconds(secs);
    return t;
  }
}
