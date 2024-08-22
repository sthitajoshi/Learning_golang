const express = require('express');
const app = express();
const port = 3000;

// Middleware to parse JSON and form data
app.use(express.json());
app.use(express.urlencoded({ extended: true }));

// Route for GET request
app.get('/', (req, res) => {
    res.send('<h1>GET request received</h1>');
});

// Route for POST request with JSON
app.post('/post', (req, res) => {
    console.log('Received JSON:', req.body);
    res.json({
        message: 'POST request with JSON received',
        data: req.body
    });
});

// Route for POST request with Form data
app.post('/postform', (req, res) => {
    console.log('Received Form Data:', req.body);
    res.send(`<h1>POST request with Form data received</h1><p>Data: ${JSON.stringify(req.body)}</p>`);
});

// Start the server
app.listen(port, () => {
    console.log(`Server is running on http://localhost:${port}`);
});
