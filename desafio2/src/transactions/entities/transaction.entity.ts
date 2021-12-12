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
        defaultValue: TransactionStatus.PENDING,
    })
    status: TransactionStatus;


    @Column({ type: DataType.STRING(16), allowNull: false })
    credit_card_number: string;

    @Column({ type: DataType.STRING(255), allowNull: false })
    credit_card_name: string;


    @Column({ allowNull: false })
    credit_card_month: number;

    @Column({ allowNull: false })
    credit_card_year: number;

    @Column({ allowNull: false })
    credit_card_cvv: number;

}
