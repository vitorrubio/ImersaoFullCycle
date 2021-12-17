import { Controller, Get, Param } from '@nestjs/common';
import { TesteService } from './teste.service';

@Controller('teste')
//essa classe é responsável por receber os dados do frontend e enviar para o backend
export class TesteController {

    constructor(private readonly testeSvc: TesteService) { }

    @Get(':id')

    //esse metodo vai receber um parametro
    acao(@Param('id') id:string):string {
        return this.testeSvc.metodo1(id);
    }
}
