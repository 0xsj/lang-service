import { Injectable } from '@nestjs/common';

@Injectable()
export class AppService {
  private readonly ping = { status: 'brand-service - ok' };
  getData(): { message: string } {
    return { message: 'Hello API' };
  }

  checkHealth() {
    return this.ping;
  }
}
