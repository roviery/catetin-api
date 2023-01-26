const dbo = require("../../db/conn");
const bcrypt = require("bcrypt")
const ObjectId = require("mongodb").ObjectId

const register = async (reqBody) => {
  try {
    console.log(reqBody)
    const {fullname, email, password} = reqBody
    const dbConnect = dbo.getDb("users")
    const createdAt = new Date
    const hashedPassword = await bcrypt.hash(password, 10)

    const newUser = {
      fullname: fullname,
      email: email,
      password: hashedPassword,
      createdAt: createdAt.toUTCString()
    }
    dbConnect.collection("users").insertOne(newUser, (err, qres) => {
      if (err) return err
      return qres
    })
  }catch (error) {
    return error
  }
}

module.exports = {register}