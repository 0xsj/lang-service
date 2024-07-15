import { Module } from '@nestjs/common';

import { AppController } from './app.controller';
import { AppService } from './app.service';
import { ClientsModule, Transport } from '@nestjs/microservices';

@Module({
  imports: [
    ClientsModule.register([
      {
        name: 'AUTH_SERVICE',
        transport: Transport.TCP,
        options: {
          host: '127.0.0.1',
          port: parseInt(process.env.AUTH_SERVICE_PORT),
        },
      },
      {
        name: 'BRAND_SERVICE',
        transport: Transport.TCP,
        options: {
          host: '127.0.0.1',
          port: parseInt(process.env.BRAND_SERVICE_PORT),
        },
      },
      {
        name: 'DOCUMENT_SERVICE',
        transport: Transport.TCP,
        options: {
          host: '127.0.0.1',
          port: parseInt(process.env.DOCUMENT_SERVICE_PORT),
        },
      },
      {
        name: 'ENTITY_SERVICE',
        transport: Transport.TCP,
        options: {
          host: '127.0.0.1',
          port: parseInt(process.env.ENTITY_SERVICE_PORT),
        },
      },
      {
        name: 'USER_SERVICE',
        transport: Transport.TCP,
        options: {
          host: '127.0.0.1',
          port: parseInt(process.env.USER_SERVICE_PORT),
        },
      },
    ]),
  ],
  controllers: [AppController],
  providers: [AppService],
})
export class AppModule {}
