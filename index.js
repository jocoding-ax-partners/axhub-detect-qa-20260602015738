const http = require('http');
const port = process.env.PORT || 8080;
http.createServer((_req, res) => { res.writeHead(200, {'content-type':'text/plain'}); res.end('axhub remix-compatible ok'); }).listen(port, '0.0.0.0');
