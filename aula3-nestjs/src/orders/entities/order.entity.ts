import { BelongsTo, Column, DataType, ForeignKey, Model, PrimaryKey, Table } from "sequelize-typescript";
import { Account } from "src/accounts/entities/account.entity";

export enum OrderStatus {
    PENDING = "pending",
    APPROVED = "approved",
    CANCELED = "canceled",
}

@Table({
    tableName: "orders",
    createdAt: "created_at",
    updatedAt: "updated_at",
})
export class Order extends Model {

    @PrimaryKey
    @Column({type: DataType.UUIDV4, defaultValue: DataType.UUIDV4})
    id: string;

    @Column({type: DataType.DECIMAL(10,2), allowNull: false})
    amount: number;

    @Column({type: DataType.STRING(16), allowNull: false})
    credit_card_number: string;

    @Column({ type: DataType.STRING(255), allowNull: false })
    credit_card_name: string;

    @Column({ allowNull: false, defaultValue: OrderStatus.PENDING })
    status: OrderStatus;

    @ForeignKey(() => Account)
    @Column({
        type: DataType.UUID,
        allowNull: false,
    })
    account_id: string;

    @BelongsTo(() => Account)
    account: Account;

}
