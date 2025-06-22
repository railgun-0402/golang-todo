import http from "k6/http";
import { sleep } from "k6";

export const options = {
  // 100ユーザー/300秒を実行する
  vus: 100,
  duration: "300s",
};

export default function () {
  http.get("http://localhost:8080");
  sleep(1);
}
