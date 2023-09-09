import mongoose, { Schema, Document } from 'mongoose';

// Define an interface for the model
export interface HelloWorldType {
  message: string;
}
type IHelloWorld = HelloWorldType & Document;

// Define the schema for the "HelloWorld" model
const helloWorldSchema = new Schema({
  message: { type: String, required: true },
});
// Create the "HelloWorld" model
export const HelloWorld = mongoose.model<IHelloWorld>(
  'HelloWorld',
  helloWorldSchema
);
