import { Module } from '@nestjs/common';
import { SequelizeModule } from '@nestjs/sequelize';

import { AccountsController } from './accounts.controller';
import { AccountsService } from './accounts.service';
import { Account } from './entities/account.entity';
import { AccountStorageService } from './account-storage/account-storage.service';

@Module({

  imports: [

    SequelizeModule.forFeature([Account]),

  ],

  controllers: [AccountsController],
  providers: [AccountsService, AccountStorageService]
})
export class AccountsModule {}
