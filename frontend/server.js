const http = require('http');
const fs = require('fs');
const path = require('path');
const PORT = process.env.PORT || 3000;
const IP_BIND = process.env.IP || '0.0.0.0';

http.createServer((request, response) => {
  let filePath = '.' + request.url;

  if (filePath[filePath.length - 1] === '/') filePath += 'index.html';

  let extname = path.extname(filePath);

  let fullpath = path.join(__dirname, filePath);

  let contentType = 'text/html';
  switch (extname) {
    case '.js':
      contentType = 'text/javascript';
      break;
    case '.css':
      contentType = 'text/css';
      break;
    case '.json':
      contentType = 'application/json';
      break;
    case '.png':
      contentType = 'image/png';
      break;
    case '.jpg':
      contentType = 'image/jpg';
      break;
    case '.wav':
      contentType = 'audio/wav';
      break;
  }

  fs.readFile(fullpath, function (error, content) {
    if (error) {
      if (filePath.includes('favicon.ico')) {
        response.writeHead(200, {
          'Content-Type': 'image/png'
        });
        response.end('', 'utf-8');
        return
      }


      response.writeHead(404);

      error.statusCode = response.statusCode;
      error.statusMessage = response.statusMessage;

      let errorMessage =
        `Error: ${error.statusMessage}\nStatus: ${error.statusCode}\nPath: ${error.path}`;

      console.error(errorMessage);
      response.end(errorMessage);
    } else {
      response.setHeader('Access-Control-Allow-Origin', '*');
      response.setHeader('Access-Control-Allow-Methods', 'GET, POST, OPTIONS, PUT, PATCH, DELETE');
      response.setHeader('Access-Control-Allow-Headers', 'X-Requested-With,content-type');
      response.setHeader('Access-Control-Allow-Credentials', true);
      response.writeHead(200, {
        'Content-Type': contentType
      });
      response.end(content, 'utf-8');
      console.log(`${request.method} ${request.url}`);
    }
  });

}).listen(PORT, IP_BIND, (err) => {
  if (err) {
    console.log(err);
    return;
  }

  let appUrl = `http://${IP_BIND}:${PORT}`;
  console.log('Server running on', appUrl);
});
