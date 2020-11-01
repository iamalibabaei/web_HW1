const express = require('express')
const app = express()
const url = require('url')
const fs = require('fs')
const crypto = require('crypto')
const port = 3214

app.use(
  express.urlencoded({
    extended: true
  })
)

app.use(express.json())

app.get('/write/', (req, res) => {
    var q = url.parse(req.url, true).query;
    var txt = q.line;
    fs.readFile('file.txt', function(err, fileData) {
      if (err) throw err;
      var line = fileData.toString().split('\n')[txt - 1]
      //console.log(line)
      res.send(line);
    });
});

app.post('/sha-256', (req, res) => {
  var myInt = req.body.num1 + req.body.num2
  var myHash = crypto.createHash('sha256').update(myInt.toString()).digest('base64');
  var myJson = JSON.stringify({myHash});
  //console.log(myJson)
  res.send(myJson)
})

app.listen(port, () => {
  console.log(`Nodejs backend listening at http://localhost:${port}`)
})
