import { string } from 'zod';

export const mongoEnvSchema = {
  NOSQL_URL: string(),
};
