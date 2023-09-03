/**
 * This is not a production server yet!
 * This is only a minimal backend to get started.
 */

import express from "express";
import { serve, setup } from "swagger-ui-express";
import { exampleRouter } from "./routes/example/example";
import { specs } from "./config/swagger";

const app = express();

app.use("/api", exampleRouter);
app.get("/swagger.json", (req, res) => {
  res.setHeader("Content-Type", "application/json");
  res.send(specs);
});
app.use("/api-docs", serve, setup(specs));

const port = process.env.PORT || 3333;
const server = app.listen(port, () => {
  console.log(`Listening at http://localhost:${port}/api`);
});
server.on("error", console.error);
