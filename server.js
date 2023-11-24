const express = require('express');
const bodyParser = require('body-parser');

const app = express();
const port = 4000;

// Middleware to parse JSON
app.use(bodyParser.json());

// Endpoint 1: REST API to return probability is 50% when catching Pokemon
app.get('/catch-pokemon', (req, res) => {
    const probability = Math.random() < 0.5 ? 1 : 0; // 50% probability
    res.json({ "result": probability });
});

// Endpoint 2: REST API to release Pokemon
var number = 0
app.get('/release-pokemon', (req, res) => {
    number++;
    console.log(number)

    if (isPrime(number)) {
        res.json({ result: true, message: "release success" });
    } else {
        res.status(400).json({ error: 'Release failed. The provided number is not prime.' });
    }
});

function isPrime(num) {
    if (num <= 1) return false;
    for (let i = 2; i <= Math.sqrt(num); i++) {
        if (num % i === 0) return false;
    }
    return true;
}

// Start the server
app.listen(port, () => {
    console.log(`Server is running on port ${port}`);
});
