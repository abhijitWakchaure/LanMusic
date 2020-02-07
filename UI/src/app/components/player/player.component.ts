import { Component, OnInit, Input } from "@angular/core";
import { ISongMetadata } from "src/app/models/iSongMetadata";
import { PlayerService } from "src/app/services/player.service";

@Component({
  selector: "app-player",
  templateUrl: "./player.component.html",
  styleUrls: ["./player.component.css"]
})
export class PlayerComponent implements OnInit {
  constructor(public ps: PlayerService) {}
  ngOnInit() {}
}
