const express = require('express')
const app = express()
const port = 1699 

app.use(express.json())
app.use(express.urlencoded({extended:true}))

app.get('/',(req,res)=>{
    res.status(200).send("Welcome to goland server")
})

app.post('/post',(req,res)=>{
    let myjson = req.body
    res.status(200).send(myjson)
})

app.post('/postform',(req,res)=>{
    res.status(200).send(JSON.stringify(req.body))
})

app.listen(port,()=>{
    console.log("Server is running on port 1699")
})