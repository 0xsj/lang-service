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

  async testMicroservicesConnections(): Promise<void> {
    try {
      const authResponse = await this.authClient
        .send('auth_ping', {})
        .toPromise();
    } catch (error) {
      console.log(error);
      throw error;
    }
  }
}
