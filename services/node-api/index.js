const http = require("http");

const PORT = Number(process.env.PORT) || 8080;

function json(res, status, body) {
  res.writeHead(status, { "Content-Type": "application/json" });
  res.end(JSON.stringify(body));
}

const server = http.createServer((req, res) => {
  if (req.method === "GET" && req.url === "/healthz/live") {
    json(res, 200, { status: "live" });
    return;
  }
  if (req.method === "GET" && req.url === "/healthz/ready") {
    json(res, 200, { status: "ready" });
    return;
  }

  json(res, 404, { error: "not found" });
});

server.listen(PORT, () => {
  console.log(`listening on :${PORT}`);
});
