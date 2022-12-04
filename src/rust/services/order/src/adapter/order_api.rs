use actix_web::{get, post, web, HttpResponse};

pub fn init_routes(config: &mut web::ServiceConfig) {
    config.service(get_orders);
    config.service(post_orders);
    config.service(get_order);
}

#[get("/orders")]
pub async fn get_orders() -> HttpResponse {
    HttpResponse::Ok().body("Hello world!")
}

#[post("/orders")]
pub async fn post_orders() -> HttpResponse {
    HttpResponse::Ok().body("Hello world!")
}

#[get("/orders/{order_id}")]
pub async fn get_order(path: web::Path<String>) -> HttpResponse {
    HttpResponse::Ok().body("Hello world!")
}
