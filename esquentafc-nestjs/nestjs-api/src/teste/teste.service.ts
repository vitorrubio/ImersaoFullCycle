import { Injectable } from '@nestjs/common';

@Injectable()
export class TesteService {

    metodo1(id:string): string {
        return 'deu certo mesmo !!! ' + id;
    }
}
