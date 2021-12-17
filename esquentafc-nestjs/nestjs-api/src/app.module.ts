import { Module } from '@nestjs/common';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { TesteController } from './teste/teste.controller';
import { TesteService } from './teste/teste.service';
import { OrdersModule } from './orders/orders.module';
import { SequelizeModule } from '@nestjs/sequelize';
import { Order } from './orders/entities/order.entity';
import { join } from 'path';

//es7 decorators
@Module({
  imports: [OrdersModule, 
    // SequelizeModule.forRoot({
    //   dialect: 'sqlite',
    //   host:join(__dirname, 'database.sqlite'),
    //   models:[Order],
    //   autoLoadModels: true,

    SequelizeModule.forRoot({
      dialect: 'mysql',
      host: 'db',
      port: 3306,
      database: 'fin',
      username: 'root',
      password: 'root',
      models: [Order],
      autoLoadModels: true,
  })],
  controllers: [AppController, TesteController],
  providers: [AppService, TesteService],
})
export class AppModule {}
