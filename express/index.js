
const express = require('express')
const app = express()
const port = process.env.PORT


function sleep(ms) {
  return new Promise((resolve) => {
    setTimeout(resolve, ms);
  });
}

app.get('/', async (req, res) => {
  // await sleep(process.env.TIMEOUT)
  res.send({hello:"world"})
})

app.listen(port, () => {
  console.log(`Example app listening at http://localhost:${port}`)
})