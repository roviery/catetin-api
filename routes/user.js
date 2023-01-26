const express = require("express")
const userRoutes = express.Router()
const {register} = require("../users/controller/user")

userRoutes.post("/register", register)

module.exports = userRoutes