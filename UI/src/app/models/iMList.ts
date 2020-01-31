import { ICursor } from "./iCursor";
import { ISongMetadata } from "./iSongMetadata";

export interface IMList {
  songs: ISongMetadata[];
  cursor: ICursor;
}
