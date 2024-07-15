import { Controller, Get, Logger, Post } from '@nestjs/common';

import { AppService } from './app.service';

const logger = new Logger();

@Controller()
export class AppController {
  constructor(private readonly appService: AppService) {}

  @Get()
  getData() {
    return this.appService.getData();
  }

  @Get('auth-service')
  async checkHealth() {
    const result = await this.appService.checkHealth();
    logger.log('gateway controller - auth service healthcheck');
    return result;
  }
}
