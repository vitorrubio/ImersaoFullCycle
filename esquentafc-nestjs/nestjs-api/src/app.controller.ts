import { Controller, Get } from '@nestjs/common';
import { AppService } from './app.service';

//esse aqui seria o equivalente ao controller do .net mvc
@Controller('api')
export class AppController {
  constructor(private readonly appService: AppService) {}

  //esses métodos são chamados pelo cliente
  //esse método seria a action
  @Get("")
  getHello(): string {
    return this.appService.getHello();
  }
}
