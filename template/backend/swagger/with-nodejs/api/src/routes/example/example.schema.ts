import { generateSchema } from '@anatine/zod-openapi';
import z from 'zod';

export const exampleCreateSchema = z.object({ message: z.string() });
export const exampleSchema = exampleCreateSchema;
// Naming convention <resource>OpenApiDefinitions
export const exampleOpenApiDefinitions = {
  Example: generateSchema(exampleSchema),
  ExampleCreate: generateSchema(exampleCreateSchema),
};
