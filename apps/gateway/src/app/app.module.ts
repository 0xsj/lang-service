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
          // host: '0.0.0.0',
          host: '::',
          port: 3001,
        },
      },
      {
        name: 'BRAND_SERVICE',
        transport: Transport.TCP,
        options: {
          // host: '0.0.0.0',
          host: '::',
          port: 3002,
        },
      },
      {
        name: 'DOCUMENT_SERVICE',
        transport: Transport.TCP,
        options: {
          // host: '0.0.0.0',
          host: '::',
          port: 3003,
        },
      },
      {
        name: 'ENTITY_SERVICE',
        transport: Transport.TCP,
        options: {
          // host: '0.0.0.0',
          host: '::',
          port: 3004,
        },
      },
      {
        name: 'USER_SERVICE',
        transport: Transport.TCP,
        options: {
          // host: '0.0.0.0',
          host: '::',
          port: 3005,
        },
      },
    ]),
  ],
  controllers: [AppController],
  providers: [AppService],
})
export class AppModule {}
