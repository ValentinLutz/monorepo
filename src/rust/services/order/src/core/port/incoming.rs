use std::io::Error;

use crate::core::entity::entity;

pub trait OrderService {
    fn get_orders(limit: i32, offset: i32) -> Result<Vec<entity::Order>, Error>;
    fn place_order(item_names: Vec<String>) -> Result<entity::Order, Error>;
    fn get_order(order_id: String) -> Result<entity::Order, Error>;
}