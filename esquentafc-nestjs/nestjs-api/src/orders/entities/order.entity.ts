import { Column, DataType, Model, PrimaryKey, Table } from "sequelize-typescript";

export enum OrderStatus {
    Pending = 'pending',
    Approved = 'approved'
}



@Table({
    tableName: 'orders',
    createdAt: 'created_at',
    updatedAt: 'updated_at',
})
//essa classe representa a tabela orders
export class Order extends Model {
    
    @PrimaryKey
    @Column({
        type: DataType.UUID,
        defaultValue: DataType.UUIDV4,
    })
    id: string;

    @Column({
        type: DataType.DECIMAL(10, 2), 
        allowNull: false,
    })
    amount: number;

    @Column({
        allowNull: false,
        defaultValue: OrderStatus.Pending,
    })
    status: OrderStatus;


}
