import swaggerJsdoc from 'swagger-jsdoc';

import { extendZodWithOpenApi } from '@anatine/zod-openapi';
import z from 'zod';
import { exampleOpenApiDefinitions } from '../routes/example/example.schema';

extendZodWithOpenApi(z);

const openApiDefinitions = {
  ...exampleOpenApiDefinitions,
};

const options = {
  swaggerDefinition: {
    openapi: '3.0.0',
    info: {
      title: 'Express Swagger API',
      version: '1.0.0',
      description: 'API documentation using Swagger',
    },
  },
  apis: ['./apps/api/src/routes/**/*.ts'], // Path to the API routes containing JSDoc comments
};

export const specs = {
  ...swaggerJsdoc(options),
  components: {
    schemas: openApiDefinitions,
  },
};
console.log(specs);
