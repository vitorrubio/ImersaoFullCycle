import { Injectable, NotFoundException } from '@nestjs/common';
import { InjectModel } from '@nestjs/sequelize';
import { CreateOrderDto } from './dto/create-order.dto';
import { UpdateOrderDto } from './dto/update-order.dto';
import { Order } from './entities/order.entity';

@Injectable()
export class OrdersService {

  constructor(@InjectModel(Order) private orderModel: typeof Order) { }

  create(createOrderDto: CreateOrderDto) {
    return  this.orderModel.create(createOrderDto);
  }

  findAll() {
    return this.orderModel.findAll();
  }

  async findOne(id: string) {
    const order:Order = await this.orderModel.findByPk(id);
    if (order == null) {
      throw new NotFoundException('invalid order id');
    }

    return order;
  }

  async update(id: string, updateOrderDto: UpdateOrderDto) {
    const ordr = await this.orderModel.findByPk(id);
    if (ordr == null) {
      throw new NotFoundException('invalid order id');
    }

    return ordr.update(updateOrderDto);
  }

  async remove(id: string) {
    const ordr = await this.orderModel.findByPk(id);
    if(ordr != null)
      ordr.destroy();
  }
}
