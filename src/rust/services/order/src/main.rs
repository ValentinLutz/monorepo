use actix_web::{web, App, HttpServer};
use postgres::{Client, NoTls};

use adapter::order_api;

mod adapter;
mod core;
mod infrastructure;

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    let client = Client::connect(
        "host=localhost port=9432 user=test password=test dbname=dev_db",
        NoTls,
    )?;

    HttpServer::new(|| {
        App::new()
            .app_data(web::Data::new(client))
            .service(web::scope("/api").configure(order_api::init_routes))
    })
    .bind(("0.0.0.0", 8080))?
    .run()
    .await
}
