const express = require('express');
const app = express();
const message = process.env.APPHUB_QA_NODE_MESSAGE || 'node ok';
app.get('/', (_req, res) => res.json({ language: 'node', message }));
const port = process.env.PORT || 3000;
app.listen(port, () => console.log(`node listening on ${port}`));
