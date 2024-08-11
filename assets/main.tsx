import "./css/main.css";

const evtSource = new EventSource("/sse");

evtSource.onmessage = (e) => {
  console.log("sse message", e);
};

console.log("hello");
