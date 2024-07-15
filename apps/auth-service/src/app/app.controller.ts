import { Controller, Get, Logger } from '@nestjs/common';

import { AppService } from './app.service';
import { MessagePattern } from '@nestjs/microservices';

const logger = new Logger();

@Controller()
export class AppController {
  constructor(private readonly appService: AppService) {}

  @Get()
  getData() {
    return this.appService.getData();
  }

  @MessagePattern({ cmd: 'auth-service-ping' })
  async checkHealth() {
    const result = await this.appService.checkHealth();
    logger.log('Received ping from auth service:', result);
    return result;
  }
}
