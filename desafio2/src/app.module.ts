import { Module } from '@nestjs/common';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { SequelizeModule } from '@nestjs/sequelize';
import { join } from 'path';
import { TransactionsModule } from './transactions/transactions.module';
import { Transaction } from './transactions/entities/transaction.entity';

//es7 decorators
@Module({
  imports: [
    TransactionsModule,
    SequelizeModule.forRoot({
      dialect: 'sqlite',
      host:join(__dirname, 'database.sqlite'),
      models: [Transaction],
      autoLoadModels: true
    })
  ],
  controllers: [AppController],
  providers: [AppService],
})
export class AppModule {}
