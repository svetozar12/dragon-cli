import { HelloWorld, connectToMongo } from './database/mongo';
import express from 'express';
import { serve, setup } from 'swagger-ui-express';
import { exampleRouter } from './routes/example';
import { specs } from './utils/swagger';
import { getEnvVariables } from './utils/env';

export async function bootstrap() {
  const { PORT } = getEnvVariables();
  await connectToMongo();
  await HelloWorld.create({ message: 'Hi' });
  const app = express();

  app.use('/api', exampleRouter);
  app.get('/swagger.json', (req, res) => {
    res.setHeader('Content-Type', 'application/json');
    res.send(specs);
  });
  app.use('/api-docs', serve, setup(specs));

  const port = PORT || 3333;
  const server = app.listen(port, () => {
    console.log(`Listening at http://localhost:${port}/api`);
  });
  server.on('error', console.error);
}
