const express = require('express');
const path = require('path');
const bodyParser = require('body-parser');
const axios = require('axios');

const app = express();

app.use(express.json());

const port = 8080;
const servicePort = process.env.SERVICE_PORT ?? 12345;

const serviceUrl = `http://localhost:${servicePort}`;

/**
The following routes forward requests (using pipe) from our React client to our dapr-enabled services. Our Dapr sidecar lives on localhost:<daprPort>. We invoke other Dapr enabled services by calling /v1.0/invoke/<DAPR_ID>/method/<SERVICE'S_ROUTE>.
*/

app.post('/calculate/add', async (req, res) => {
  try {
      // Invoke Dapr add app
      const appResponse = await axios.post(`${serviceUrl}/addapp/method/add`, req.body);

      // Return expected string result to client
      return res.send(`${appResponse.data}`); 
  } catch (err) {
      console.log(err);
  }
});

app.post('/calculate/subtract', async (req, res) => {
  try {
      // Invoke Dapr subtract app
      console.log("subtract app** 1")
      const appResponse = await axios.post(`${serviceUrl}/subtractapp/method/subtract`, req.body);
      console.log("subtract app** 2")

      // Return expected string result to client
      return res.send(`${appResponse.data}`); 
  } catch (err) {
      console.log(err);
  }
});

app.post('/calculate/multiply', async (req, res) => {
  try {
      // Dapr invoke multiply app
      const appResponse = await axios.post(`${serviceUrl}/multiplyapp/method/multiply`, req.body);

      // Return expected string result to client
      return res.send(`${appResponse.data}`); 
  } catch (err) {
      console.log(err);
  }
});

app.post('/calculate/divide', async (req, res) => {
  try {
      // Dapr invoke divide app
      const appResponse = await axios.post(`${serviceUrl}/divideapp/method/divide`, req.body);

      // Return expected string result to client
      return res.send(`${appResponse.data}`); 
  } catch (err) {
      console.log(err);
  }
});

var state = {};

// Store state in process
app.get('/state', async (req, res) => {
  res.send(JSON.stringify(state))
});

// Store state in process
app.post('/persist', async (req, res) => {
  state = req.body;
  res.send(JSON.stringify(state))
});

// Serve static files
app.use(express.static(path.join(__dirname, 'client/build')));

// For default home request route to React client
app.get('/', async function (_req, res) {
  try {
    return await res.sendFile(path.join(__dirname, 'client/build', 'index.html'));
  } catch (err) {
    console.log(err);
  }
});

app.listen(process.env.PORT || port, () => console.log(`Listening on port ${port}!`));
