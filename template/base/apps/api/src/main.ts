/**
 * This is not a production server yet!
 * This is only a minimal backend to get started.
 */

import express from 'express';
import * as path from 'path';
const swaggerUi = require('swagger-ui-express');
import { exampleRouter } from './routes/example/example';
import { specs } from './config/swagger';

const app = express();

app.use('/api', exampleRouter);
app.use('/api-docs', swaggerUi.serve, swaggerUi.setup(specs));

const port = process.env.PORT || 3333;
const server = app.listen(port, () => {
  console.log(`Listening at http://localhost:${port}/api`);
});
server.on('error', console.error);
