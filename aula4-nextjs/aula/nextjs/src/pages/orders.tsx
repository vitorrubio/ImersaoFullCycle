import { GetServerSideProps } from "next";

const OrdersPage = () => {
    return (
        <div>
            <h2>Orders</h2>
        </div>
    );
};

export default OrdersPage;

const getServerSideProps: GetServerSideProps = async (context) => {
    
    return {
        props: {
            name: "FullCycle"
        }
    }
}

