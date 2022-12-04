use crate::core::port::incoming;

impl incoming::OrderService for Order {
    fn get_orders(limit: i32, offset: i32) -> Result<Vec<Order>, Error> {
        todo!()
    }

    fn place_order(item_names: Vec<String>) -> Result<Order, Error> {
        todo!()
    }

    fn get_order(order_id: String) -> Result<Order, Error> {
        todo!()
    }
}