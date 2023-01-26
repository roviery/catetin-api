const {register: registerUsecase} = require("../usecase/user") 

const register = async (req, res) => {
  // const {fullname, email, password} = req.body
  try {
    // const dbConnect = dbo.getDb("users")
    // const createdAt = new Date
    // const hashedPassword = await bcrypt.hash(password, 10)
    // console.log(hashedPassword)

    // const newUser = {
    //   fullname: fullname,
    //   email: email,
    //   password: hashedPassword,
    //   createdAt: createdAt.toUTCString()
    // }
    // dbConnect.collection("users").insertOne(newUser, (err, qres) => {
    //   if (err) res.json({
    //     error: err
    //   })
    //   res.json(qres)
    // })

    const result = registerUsecase(req.body)
    res.json(result)
  } catch (err) {
    res.status(500).json({error: err})
  }
}

module.exports = {register}