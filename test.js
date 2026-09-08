import http from "k6/http";
import { sleep } from "k6";

export const options = {
  stages: [
    { duration: "5s", target: 1 },
    { duration: "5s", target: 2 },
    { duration: "5s", target: 0 },
  ],
};

export default function () {
  http.get("https://quickpizza.grafana.com");
  sleep(0.5);
}
