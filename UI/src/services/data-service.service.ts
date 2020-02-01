import { Injectable } from "@angular/core";
import {
  HttpHeaders,
  HttpClient,
  HttpErrorResponse
} from "@angular/common/http";
import { Observable, of } from "rxjs";
import { catchError, map } from "rxjs/operators";

const httpOptions = {
  headers: new HttpHeaders({ "Content-Type": "application/json" }),
  observe: "response" as "body"
};
const baseURL = "http://localhost:9000/music";

@Injectable({
  providedIn: "root"
})
export class DataServiceService {
  constructor(private http: HttpClient) {}

  doGetMusic() {
    return this.http.get(baseURL, httpOptions).pipe(
      map((res: any) => {
        console.log("Server Response", res.body);
        return res.body;
      }),
      catchError(this.handleError)
    );
  }

  doGetPrevious(index: number) {
    return this.http.get(baseURL + "/" + (index - 20), httpOptions).pipe(
      map((res: any) => {
        console.log("Server Response", res.body);
        return res.body;
      }),
      catchError(this.handleError)
    );
  }
  doGetNext(index: number) {
    return this.http.get(baseURL + "/" + index, httpOptions).pipe(
      map((res: any) => {
        console.log("Server Response", res.body);
        return res.body;
      }),
      catchError(this.handleError)
    );
  }

  handleError() {
    let res = {
      status: "error",
      message: "Unable to contact server. Please try again after some time."
    };
    return of(res);
  }
}
