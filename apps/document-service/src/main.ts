/**
 * This is not a production server yet!
 * This is only a minimal backend to get started.
 */

import { Logger } from '@nestjs/common';
import { NestFactory } from '@nestjs/core';

import { AppModule } from './app/app.module';
import { Transport } from '@nestjs/microservices';

const logger = new Logger();

async function bootstrap() {
  const app = await NestFactory.createMicroservice(AppModule, {
    transport: Transport.TCP,
    options: {
      // host: '0.0.0.0',
      host: '::',
      port: process.env.DOCUMENT_SERVICE_PORT || 3003,
    },
  });
  logger.log('Document service is listening');
  await app.listen();
}

bootstrap();
