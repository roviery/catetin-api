const express = require("express")
const app = express()
const cors = require("cors")

require("dotenv").config({path: "./config.env"})
const port = process.env.PORT || 8080
app.use(cors())
app.use(express.json())
app.use('/api/v1/user/', require("./routes/user"))

const dbo = require("./db/conn")

app.listen(port, () => {
  dbo.connectToServer( (err) => {
    if (err) console.error(err)
  })
  console.log(`Server is running on port: ${port}`)
})