import { object, string, number } from 'zod';
import { mongoEnvSchema } from './mongo.env';
const envSchema = object({
  NODE_ENV: string(),
  PORT: number().min(0).max(65535).default(3333),
  ...mongoEnvSchema,
});

export const getEnvVariables = () => {
  try {
    return envSchema.parse(process.env);
  } catch (error) {
    console.error('Error validating environment variables:', error.message);
    process.exit(1);
  }
};
