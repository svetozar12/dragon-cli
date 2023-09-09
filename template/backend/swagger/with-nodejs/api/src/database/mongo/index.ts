import { getEnvVariables } from '../../utils/env';
import mongoose from 'mongoose';

let connection: typeof mongoose;

async function connectToMongo() {
  try {
    if (!connection) {
      const { NOSQL_URL } = getEnvVariables();
      connection = await mongoose.connect(NOSQL_URL);
      console.log('Connected to MongoDB');
    }
  } catch (error) {
    console.error('MongoDB connection error:', error);
    // Optionally, you can throw the error if needed for higher-level handling
    throw error;
  }
}

async function disconnectFromMongo() {
  connection.disconnect();
}

export {
  connectToMongo,
  disconnectFromMongo as disconnectFromMongo,
  connection,
};

export * from './models/HelloWorld.model';
