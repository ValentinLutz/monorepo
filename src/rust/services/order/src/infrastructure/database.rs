use postgres::{Client, Error, NoTls};

pub fn connect() -> Result<Client, Error> {
    Client::connect("host=localhost port=9432 user=test password=test dbname=dev_db", NoTls)
}