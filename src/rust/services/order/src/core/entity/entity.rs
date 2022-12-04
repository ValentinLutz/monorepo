use std::time::Instant;

type OrderId = String;

enum Status {
    OrderPlaced,
    OrderInProgress,
    OrderCanceled,
    OrderCompleted,
}

pub struct Order {
    order_id: OrderId,
    creation_date: Instant,
    status: Status,
    workflow: String,
    items: Vec<OrderItem>,
}

struct OrderItem {
    order_item_id: i32,
    order_id: OrderId,
    name: String,
    creation_date: Instant,
}
