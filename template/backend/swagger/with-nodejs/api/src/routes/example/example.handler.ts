import { Request, Response } from 'express';
import { HelloWorld } from '../../database/mongo';
import { exampleSchema } from './example.schema';

export async function getExample(req: Request, res: Response) {
  const data = await HelloWorld.findOne();
  return res.json(data);
}

export async function createExample(req: Request, res: Response) {
  const payload = exampleSchema.parse(req.body);
  const example = await HelloWorld.create(payload);
  return res.json(example);
}
