import http from "k6/http";
import { sleep } from "k6";

export default function () {
  http.get("https://geocache-stage.dev.clarkezone.dev/");
  sleep(1);
}
