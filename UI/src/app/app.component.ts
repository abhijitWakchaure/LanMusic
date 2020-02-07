import { Component, OnInit } from "@angular/core";
import { DataServiceService } from "src/services/data-service.service";
import { IMList } from "./models/iMList";
import { ISongMetadata } from "./models/iSongMetadata";
import { PlayerService } from "./services/player.service";

@Component({
  selector: "app-root",
  templateUrl: "./app.component.html",
  styleUrls: ["./app.component.css"]
})
export class AppComponent implements OnInit {
  title = "lanMusic";
  isDataLoaded: boolean;
  isError: boolean;
  ErrorMsg: string;
  musicList: IMList;

  constructor(
    private dataService: DataServiceService,
    private ps: PlayerService
  ) {}

  ngOnInit() {
    this.isDataLoaded = false;
    this.isError = false;
    this.getMusic();
  }

  getMusic() {
    this.dataService.doGetMusic().subscribe((res: any) => {
      this.isDataLoaded = !this.isDataLoaded;
      if (res.status == "success") {
        this.musicList = res.data;
      } else {
        this.isError = true;
        this.ErrorMsg = res.message;
        console.log("Unable to getMusic", res);
      }
    });
  }

  getNext() {
    this.isDataLoaded = !this.isDataLoaded;
    this.dataService
      .doGetNext(this.musicList.cursor.index)
      .subscribe((res: any) => {
        this.isDataLoaded = !this.isDataLoaded;
        if (res.status == "success") {
          this.musicList = res.data;
        } else {
          this.isError = true;
          this.ErrorMsg = res.message;
          console.log("Unable to getMusic", res);
        }
      });
  }

  getPrevious() {
    this.isDataLoaded = !this.isDataLoaded;
    this.dataService
      .doGetPrevious(this.musicList.cursor.index)
      .subscribe((res: any) => {
        this.isDataLoaded = !this.isDataLoaded;
        if (res.status == "success") {
          this.musicList = res.data;
        } else {
          this.isError = true;
          this.ErrorMsg = res.message;
          console.log("Unable to getMusic", res);
        }
      });
  }

  updateNowPlaying(md: ISongMetadata) {
    this.ps.nowPlaying = md;
    this.ps.observableNowPlaying.next(md)
  }
}
