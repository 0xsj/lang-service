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

  @Get('brand-service')
  async checkBrandHealth() {
    const result = await this.appService.checkBrandHealth();
    logger.log('gateway controller - brand service healthcheck');
    return result;
  }
  @Get('document-service')
  async checkDocumentHealth() {
    const result = await this.appService.checkDocumentHealth();
    logger.log('gateway controller - document service healthcheck');
    return result;
  }
  @Get('entity-service')
  async checkEntityHealth() {
    const result = await this.appService.checkEntityHealth();
    logger.log('gateway controller - entity service healthcheck');
    return result;
  }
  @Get('user-service')
  async checkUserHealth() {
    const result = await this.appService.checkUserHealth();
    logger.log('gateway controller - user service healthcheck');
    return result;
  }
}
