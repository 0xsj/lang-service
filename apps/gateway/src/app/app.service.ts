import { Inject, Injectable } from '@nestjs/common';
import { ClientProxy } from '@nestjs/microservices';

@Injectable()
export class AppService {
  constructor(
    @Inject('AUTH_SERVICE') private readonly authClient: ClientProxy,
    @Inject('BRAND_SERVICE') private readonly brandClient: ClientProxy,
    @Inject('DOCUMENT_SERVICE') private readonly documentClient: ClientProxy,
    @Inject('ENTITY_SERVICE') private readonly entityClient: ClientProxy,
    @Inject('USER_SERVICE') private readonly userClient: ClientProxy
  ) {}

  getData(): { message: string } {
    return { message: 'Hello API' };
  }

  checkHealth() {
    return this.authClient.send({ cmd: 'auth-service-ping' }, {});
  }

  checkBrandHealth() {
    return this.brandClient.send({ cmd: 'brand-service-ping' }, {});
  }

  checkDocumentHealth() {
    return this.documentClient.send({ cmd: 'document-service-ping' }, {});
  }

  checkEntityHealth() {
    return this.entityClient.send({ cmd: 'entity-service-ping' }, {});
  }

  checkUserHealth() {
    return this.userClient.send({ cmd: 'user-service-ping' }, {});
  }
}
