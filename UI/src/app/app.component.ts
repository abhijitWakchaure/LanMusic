import { Component, OnInit } from '@angular/core';
import { DataServiceService } from 'src/services/data-service.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit {
  title = 'lanMusic';
  isDataLoaded: boolean
  isError: boolean
  ErrorMsg: string
  musicData = []


  constructor(private dataService: DataServiceService){}

  ngOnInit() {
    this.isDataLoaded = false
    this.isError = false
    this.getMusic()
  }

  getMusic() {
    this.dataService.doGetMusic().subscribe((res: any) => {
      this.isDataLoaded = !this.isDataLoaded
      if (res.status == 'success') {
        this.musicData = res.data
      } else {
        this.isError = true
        this.ErrorMsg = res.message
        console.log("Unable to getMusic", res)
      }
    })
  }
}
