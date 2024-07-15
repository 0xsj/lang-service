import { Module } from '@nestjs/common';

import { AppController } from './app.controller';
import { AppService } from './app.service';
import { ClientsModule } from '@nestjs/microservices';

@Module({
  imports: [ClientsModule.register([])],
  controllers: [AppController],
  providers: [AppService],
})
export class AppModule {}
