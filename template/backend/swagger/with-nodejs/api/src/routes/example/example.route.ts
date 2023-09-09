import { Router } from 'express';
import { createExample, getExample } from './example.handler';

export const exampleRouter = Router();

/**
 * @swagger
 * /api/example:
 *   get:
 *     summary: Returns fist example entity
 *     responses:
 *       200:
 *         description: A successful response
 */
exampleRouter.get('/example', async (req, res) => {
  return getExample(req, res);
});

/**
 * @swagger
 * /api/example:
 *   post:
 *     summary: Create example entity
 *     requestBody:
 *       required: true
 *       content:
 *         application/json:
 *           schema:
 *             $ref: '#/components/schemas/ExampleCreate'
 *     responses:
 *       200:
 *         description: A successful response
 *         content:
 *           application/json:
 *              schema:
 *                $ref: '#/components/schemas/Example'
 */
exampleRouter.post('/example', async (req, res) => {
  return createExample(req, res);
});
