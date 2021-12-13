import { Column, DataType, Model, PrimaryKey, Table } from "sequelize-typescript";


export enum TransactionStatus {
    PENDING = "pending",
    APPROVED = "approved",
    CANCELED = "canceled",
}

@Table({
    tableName: "transactions",
    createdAt: "created_at",
    updatedAt: "updated_at",
})
export class Transaction extends Model {

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
    })
    account_id: string;

}
