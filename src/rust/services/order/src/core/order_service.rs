use std::io::Error;
use crate::core::entity::entity::Order;
use crate::core::port::incoming;
use crate::core::port::incoming::OrderService;

impl incoming::OrderService for OrderService {
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