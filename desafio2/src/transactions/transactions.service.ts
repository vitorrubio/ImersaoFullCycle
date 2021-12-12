import { Injectable, NotFoundException } from '@nestjs/common';
import { InjectModel } from '@nestjs/sequelize';
import { CreateTransactionDto } from './dto/create-transaction.dto';
import { UpdateTransactionDto } from './dto/update-transaction.dto';
import { Transaction } from './entities/transaction.entity';

@Injectable()
export class TransactionsService {


  constructor(
    @InjectModel(Transaction)
    private transactionModel: typeof Transaction
  ) { }

  create(createTransactionDto: CreateTransactionDto) {
    return this.transactionModel.create(createTransactionDto);
  }

  findAll() {
    return this.transactionModel.findAll();
  }

  async findOne(id: string) {

    const tr: Transaction = await this.transactionModel.findByPk(id);

    if (tr == null) {
      throw new NotFoundException('invalid transaction id');
    }

    return tr;
  }

  async update(id: string, updateTransactionDto: UpdateTransactionDto) {
    const trans = await this.findOne(id);
    trans.update(updateTransactionDto);
    return trans;
  }

  async remove(id: string) {
    const order = await this.findOne(id);
    order.destroy();
  }
}
