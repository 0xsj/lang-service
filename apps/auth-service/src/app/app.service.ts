import { Injectable } from '@nestjs/common';

@Injectable()
export class AppService {
  private readonly ping = { status: 'auth-service - ok' };
  getData(): { message: string } {
    return { message: 'Hello API' };
  }

  checkHealth() {
    return this.ping;
  }
}
