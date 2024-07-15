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
      port: process.env.BRAND_SERVICE_PORT || 3002,
    },
  });

  logger.log('Brand service is listening');
  await app.listen();
}

bootstrap();
