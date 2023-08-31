const app = require('express')(),
    PORT = 4000;

app.get('/', (req, res) => {
    res.send(JSON.stringify({Code: 200, Message: "Get Y SO Serious?!"}));
})
app.post('/', (req, res) => {
    const {Code, Message} = req?.body;
    console.log(Code, Message);
    res.send(JSON.stringify({Code: 200, Message: "Post Y SO Serious?!"}));
})
app.listen(PORT, () => {
    console.log(`Running on :${PORT}`);
})
